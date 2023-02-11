package nettool

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"

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

// [old], [new] must be without port
// if [onScheme] is empty, deal with both "http" and "https" scheme
// if [onPort] is -1, deal with all valid port
func ModifyFileOriginIP(old, new, onScheme string, onPort int, keepScheme, keepPort bool, fPaths ...string) error {
	for _, fPath := range fPaths {
		data, err := os.ReadFile(fPath)
		if err != nil {
			return err
		}
		text, err := strs.ModifyOriginIP(string(data), old, new, onScheme, onPort, keepScheme, keepPort, false)
		if err != nil {
			return err
		}
		if err := os.WriteFile(fPath, []byte(text), os.ModePerm); err != nil {
			return err
		}
	}
	return nil
}

// [old], [new] must be without port
// if [onScheme] is empty, deal with both "http" and "https" scheme
// if [onPort] is -1, deal with all valid port
func ModifyFile1stOriginIP(old, new, onScheme string, onPort int, keepScheme, keepPort bool, fPaths ...string) error {
	for _, fPath := range fPaths {
		data, err := os.ReadFile(fPath)
		if err != nil {
			return err
		}
		text, err := strs.ModifyOriginIP(string(data), old, new, onScheme, onPort, keepScheme, keepPort, true)
		if err != nil {
			return err
		}
		if err := os.WriteFile(fPath, []byte(text), os.ModePerm); err != nil {
			return err
		}
	}
	return nil
}
