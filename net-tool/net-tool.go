package nettool

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math"
	"net"
	"net/http"
	"os"
	"regexp"
	"sort"
	"strings"

	. "github.com/digisan/go-generics/v2"
	"github.com/digisan/gotk/strs"
)

func CheckIP(ip string) bool {
	return net.ParseIP(ip) != nil
}

func LocalIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return fmt.Sprintf("%v", localAddr.IP)
}

func PublicIP() string {
	req, err := http.Get("http://ip-api.com/json/")
	if err != nil {
		return err.Error()
	}
	defer req.Body.Close()

	body, err := io.ReadAll(req.Body)
	if err != nil {
		return err.Error()
	}

	var ip struct {
		Query string
	}
	err = json.Unmarshal(body, &ip)
	if err != nil {
		return err.Error()
	}
	return ip.Query
}

// if [scheme] is false && [portOld] is -1, then replace any local url's old port to new port.
func ReplacePort4LocalUrl(portOld, portNew int, scheme, only1st bool, fPaths ...string) error {

	for _, fPath := range fPaths {

		data, err := os.ReadFile(fPath)
		if err != nil {
			return err
		}

		text := string(data)
		locIP := LocalIP()

		type sp struct {
			s string
			p int
		}
		spGrp := make([]sp, 3)
		for i := 0; i < len(spGrp); i++ {
			spGrp[i].p = math.MaxInt32
		}

		var rLocalIPs []*regexp.Regexp
		if scheme {
			rLocalIPs = []*regexp.Regexp{
				regexp.MustCompile(fmt.Sprintf(`https?://localhost:%d/?`, portOld)),
				regexp.MustCompile(fmt.Sprintf(`https?://127.0.0.1:%d/?`, portOld)),
				regexp.MustCompile(fmt.Sprintf(`https?://%s:%d/?`, locIP, portOld)),
			}
		} else {
			if portOld != -1 {
				rLocalIPs = []*regexp.Regexp{
					regexp.MustCompile(fmt.Sprintf(`localhost:%d/?`, portOld)),
					regexp.MustCompile(fmt.Sprintf(`127.0.0.1:%d/?`, portOld)),
					regexp.MustCompile(fmt.Sprintf(`%s:%d/?`, locIP, portOld)),
				}
			} else {
				rLocalIPs = []*regexp.Regexp{
					regexp.MustCompile(`localhost:\d+/?`),
					regexp.MustCompile(`127.0.0.1:\d+/?`),
					regexp.MustCompile(fmt.Sprintf(`%s:\d+/?`, locIP)),
				}
			}
		}

		var (
			sPortOld = fmt.Sprintf(":%d", portOld)
			sPortNew = fmt.Sprintf(":%d", portNew)
		)

		for i, rip := range rLocalIPs {
			if only1st {
				if found := rip.FindString(text); found != "" {
					if p := strings.Index(text, found); p >= 0 {
						new := ""
						if portOld == -1 {
							new = strs.TrimTailFromLast(found, ":") + sPortNew
							if strings.HasSuffix(found, "/") {
								new += "/"
							}
						} else {
							new = strings.Replace(found, sPortOld, sPortNew, 1)
						}
						spGrp[i].s = strings.Replace(text, found, new, 1)
						spGrp[i].p = p
					}
				}
			} else {
				text = rip.ReplaceAllStringFunc(text, func(s string) string {
					if portOld == -1 {
						if strings.HasSuffix(s, "/") {
							return strs.TrimTailFromLast(s, ":") + sPortNew + "/"
						}
						return strs.TrimTailFromLast(s, ":") + sPortNew
					}
					return strings.ReplaceAll(s, sPortOld, sPortNew)
				})
			}
		}

		if only1st {
			sort.Slice(spGrp, func(i, j int) bool {
				return spGrp[i].p < spGrp[j].p
			})
			if spGrp[0].p < math.MaxInt32 {
				text = spGrp[0].s
			}
		}

		if err := os.WriteFile(fPath, []byte(text), os.ModePerm); err != nil {
			return err
		}
	}
	return nil
}

// if 'toPub' or 'toLoc' is true, 'aimIP' is ignored.
// if 'toPub' & 'toLoc' are both true, return error.
// otherwise, 'aimIP' must be valid.
func Loc127To(toPub, toLoc bool, aimIP string, scheme, only1st bool, fPaths ...string) error {

	if !toPub && !toLoc && !CheckIP(aimIP) {
		return fmt.Errorf("[%v] is invalid IP address", aimIP)
	}

	if toPub && toLoc {
		return fmt.Errorf("one of [toPub, toLoc] can only be true")
	}

	if toPub {
		aimIP = PublicIP()
	}
	if toLoc {
		aimIP = LocalIP()
	}

	for _, fPath := range fPaths {

		data, err := os.ReadFile(fPath)
		if err != nil {
			return err
		}
		text := string(data)

		type sp struct {
			s string
			p int
		}
		spGrp := make([]sp, 3)
		for i := 0; i < len(spGrp); i++ {
			spGrp[i].p = math.MaxInt32
		}

		var rLocalIPs []*regexp.Regexp
		if scheme {
			rLocalIPs = []*regexp.Regexp{
				regexp.MustCompile(`https?://localhost(:\d+)?/?`),
				regexp.MustCompile(`https?://127.0.0.1(:\d+)?/?`),
			}
		} else {
			rLocalIPs = []*regexp.Regexp{
				regexp.MustCompile(`localhost(:\d+)?`),
				regexp.MustCompile(`127.0.0.1(:\d+)?`),
			}
		}

		for i, rip := range rLocalIPs {
			if only1st {
				if found := rip.FindString(text); found != "" {
					if p := strings.Index(text, found); p >= 0 {
						new := strs.ReplaceFirstOnAnyOf(found, aimIP, "localhost", "127.0.0.1")
						spGrp[i].s = strings.Replace(text, found, new, 1)
						spGrp[i].p = p
					}
				}
			} else {
				text = rip.ReplaceAllStringFunc(text, func(s string) string {
					return strs.ReplaceAllOnAnyOf(s, aimIP, "localhost", "127.0.0.1")
				})
			}
		}

		if only1st {
			sort.Slice(spGrp, func(i, j int) bool {
				return spGrp[i].p < spGrp[j].p
			})
			if spGrp[0].p < math.MaxInt32 {
				text = spGrp[0].s
			}
		}

		if err := os.WriteFile(fPath, []byte(text), os.ModePerm); err != nil {
			return err
		}
	}
	return nil
}

// e.g 192.168.1.10 to public IP
func LocIP2PubIP(scheme, only1st bool, fPaths ...string) error {

	var (
		pubIP = PublicIP()
		locIP = LocalIP()
	)

	for _, fPath := range fPaths {

		data, err := os.ReadFile(fPath)
		if err != nil {
			return err
		}
		text := string(data)

		type sp struct {
			s string
			p int
		}
		spGrp := make([]sp, 3)
		for i := 0; i < len(spGrp); i++ {
			spGrp[i].p = math.MaxInt32
		}

		var rLocalIPs []*regexp.Regexp
		if scheme {
			rLocalIPs = []*regexp.Regexp{
				regexp.MustCompile(fmt.Sprintf(`https?://%s(:\d+)?/?`, locIP)),
			}
		} else {
			rLocalIPs = []*regexp.Regexp{
				regexp.MustCompile(fmt.Sprintf(`%s(:\d+)?`, locIP)),
			}
		}

		for i, rip := range rLocalIPs {
			if only1st {
				if found := rip.FindString(text); found != "" {
					if p := strings.Index(text, found); p >= 0 {
						new := strs.ReplaceFirstOnAnyOf(found, pubIP, locIP)
						spGrp[i].s = strings.Replace(text, found, new, 1)
						spGrp[i].p = p
					}
				}
			} else {
				text = rip.ReplaceAllStringFunc(text, func(s string) string {
					return strs.ReplaceAllOnAnyOf(s, pubIP, locIP)
				})
			}
		}

		if only1st {
			sort.Slice(spGrp, func(i, j int) bool {
				return spGrp[i].p < spGrp[j].p
			})
			if spGrp[0].p < math.MaxInt32 {
				text = spGrp[0].s
			}
		}

		if err := os.WriteFile(fPath, []byte(text), os.ModePerm); err != nil {
			return err
		}
	}
	return nil
}

func ModifyOriginOrIP(src, dest string, scheme, only1st, rmPort bool, fPaths ...string) error {

	r := IF(scheme, regexp.MustCompile(fmt.Sprintf(`https?://%s(:\d+)?/?`, src)), regexp.MustCompile(fmt.Sprintf(`%s(:\d+)?`, src)))

	for _, fPath := range fPaths {

		data, err := os.ReadFile(fPath)
		if err != nil {
			return err
		}
		text := string(data)

		if only1st {
			if found := r.FindString(text); found != "" {
				if !rmPort {
					new := strings.Replace(found, src, dest, 1)
					text = strings.Replace(text, found, new, 1)
				} else {
					text = strings.Replace(text, found, dest, 1)
				}
			}
		} else {
			if !rmPort {
				text = r.ReplaceAllStringFunc(text, func(s string) string {
					return strings.ReplaceAll(s, src, dest)
				})
			} else {
				text = r.ReplaceAllStringFunc(text, func(s string) string {
					return dest
				})
			}
		}

		if err := os.WriteFile(fPath, []byte(text), os.ModePerm); err != nil {
			return err
		}
	}
	return nil
}
