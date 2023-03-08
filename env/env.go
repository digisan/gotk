package env

import (
	"fmt"
	"log"
	"os"
	"strings"

	. "github.com/digisan/go-generics/v2"
	fd "github.com/digisan/gotk/file-dir"
	"github.com/digisan/gotk/strs"
)

// ChunkToMap :
func ChunkToMap(fPath, markStart, markEnd, sep string, env, pathVal2abs bool) map[string]string {

	m := make(map[string]string)

	proc := false
	_, err := fd.FileLineScan(fPath, func(ln string) (bool, string) {

		ln = strings.Trim(ln, " \t")
		if !strings.HasPrefix(ln, "#") {
			ln = strings.Trim(strings.SplitN(ln, "#", 2)[0], " \t") // support comment
		}

		if ln == markStart && !proc {
			proc = true
			return false, ""
		}
		if ln == markEnd && proc {
			proc = false
			return false, ""
		}
		if proc && strings.Contains(ln, sep) {
			ss := strings.SplitN(ln, sep, 2)
			ss[0] = strings.Trim(ss[0], " \t")
			ss[1] = strings.Trim(ss[1], " \t")
			m["$"+ss[0]] = ss[1]
		}
		return false, ""
	}, "")
	if err != nil {
		log.Fatalf("%v", err)
	}

	keyVars, _ := MapToKVs(m, func(i string, j string) bool { return len(i) > len(j) }, nil)

	// replace '$XYZ $XY $X' to its value
AGAIN1:
	for key, value := range m {
		for _, variable := range keyVars {
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

	keyVars, _ = MapToKVs(m, func(i string, j string) bool { return len(i) > len(j) }, nil)

	// replace '${XYZ} ${XY} ${X}' to its value
AGAIN3:
	for key, value := range m {
		for _, variable := range keyVars {
			valued := fmt.Sprintf("${%s}", variable)
			if strings.Contains(value, valued) {
				m[key] = strings.ReplaceAll(value, valued, m[variable])
				goto AGAIN3
			}
		}
	}

	if pathVal2abs {
		for key, value := range m {
			if strs.HasAnyPrefix(value, "~/", "./", "../") {
				absPath, _ := fd.AbsPath(value, false)
				m[key] = absPath
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

// EnvValued : if m is empty, environment variables as m apply to s
func EnvValued(s string, m map[string]string) (valStr string) {

	if len(m) == 0 {
		m = make(map[string]string)
		for _, e := range os.Environ() {
			pair := strings.SplitN(e, "=", 2)
			m[pair[0]] = pair[1]
		}
	}

	keyVars, _ := MapToKVs(m, func(i string, j string) bool { return len(i) > len(j) }, nil)
	valStr = s

	// replace '$XYZ $XY $X' to its value
	for _, variable := range keyVars {
		valued := fmt.Sprintf("$%s", variable)
		valStr = strings.ReplaceAll(valStr, valued, m[variable])
	}

	// replace '${XYZ} ${XY} ${X}' to its value
	for _, variable := range keyVars {
		valued := fmt.Sprintf("${%s}", variable)
		valStr = strings.ReplaceAll(valStr, valued, m[variable])
	}

	return
}
