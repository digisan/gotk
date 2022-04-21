package strs

import (
	"regexp"
	"runtime"
	"strings"

	. "github.com/digisan/go-generics/v2"
)

func Maxlen(s string, length int) string {
	if len(s) < length {
		return s
	}
	return s[:length]
}

func IndexAll(s, sub string) (starts, ends []int) {
	for i := strings.Index(s, sub); i != -1; i = strings.Index(s, sub) {
		last := 0
		if len(starts) > 0 {
			last = starts[len(starts)-1] + 1
		}
		start := last + i
		starts = append(starts, start)
		ends = append(ends, start+len(sub))
		s = s[i+1:]
	}
	return
}

func IndexAllByReg(s, sub string) (starts, ends []int) {
	r := regexp.MustCompile(regexp.QuoteMeta(sub))
	for _, pair := range r.FindAllStringIndex(s, -1) {
		starts = append(starts, pair[0])
		ends = append(ends, pair[1])
	}
	return
}

func RangeReplace(s string, ranges [][2]int, ns []string) string {
	for i := 0; i < Min(len(ranges), len(ns)); i++ {
		lenPrev := len(s)
		start, end := ranges[i][0], ranges[i][1]
		s = s[:start] + ns[i] + s[end:]
		diff := len(s) - lenPrev
		for j := i + 1; j < len(ranges); j++ {
			ranges[j][0] += diff
			ranges[j][1] += diff
		}
	}
	return s
}

func HasAnyPrefix(s string, prefixGrp ...string) bool {
	for _, prefix := range prefixGrp {
		if strings.HasPrefix(s, prefix) {
			return true
		}
	}
	return false
}

func HasAnySuffix(s string, suffixGrp ...string) bool {
	for _, suffix := range suffixGrp {
		if strings.HasSuffix(s, suffix) {
			return true
		}
	}
	return false
}

func SplitPart(s, sep string, idx int) string {
	ss := strings.Split(s, sep)
	if idx < 0 && idx >= len(ss) {
		panic("[idx] is out of range")
	}
	return ss[idx]
}

func SplitPartFromLast(s, sep string, idx int) string {
	return Last(strings.Split(s, sep), idx)
}

func TrimTailFromLast(s, mark string) string {
	if i := strings.LastIndex(s, mark); i >= 0 {
		return s[:i]
	}
	return s
}

func TrimHeadToLast(s, mark string) string {
	if i := strings.LastIndex(s, mark); i >= 0 {
		return s[i+len(mark):]
	}
	return s
}

func SplitLn(s string) []string {
	sep := MATCH(runtime.GOOS, "windows", "linux", "darwin", "\r\n", "\n", "\r", "\n")
	return strings.Split(s, sep)
}
