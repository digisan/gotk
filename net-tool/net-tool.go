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

func ChangeLocalUrlPort(portOld, portNew int, strict, firstOnly bool, fpaths ...string) error {

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

func LocIP2PubIP(strict, firstOnly bool, fpaths ...string) error {

	for _, fpath := range fpaths {

		data, err := os.ReadFile(fpath)
		if err != nil {
			return err
		}

		src := string(data)
		locIP := LocalIP()
		pubIP := PublicIP()

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
				regexp.MustCompile(fmt.Sprintf(`https?://%s(:\d+)?/?`, locIP)),
			}
		} else {
			rLocalIPs = []*regexp.Regexp{
				regexp.MustCompile(`localhost(:\d+)?`),
				regexp.MustCompile(`127.0.0.1(:\d+)?`),
				regexp.MustCompile(fmt.Sprintf(`%s(:\d+)?`, locIP)),
			}
		}

		for i, rip := range rLocalIPs {
			if firstOnly {
				if found := rip.FindString(src); found != "" {
					if p := strings.Index(src, found); p >= 0 {
						new := strs.ReplaceFirstOnAnyOf(found, pubIP, "localhost", "127.0.0.1", locIP)
						spGrp[i].s = strings.Replace(src, found, new, 1)
						spGrp[i].p = p
					}
				}
			} else {
				src = rip.ReplaceAllStringFunc(src, func(s string) string {
					return strs.ReplaceAllOnAnyOf(s, pubIP, "localhost", "127.0.0.1", locIP)
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
