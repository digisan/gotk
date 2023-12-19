package strs

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
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

func IsNotIn(insensitive, ignoreSpace bool, s string, group ...string) bool {
	return !IsIn(insensitive, ignoreSpace, s, group...)
}

func MaxLen(s string, length int) string {
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

func ReplaceAllInFolder(oldStr, newStr, dir, ext string) (fPaths []string) {
	des, err := os.ReadDir(dir)
	if err != nil {
		return nil
	}
	ext = "." + strings.TrimPrefix(ext, ".")
	for _, de := range des {
		if de.IsDir() {
			continue
		}
		if strings.HasSuffix(de.Name(), ext) {
			fPath := filepath.Join(dir, de.Name())
			data, err := os.ReadFile(fPath)
			if err != nil {
				continue
			}
			s := strings.ReplaceAll(string(data), oldStr, newStr)
			if err := os.WriteFile(fPath, []byte(s), os.ModePerm); err == nil {
				fPaths = append(fPaths, fPath)
			}
		}
	}
	return
}

func HasAnyPrefix(s string, prefixGrp ...string) bool {
	for _, prefix := range prefixGrp {
		if strings.HasPrefix(s, prefix) {
			return true
		}
	}
	return false
}

func TrimAnyPrefix(s string, prefixGrp ...string) string {
	sort.Slice(prefixGrp, func(i, j int) bool {
		return len(prefixGrp[i]) > len(prefixGrp[j])
	})
	for _, prefix := range prefixGrp {
		before := s
		if s = strings.TrimPrefix(s, prefix); s != before {
			break
		}
	}
	return s
}

func HasAnySuffix(s string, suffixGrp ...string) bool {
	for _, suffix := range suffixGrp {
		if strings.HasSuffix(s, suffix) {
			return true
		}
	}
	return false
}

func TrimAnySuffix(s string, suffixGrp ...string) string {
	sort.Slice(suffixGrp, func(i, j int) bool {
		return len(suffixGrp[i]) > len(suffixGrp[j])
	})
	for _, suffix := range suffixGrp {
		before := s
		if s = strings.TrimSuffix(s, suffix); s != before {
			break
		}
	}
	return s
}

func ContainsAny(s string, aims ...string) bool {
	for _, aim := range aims {
		if strings.Contains(s, aim) {
			return true
		}
	}
	return false
}

func Replace1stOnAnyOf(s string, new string, aims ...string) string {
	type sp struct {
		s string
		p int
	}
	spGrp := []sp{}
	for _, aim := range aims {
		if p := strings.Index(s, aim); p >= 0 {
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
		return strings.Replace(s, spGrp[0].s, new, 1)
	}
	return s
}

func ReplaceAllOnAnyOf(s string, new string, aims ...string) string {
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
	return r.ReplaceAllStringFunc(s, func(string) string { return new })
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
	lastStr := Last(strings.Split(s, sep), idx)
	if rt, ok := AnyTryToType[T](lastStr); ok {
		return rt
	}
	panic(fmt.Sprintf("index@ [%d] of '%s' cannot be converted to [%T]", idx, s, *new(T)))
}

func TrimTailFromFirst(s, mark string) string {
	if i := strings.Index(s, mark); i >= 0 {
		return s[:i]
	}
	return s
}

func TrimTailFromLast(s, mark string) string {
	if i := strings.LastIndex(s, mark); i >= 0 {
		return s[:i]
	}
	return s
}

func TrimHeadToFirst(s, mark string) string {
	if i := strings.Index(s, mark); i >= 0 {
		return s[i+len(mark):]
	}
	return s
}

func TrimHeadToLast(s, mark string) string {
	if i := strings.LastIndex(s, mark); i >= 0 {
		return s[i+len(mark):]
	}
	return s
}

func HeadConsecutive(s string, cc ...rune) string {
	for i, c := range s {
		if NotIn(c, cc...) {
			return s[:i]
		}
	}
	return s
}

func HeadSpace(s string) string {
	return HeadConsecutive(s, ' ')
}

func HeadTblChar(s string) string {
	return HeadConsecutive(s, '\t')
}

func HeadBlank(s string) string {
	return HeadConsecutive(s, ' ', '\t')
}

func TailConsecutive(s string, cc ...rune) string {
	buf := []rune{}
	for _, c := range s {
		buf = append(buf, c)
	}
	for i, c := range Reverse(buf) {
		if NotIn(c, cc...) {
			return s[len(s)-i:]
		}
	}
	return s
}

func TailSpace(s string) string {
	return TailConsecutive(s, ' ')
}

func TailTblChar(s string) string {
	return TailConsecutive(s, '\t')
}

func TailBlank(s string) string {
	return TailConsecutive(s, ' ', '\t')
}

func SplitLn(s string) []string {
	sep := MATCH(runtime.GOOS, "windows", "linux", "darwin", "\r\n", "\n", "\r", "\n")
	return strings.Split(s, sep)
}

func HtmlTextContent(htmlStr string) (rt []string) {
	domDoc := html.NewTokenizer(strings.NewReader(htmlStr))
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

func ReversePath(path string) string {
	return strings.Join(Reverse(strings.Split(path, ".")), ".")
}

func SortPaths(paths ...string) []string {
	sort.SliceStable(paths, func(i, j int) bool {
		pathI, pathJ := paths[i], paths[j]
		nI, nJ := strings.Count(pathI, "."), strings.Count(pathJ, ".")
		minDot := Min(nI, nJ)
		ssI, ssJ := strings.Split(pathI, "."), strings.Split(pathJ, ".")
	NEXT:
		for i := 0; i < minDot+1; i++ {
			si, sj := ssI[i], ssJ[i]
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
		return nI < nJ
	})
	return paths
}

// ScanLine :
func ScanLine(r io.Reader, f func(line string) (bool, string)) (string, error) {
	lines := []string{}
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		if ok, line := f(scanner.Text()); ok {
			lines = append(lines, line)
		}
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}
	return strings.Join(lines, "\n"), nil
}

func ScanLineEx(r io.Reader, nAbove, nBelow int, junkLine string, f func(line string, cache []string) (bool, string)) (string, error) {
	if nAbove < 0 || nBelow < 0 {
		return "", fmt.Errorf("both nAbove: [%v] and nBelow [%v] cannot be less than 0", nAbove, nBelow)
	}
	var (
		scanner = bufio.NewScanner(r)
		lines   = []string{}
		rtLines = []string{}
	)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}
	for cache := range IterCache(lines, nAbove, nBelow, junkLine) {
		if ok, line := f(cache.Elem, cache.Cache); ok {
			rtLines = append(rtLines, line)
		}
	}
	return strings.Join(rtLines, "\n"), nil
}

// StrLineScan :
func StrLineScan(str string, f func(line string) (bool, string)) (string, error) {
	return ScanLine(strings.NewReader(str), f)
}

// StrLineScanEx :
func StrLineScanEx(str string, nAbove, nBelow int, junkLine string, f func(line string, cache []string) (bool, string)) (string, error) {
	return ScanLineEx(strings.NewReader(str), nAbove, nBelow, junkLine, f)
}
