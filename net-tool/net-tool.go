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

// [old], [new] must be without port
// if [onScheme] is empty, deal with both "http" and "https" scheme
// if [onPort] is -1, deal with all valid port
func ModifyOriginIP(text, old, new, onScheme string, onPort int, keepScheme, keepPort, only1st bool) (string, error) {

	if NotIn(onScheme, "", "http", "https") {
		return "", fmt.Errorf("[onScheme] can only be 'http', 'https' OR empty for no scheme or any scheme")
	}
	if NotInCloseRange(onPort, 0, 65535) && onPort != -1 {
		return "", fmt.Errorf("[onPort] can only be [0, 65535] OR -1 for no port or any port")
	}

	rmc, fsf := regexp.MustCompile, fmt.Sprintf

	var r *regexp.Regexp
	switch {
	case len(onScheme) > 0 && onPort > -1:
		r = rmc(fsf(`%s://%s:%d`, onScheme, old, onPort))

	case len(onScheme) == 0 && onPort > -1:
		r = rmc(fsf(`(https?://)?%s:%d`, old, onPort))

	case len(onScheme) > 0 && onPort < 0:
		r = rmc(fsf(`%s://%s(:\d+)?`, onScheme, old))

	case len(onScheme) == 0 && onPort < 0:
		r = rmc(fsf(`(https?://)?%s(:\d+)?`, old))
	}

	if only1st {

		if found := r.FindString(text); found != "" {
			switch {
			case keepScheme && keepPort:
				new := strings.Replace(found, old, new, 1)
				text = strings.Replace(text, found, new, 1)

			case keepScheme && !keepPort:
				toRm := strs.TrimAnyPrefix(found, "https://", "http://")
				text = strings.Replace(text, toRm, new, 1)

			case !keepScheme && keepPort:
				if p := strings.LastIndex(found, ":"); p > -1 && IsNumeric(found[p+1:]) {
					toRm := found[:p]
					m := map[int]string{
						strings.Index(text, "https://"+toRm): "https://" + toRm,
						strings.Index(text, "http://"+toRm):  "http://" + toRm,
						strings.Index(text, toRm):            toRm,
					}
					keys, _ := MapToKVs(m, nil, nil)
					if k, ok := MinFloor(0, true, keys...); ok {
						text = strings.Replace(text, m[k], new, 1)
					}
				}

			// case !keepScheme && !keepPort:
			default:
				m := map[int]string{
					strings.Index(text, "https://"+found): "https://" + found,
					strings.Index(text, "http://"+found):  "http://" + found,
					strings.Index(text, found):            found,
				}
				keys, _ := MapToKVs(m, nil, nil)
				if k, ok := MinFloor(0, true, keys...); ok {
					text = strings.Replace(text, m[k], new, 1)
				}
			}
		}

	} else {

		switch {
		case keepScheme && keepPort:
			return r.ReplaceAllStringFunc(text, func(s string) string {
				return strings.ReplaceAll(s, old, new)
			}), nil

		case keepScheme && !keepPort:
			return r.ReplaceAllStringFunc(text, func(s string) string {
				if p := strings.Index(s, "://"); p > -1 {
					return s[:p] + "://" + new
				}
				return new
			}), nil

		case !keepScheme && keepPort:
			return r.ReplaceAllStringFunc(text, func(s string) string {
				if p := strings.LastIndex(s, ":"); p > -1 {
					if port, ok := AnyTryToType[int](s[p+1:]); ok {
						return new + ":" + fmt.Sprint(port)
					}
				}
				return new
			}), nil

		// case !keepScheme && !keepPort:
		default:
			return r.ReplaceAllStringFunc(text, func(s string) string {
				return new
			}), nil
		}
	}

	return text, nil
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
		text, err := ModifyOriginIP(string(data), old, new, onScheme, onPort, keepScheme, keepPort, false)
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
		text, err := ModifyOriginIP(string(data), old, new, onScheme, onPort, keepScheme, keepPort, true)
		if err != nil {
			return err
		}
		if err := os.WriteFile(fPath, []byte(text), os.ModePerm); err != nil {
			return err
		}
	}
	return nil
}
