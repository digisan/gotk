package strs

import (
	"fmt"
	"regexp"
	"runtime"
	"sort"
	"strconv"
	"strings"

	. "github.com/digisan/go-generics/v2"
	"golang.org/x/net/html"
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

func ReplaceFirstOnAnyOf(str string, new string, aims ...string) string {
	type sp struct {
		s string
		p int
	}
	spGrp := []sp{}
	for _, aim := range aims {
		if p := strings.Index(str, aim); p >= 0 {
			spGrp = append(spGrp, sp{s: aim, p: p})
		}
	}
	if len(spGrp) > 0 {
		sort.Slice(spGrp, func(i, j int) bool {
			if spGrp[i].p == spGrp[j].p {
				return len(spGrp[i].s) > len(spGrp[j].s)
			}
			return spGrp[i].p < spGrp[j].p
		})
		return strings.Replace(str, spGrp[0].s, new, 1)
	}
	return str
}

func ReplaceAllOnAnyOf(str string, new string, aims ...string) string {
	sort.Slice(aims, func(i, j int) bool {
		return len(aims[i]) > len((aims[j]))
	})
	sb := strings.Builder{}
	for i, aim := range aims {
		sb.WriteString(fmt.Sprintf("(%s)", aim))
		if i < len(aims)-1 {
			sb.WriteString("|")
		}
	}
	rs := sb.String()
	// fmt.Println(rs)
	r := regexp.MustCompile(rs)
	str = r.ReplaceAllStringFunc(str, func(s string) string {
		return new
	})
	return str
}

func SplitPart(s, sep string, idx int) string {
	ss := strings.Split(s, sep)
	if idx < 0 && idx >= len(ss) {
		panic("[idx] is out of range")
	}
	return ss[idx]
}

func SplitPartToNum(s, sep string, idx int) float64 {
	part := SplitPart(s, sep, idx)
	num, err := strconv.ParseFloat(part, 64)
	if err != nil {
		panic(err)
	}
	return num
}

func SplitPartToBool(s, sep string, idx int) bool {
	part := SplitPart(s, sep, idx)
	boolean, err := strconv.ParseBool(part)
	if err != nil {
		panic(err)
	}
	return boolean
}

// idx 1 is the last element, idx 2 is the second last, etc...
func SplitPartFromLast(s, sep string, idx int) string {
	return Last(strings.Split(s, sep), idx)
}

func SplitPartFromLastToNum(s, sep string, idx int) float64 {
	part := SplitPartFromLast(s, sep, idx)
	num, err := strconv.ParseFloat(part, 64)
	if err != nil {
		panic(err)
	}
	return num
}

func SplitPartFromLastToBool(s, sep string, idx int) bool {
	part := SplitPartFromLast(s, sep, idx)
	boolean, err := strconv.ParseBool(part)
	if err != nil {
		panic(err)
	}
	return boolean
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

func HtmlTextContent(htmlstr string) (rt []string) {
	domDoc := html.NewTokenizer(strings.NewReader(htmlstr))
	previousStartTokenTest := domDoc.Token()
loopDomTest:
	for {
		tt := domDoc.Next()
		switch {
		case tt == html.ErrorToken:
			break loopDomTest // End of the document,  done
		case tt == html.StartTagToken:
			previousStartTokenTest = domDoc.Token()
		case tt == html.TextToken:
			if previousStartTokenTest.Data == "script" {
				continue
			}
			if text := strings.TrimSpace(html.UnescapeString(string(domDoc.Text()))); len(text) > 0 {
				rt = append(rt, text)
			}
		}
	}
	return
}
