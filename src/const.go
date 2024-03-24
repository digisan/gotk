package src

import (
	"fmt"
	"strings"

	. "github.com/digisan/go-generics"
	fd "github.com/digisan/gotk/file-dir"
	"github.com/digisan/gotk/strs"
)

// fSrcPath: _, fSrcPath, _, _ := runtime.Caller(0/1)
func ValuesFromConsts[T any](fSrcPath string) (values []T, err error) {

	Type := fmt.Sprintf("%T", *new(T))
	Type = strs.TrimHeadToLast(Type, ".")

	flag := false
	_, err = fd.FileLineScan(fSrcPath, func(line string) (bool, string) {
		ln := strings.TrimSpace(line)
		if ln == "const (" {
			flag = true
		}
		if flag && ln == ")" {
			flag = false
		}
		if flag {
			if strings.Contains(ln, "=") && strings.Contains(ln, Type) {
				ss := strings.Split(ln, "=")
				if strings.Contains(ss[0], Type) {
					valStr := strings.TrimSpace(ss[1])
					if strs.HasAnyPrefix(valStr, "\"", "`") { // string type value
						pCloseQ := 1
						if strings.HasPrefix(valStr, "\"") {
							for i, r := range valStr {
								if i > 0 && r == '"' && valStr[i-1:i] != "\\" {
									pCloseQ = i
									break
								}
							}
						}
						if strings.HasPrefix(valStr, "`") {
							for i, r := range valStr {
								if i > 0 && r == '`' {
									pCloseQ = i
									break
								}
							}
						}
						val := valStr[1:pCloseQ]
						if v, err := AnyToType[T](val); err == nil {
							values = append(values, v)
						} else if v, ok := AnyTryToType[T](val); ok {
							values = append(values, v)
						}
						return true, ln

					} else { // other type value
						val := strings.TrimSpace(strs.TrimTailFromFirst(valStr, "//"))
						if v, err := AnyToType[T](val); err == nil {
							values = append(values, v)
						} else if v, ok := AnyTryToType[T](val); ok {
							values = append(values, v)
						}
						return true, ln
					}
				}
			}
		}
		return false, ""

	}, "")
	return
}
