package strings

import (
	"regexp"
	"strings"

	"github.com/digisan/gotk/generics/ti"
)

var (
	sIndex = strings.Index
)

func IndexAll(s, sub string) (starts, ends []int) {
	for i := sIndex(s, sub); i != -1; i = sIndex(s, sub) {
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
	for i := 0; i < ti.Min(len(ranges), len(ns)); i++ {
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
