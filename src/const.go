package src

import (
	"fmt"
	"strings"

	. "github.com/digisan/go-generics"
	"github.com/digisan/gotk/strs"
)

// + 'import _ "embed"'
//
// + '//go:embed ***.go'
//
// + 'var src string'
func ValuesFromConsts[T any](src string) (values []T, consts []string, err error) {

	Type := fmt.Sprintf("%T", *new(T))
	Type = strs.TrimHeadToLast(Type, ".")

	flag := false
	_, err = strs.StrLineScan(src, func(line string) (bool, string) {
		ln := strings.TrimSpace(line)
		if strings.HasPrefix(ln, "//") {
			return false, ""
		}
		if !flag && ln == "const (" {
			flag = true
			return false, ""
		}
		if flag && ln == ")" {
			flag = false
			return false, ""
		}
		if flag {
			if strings.Contains(ln, "=") && strings.Contains(ln, Type) {
				ss := strings.Split(ln, "=")
				if strings.Contains(ss[0], Type) {

					// const name
					cst := strings.TrimSpace(strings.TrimSuffix(strings.TrimSpace(ss[0]), Type))

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
							consts = append(consts, cst)
						} else if v, ok := AnyTryToType[T](val); ok {
							values = append(values, v)
							consts = append(consts, cst)
						}
						return true, ln

					} else { // other type value
						val := strings.TrimSpace(strs.TrimTailFromFirst(valStr, "//"))
						if v, err := AnyToType[T](val); err == nil {
							values = append(values, v)
							consts = append(consts, cst)
						} else if v, ok := AnyTryToType[T](val); ok {
							values = append(values, v)
							consts = append(consts, cst)
						}
						return true, ln
					}
				}
			}
		}
		return false, ""
	})
	return
}

// + 'import _ "embed"'
//
// + '//go:embed ***.go'
//
// + 'var src string'
func MapFromConsts[T any](src string) (m map[string]T, err error) {
	values, consts, err := ValuesFromConsts[T](src)
	if err != nil {
		return nil, err
	}
	m = make(map[string]T)
	for i, val := range values {
		m[consts[i]] = val
	}
	return
}
