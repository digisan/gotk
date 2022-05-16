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

func ChangeLocalUrlPort(path string, portOld, portNew int) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	src := string(data)
	locIP := LocalIP()

	localIPs := []*regexp.Regexp{
		regexp.MustCompile(fmt.Sprintf(`https?://localhost:%d/`, portOld)),
		regexp.MustCompile(fmt.Sprintf(`https?://127.0.0.1:%d/`, portOld)),
		regexp.MustCompile(fmt.Sprintf(`https?://%s:%d/`, locIP, portOld)),
	}

	for _, ip := range localIPs {
		src = ip.ReplaceAllStringFunc(src, func(s string) string {
			return strings.ReplaceAll(s, fmt.Sprintf(":%d/", portOld), fmt.Sprintf(":%d/", portNew))
		})
	}
	return os.WriteFile(path, []byte(src), os.ModePerm)
}

func LocIP2PubIP(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	src := string(data)
	locIP := LocalIP()

	localIPs := []*regexp.Regexp{
		regexp.MustCompile(`https?://localhost(:\d+)?/`),
		regexp.MustCompile(`https?://127.0.0.1(:\d+)?/`),
		regexp.MustCompile(fmt.Sprintf(`https?://%s(:\d+)?/`, locIP)),
	}

	pubIP := PublicIP()

	for _, ip := range localIPs {
		src = ip.ReplaceAllStringFunc(src, func(s string) string {
			switch {
			case strings.Contains(s, "localhost"):
				return strings.ReplaceAll(s, "localhost", pubIP)
			case strings.Contains(s, "127.0.0.1"):
				return strings.ReplaceAll(s, "127.0.0.1", pubIP)
			case strings.Contains(s, locIP):
				return strings.ReplaceAll(s, locIP, pubIP)
			}
			panic("ERROR")
		})
	}
	return os.WriteFile(path, []byte(src), os.ModePerm)
}

func LocalhostToIP127(fpath string) error {
	data, err := os.ReadFile(fpath)
	if err != nil {
		return err
	}

	src := string(data)
	r := regexp.MustCompile(`https?://localhost(:\d+)?/`)
	src = r.ReplaceAllStringFunc(src, func(s string) string {
		return strings.ReplaceAll(s, "localhost", "127.0.0.1")
	})
	return os.WriteFile(fpath, []byte(src), os.ModePerm)
}

func IP127ToLocalhost(fpath string) error {
	data, err := os.ReadFile(fpath)
	if err != nil {
		return err
	}

	src := string(data)
	r := regexp.MustCompile(`https?://127.0.0.1(:\d+)?/`)
	src = r.ReplaceAllStringFunc(src, func(s string) string {
		return strings.ReplaceAll(s, "127.0.0.1", "localhost")
	})
	return os.WriteFile(fpath, []byte(src), os.ModePerm)
}
