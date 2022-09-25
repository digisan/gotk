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

func ChangeLocalUrlPort(strict, firstOnly bool, portOld, portNew int, fpaths ...string) error {

	for _, fpath := range fpaths {

		data, err := os.ReadFile(fpath)
		if err != nil {
			return err
		}

		src := string(data)
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
		if strict {
			rLocalIPs = []*regexp.Regexp{
				regexp.MustCompile(fmt.Sprintf(`https?://localhost:%d/?`, portOld)),
				regexp.MustCompile(fmt.Sprintf(`https?://127.0.0.1:%d/?`, portOld)),
				regexp.MustCompile(fmt.Sprintf(`https?://%s:%d/?`, locIP, portOld)),
			}
		} else {
			rLocalIPs = []*regexp.Regexp{
				regexp.MustCompile(fmt.Sprintf(`localhost:%d`, portOld)),
				regexp.MustCompile(fmt.Sprintf(`127.0.0.1:%d`, portOld)),
				regexp.MustCompile(fmt.Sprintf(`%s:%d`, locIP, portOld)),
			}
		}

		sPortOld, sPortNew := fmt.Sprintf(":%d", portOld), fmt.Sprintf(":%d", portNew)
		for i, rip := range rLocalIPs {
			if firstOnly {
				if found := rip.FindString(src); found != "" {
					if p := strings.Index(src, found); p >= 0 {
						new := strings.Replace(found, sPortOld, sPortNew, 1)
						spGrp[i].s = strings.Replace(src, found, new, 1)
						spGrp[i].p = p
					}
				}
			} else {
				src = rip.ReplaceAllStringFunc(src, func(s string) string {
					return strings.ReplaceAll(s, sPortOld, sPortNew)
				})
			}
		}

		if firstOnly {
			sort.Slice(spGrp, func(i, j int) bool {
				return spGrp[i].p < spGrp[j].p
			})
			if spGrp[0].p < math.MaxInt32 {
				src = spGrp[0].s
			}
		}

		if err := os.WriteFile(fpath, []byte(src), os.ModePerm); err != nil {
			return err
		}
	}
	return nil
}

// if 'toPubIP' is true, 'aimIP' is ignored.
// if 'toLocIP' is true, 'aimIP' is ignored.
// if 'toPubIP' & 'toLocIP' are both true, return error.
// otherwise, 'aimIP' must be valid.
func ChangeLocalhost(strict, firstOnly, toPubIP, toLocIP bool, aimIP string, fpaths ...string) error {

	if !toPubIP && !toLocIP && !CheckIP(aimIP) {
		return fmt.Errorf("[%v] is invalid IP address", aimIP)
	}

	if toPubIP && toLocIP {
		return fmt.Errorf("only one of [toPubIP, toLocIP] can be set true")
	}

	if toPubIP {
		aimIP = PublicIP()
	}
	if toLocIP {
		aimIP = LocalIP()
	}

	for _, fpath := range fpaths {

		data, err := os.ReadFile(fpath)
		if err != nil {
			return err
		}
		src := string(data)

		type sp struct {
			s string
			p int
		}
		spGrp := make([]sp, 3)
		for i := 0; i < len(spGrp); i++ {
			spGrp[i].p = math.MaxInt32
		}

		var rLocalIPs []*regexp.Regexp
		if strict {
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
			if firstOnly {
				if found := rip.FindString(src); found != "" {
					if p := strings.Index(src, found); p >= 0 {
						new := strs.ReplaceFirstOnAnyOf(found, aimIP, "localhost", "127.0.0.1")
						spGrp[i].s = strings.Replace(src, found, new, 1)
						spGrp[i].p = p
					}
				}
			} else {
				src = rip.ReplaceAllStringFunc(src, func(s string) string {
					return strs.ReplaceAllOnAnyOf(s, aimIP, "localhost", "127.0.0.1")
				})
			}
		}

		if firstOnly {
			sort.Slice(spGrp, func(i, j int) bool {
				return spGrp[i].p < spGrp[j].p
			})
			if spGrp[0].p < math.MaxInt32 {
				src = spGrp[0].s
			}
		}

		if err := os.WriteFile(fpath, []byte(src), os.ModePerm); err != nil {
			return err
		}
	}
	return nil
}

// e.g 192.168.1.10 to public IP
func LocIPToPubIP(strict, firstOnly bool, fpaths ...string) error {

	var (
		pubIP = PublicIP()
		locIP = LocalIP()
	)

	for _, fpath := range fpaths {

		data, err := os.ReadFile(fpath)
		if err != nil {
			return err
		}
		src := string(data)

		type sp struct {
			s string
			p int
		}
		spGrp := make([]sp, 3)
		for i := 0; i < len(spGrp); i++ {
			spGrp[i].p = math.MaxInt32
		}

		var rLocalIPs []*regexp.Regexp
		if strict {
			rLocalIPs = []*regexp.Regexp{
				regexp.MustCompile(fmt.Sprintf(`https?://%s(:\d+)?/?`, locIP)),
			}
		} else {
			rLocalIPs = []*regexp.Regexp{
				regexp.MustCompile(fmt.Sprintf(`%s(:\d+)?`, locIP)),
			}
		}

		for i, rip := range rLocalIPs {
			if firstOnly {
				if found := rip.FindString(src); found != "" {
					if p := strings.Index(src, found); p >= 0 {
						new := strs.ReplaceFirstOnAnyOf(found, pubIP, locIP)
						spGrp[i].s = strings.Replace(src, found, new, 1)
						spGrp[i].p = p
					}
				}
			} else {
				src = rip.ReplaceAllStringFunc(src, func(s string) string {
					return strs.ReplaceAllOnAnyOf(s, pubIP, locIP)
				})
			}
		}

		if firstOnly {
			sort.Slice(spGrp, func(i, j int) bool {
				return spGrp[i].p < spGrp[j].p
			})
			if spGrp[0].p < math.MaxInt32 {
				src = spGrp[0].s
			}
		}

		if err := os.WriteFile(fpath, []byte(src), os.ModePerm); err != nil {
			return err
		}
	}
	return nil
}

func LocalhostToIP127(strict, firstOnly bool, fpaths ...string) error {

	for _, fpath := range fpaths {

		data, err := os.ReadFile(fpath)
		if err != nil {
			return err
		}
		src := string(data)

		var r *regexp.Regexp
		if strict {
			r = regexp.MustCompile(`https?://localhost(:\d+)?/?`)
		} else {
			r = regexp.MustCompile(`localhost(:\d+)?`)
		}

		if firstOnly {
			if found := r.FindString(src); found != "" {
				new := strings.Replace(found, "localhost", "127.0.0.1", 1)
				src = strings.Replace(src, found, new, 1)
			}
		} else {
			src = r.ReplaceAllStringFunc(src, func(s string) string {
				return strings.ReplaceAll(s, "localhost", "127.0.0.1")
			})
		}

		if err := os.WriteFile(fpath, []byte(src), os.ModePerm); err != nil {
			return err
		}
	}
	return nil
}

func IP127ToLocalhost(strict, firstOnly bool, fpaths ...string) error {

	for _, fpath := range fpaths {

		data, err := os.ReadFile(fpath)
		if err != nil {
			return err
		}
		src := string(data)

		var r *regexp.Regexp
		if strict {
			r = regexp.MustCompile(`https?://127.0.0.1(:\d+)?/?`)
		} else {
			r = regexp.MustCompile(`127.0.0.1(:\d+)?`)
		}

		if firstOnly {
			if found := r.FindString(src); found != "" {
				new := strings.Replace(found, "127.0.0.1", "localhost", 1)
				src = strings.Replace(src, found, new, 1)
			}
		} else {
			src = r.ReplaceAllStringFunc(src, func(s string) string {
				return strings.ReplaceAll(s, "127.0.0.1", "localhost")
			})
		}

		if err := os.WriteFile(fpath, []byte(src), os.ModePerm); err != nil {
			return err
		}
	}
	return nil
}
