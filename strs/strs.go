package strs

import (
	"fmt"
	"log"
	"regexp"
	"runtime"
	"sort"
	"strings"

	. "github.com/digisan/go-generics/v2"
	"golang.org/x/net/html"
)

func IsIn(ignoreCase, ignoreSpace bool, s string, group ...string) bool {
	if ignoreCase {
		s = strings.ToLower(s)
		group = FilterMap4SglTyp(group, nil, func(i int, e string) string { return strings.ToLower(e) })
	}
	if ignoreSpace {
		s = strings.TrimSpace(s)
		group = FilterMap4SglTyp(group, nil, func(i int, e string) string { return strings.TrimSpace(e) })
	}
	return In(s, group...)
}

func IsNotIn(insensitive, ignorespace bool, s string, group ...string) bool {
	return !IsIn(insensitive, ignorespace, s, group...)
}

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

func ContainsAny(s string, aims ...string) bool {
	for _, aim := range aims {
		if strings.Contains(s, aim) {
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

func SplitPartTo[T any](s, sep string, idx int) T {
	ss := strings.Split(s, sep)
	if idx < 0 && idx >= len(ss) {
		log.Fatalf("index@ [%d] of '%s' is out of range", idx, s)
	}
	if rt, ok := AnyTryToType[T](ss[idx]); ok {
		return rt
	}
	panic(fmt.Sprintf("index@ [%d] of '%s' cannot be converted to [%T]", idx, s, *new(T)))
}

// idx 1 is the last element, idx 2 is the second last, etc...
func SplitPartFromLastTo[T any](s, sep string, idx int) T {
	laststr := Last(strings.Split(s, sep), idx)
	if rt, ok := AnyTryToType[T](laststr); ok {
		return rt
	}
	panic(fmt.Sprintf("index@ [%d] of '%s' cannot be converted to [%T]", idx, s, *new(T)))
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

func SortPaths(paths ...string) []string {
	sort.SliceStable(paths, func(i, j int) bool {
		pathi, pathj := paths[i], paths[j]
		ni, nj := strings.Count(pathi, "."), strings.Count(pathj, ".")
		minDot := Min(ni, nj)
		ssi, ssj := strings.Split(pathi, "."), strings.Split(pathj, ".")
	NEXT:
		for i := 0; i < minDot+1; i++ {
			si, sj := ssi[i], ssj[i]
			if si == sj {
				continue NEXT
			}
			if IsUint(si) && IsUint(sj) {
				idxI, _ := AnyTryToType[uint](si)
				idxJ, _ := AnyTryToType[uint](sj)
				if idxI == idxJ {
					continue NEXT
				}
				return idxI < idxJ
			}
			return si < sj // ascii ASC, uppercase first
		}
		return ni < nj
	})
	return paths
}
