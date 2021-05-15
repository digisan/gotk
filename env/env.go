package env

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/digisan/gotk/io"
	"github.com/digisan/gotk/mapslice"
)

// Chunk2Map :
func Chunk2Map(filepath, markstart, markend, sep string, env bool) map[string]string {

	m := make(map[string]string)

	proc := false
	_, err := io.FileLineScan(filepath, func(ln string) (bool, string) {
		if ln == markstart && !proc {
			proc = true
			return false, ""
		}
		if ln == markend && proc {
			proc = false
			return false, ""
		}
		if proc && strings.Contains(ln, sep) {
			ss := strings.SplitN(ln, sep, 2)
			m["$"+ss[0]] = ss[1]
		}
		return false, ""
	}, "")
	if err != nil {
		log.Fatalf("%v", err)
	}

	keyvars, _ := mapslice.KsVs2Slc(m, "KL-DESC")

	// replace '$XYZ $XY $X' to its value
AGAIN1:
	for key, value := range m {
		for _, variable := range keyvars {
			if strings.Contains(value, variable) {
				m[key] = strings.ReplaceAll(value, variable, m[variable])
				goto AGAIN1
			}
		}
	}

	// remove map each key's prefix '$'
AGAIN2:
	for key, value := range m {
		if strings.HasPrefix(key, "$") {
			m[key[1:]] = value
			delete(m, key)
			goto AGAIN2
		}
	}

	keyvars, _ = mapslice.KsVs2Slc(m, "KL-DESC")

	// replace '${XYZ} ${XY} ${X}' to its value
AGAIN3:
	for key, value := range m {
		for _, variable := range keyvars {
			valued := fmt.Sprintf("${%s}", variable)
			if strings.Contains(value, valued) {
				m[key] = strings.ReplaceAll(value, valued, m[variable])
				goto AGAIN3
			}
		}
	}

	if env {
		for key, value := range m {
			os.Setenv(key, value)
		}
	}

	return m
}
