package nettool

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"regexp"
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

func ChangeLocalUrlPort(path string, portOld, portNew int, strict, onlyFirst bool) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	src := string(data)
	locIP := LocalIP()

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
	for _, rip := range rLocalIPs {
		if onlyFirst {
			if found := rip.FindString(src); found != "" {
				new := strings.Replace(found, sPortOld, sPortNew, 1)
				src = strings.Replace(src, found, new, 1)
				break
			}
		} else {
			src = rip.ReplaceAllStringFunc(src, func(s string) string {
				return strings.ReplaceAll(s, sPortOld, sPortNew)
			})
		}
	}
	return os.WriteFile(path, []byte(src), os.ModePerm)
}

func LocIP2PubIP(path string, strict, onlyFirst bool) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	src := string(data)
	locIP := LocalIP()
	pubIP := PublicIP()

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

	for _, rip := range rLocalIPs {
		if onlyFirst {
			if found := rip.FindString(src); found != "" {
				new := strs.ReplaceFirstOnAnyOf(found, pubIP, "localhost", "127.0.0.1", locIP)
				src = strings.Replace(src, found, new, 1)
				break
			}
		} else {
			src = rip.ReplaceAllStringFunc(src, func(s string) string {
				return strs.ReplaceAllOnAnyOf(s, pubIP, "localhost", "127.0.0.1", locIP)
			})
		}
	}
	return os.WriteFile(path, []byte(src), os.ModePerm)
}

func LocalhostToIP127(fpath string, strict, onlyFirst bool) error {
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

	if onlyFirst {
		if found := r.FindString(src); found != "" {
			new := strings.Replace(found, "localhost", "127.0.0.1", 1)
			src = strings.Replace(src, found, new, 1)
		}
	} else {
		src = r.ReplaceAllStringFunc(src, func(s string) string {
			return strings.ReplaceAll(s, "localhost", "127.0.0.1")
		})
	}

	return os.WriteFile(fpath, []byte(src), os.ModePerm)
}

func IP127ToLocalhost(fpath string, strict, onlyFirst bool) error {
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

	if onlyFirst {
		if found := r.FindString(src); found != "" {
			new := strings.Replace(found, "127.0.0.1", "localhost", 1)
			src = strings.Replace(src, found, new, 1)
		}
	} else {
		src = r.ReplaceAllStringFunc(src, func(s string) string {
			return strings.ReplaceAll(s, "127.0.0.1", "localhost")
		})
	}
	return os.WriteFile(fpath, []byte(src), os.ModePerm)
}
