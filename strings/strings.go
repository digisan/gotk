package strings

import "strings"

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
