package strs

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestIsIn(t *testing.T) {
	type args struct {
		ignoreCase  bool
		ignoreSpace bool
		s           string
		group       []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				ignoreCase:  true,
				ignoreSpace: false,
				s:           "A",
				group:       []string{"a ", "b", "c", "c"},
			},
			want: false,
		},
		{
			name: "",
			args: args{
				ignoreCase:  false,
				ignoreSpace: false,
				s:           "AA",
				group:       []string{"aa", "bb", "cc", ""},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsIn(tt.args.ignoreCase, tt.args.ignoreSpace, tt.args.s, tt.args.group...); got != tt.want {
				t.Errorf("IsIn() = %v, want %v", got, tt.want)
			}
		})

		fmt.Println(tt.args.s, tt.args.group)
	}
}

func TestIsNotIn(t *testing.T) {
	type args struct {
		ignoreCase  bool
		ignoreSpace bool
		s           string
		group       []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				ignoreCase:  true,
				ignoreSpace: true,
				s:           "a",
				group:       []string{"A", "B", "C"},
			},
			want: false,
		},
		{
			name: "",
			args: args{
				ignoreCase:  false,
				ignoreSpace: true,
				s:           "a",
				group:       []string{"A", "B", "C"},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsNotIn(tt.args.ignoreCase, tt.args.ignoreSpace, tt.args.s, tt.args.group...); got != tt.want {
				t.Errorf("IsNotIn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxlen(t *testing.T) {
	type args struct {
		s      string
		length int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Maxlen(tt.args.s, tt.args.length); got != tt.want {
				t.Errorf("Maxlen() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIndexAll(t *testing.T) {
	type args struct {
		s   string
		sub string
	}
	tests := []struct {
		name       string
		args       args
		wantStarts []int
		wantEnds   []int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotStarts, gotEnds := IndexAll(tt.args.s, tt.args.sub)
			if !reflect.DeepEqual(gotStarts, tt.wantStarts) {
				t.Errorf("IndexAll() gotStarts = %v, want %v", gotStarts, tt.wantStarts)
			}
			if !reflect.DeepEqual(gotEnds, tt.wantEnds) {
				t.Errorf("IndexAll() gotEnds = %v, want %v", gotEnds, tt.wantEnds)
			}
		})
	}
}

func TestIndexAllByReg(t *testing.T) {
	type args struct {
		s   string
		sub string
	}
	tests := []struct {
		name       string
		args       args
		wantStarts []int
		wantEnds   []int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotStarts, gotEnds := IndexAllByReg(tt.args.s, tt.args.sub)
			if !reflect.DeepEqual(gotStarts, tt.wantStarts) {
				t.Errorf("IndexAllByReg() gotStarts = %v, want %v", gotStarts, tt.wantStarts)
			}
			if !reflect.DeepEqual(gotEnds, tt.wantEnds) {
				t.Errorf("IndexAllByReg() gotEnds = %v, want %v", gotEnds, tt.wantEnds)
			}
		})
	}
}

func TestRangeReplace(t *testing.T) {
	type args struct {
		s      string
		ranges [][2]int
		ns     []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RangeReplace(tt.args.s, tt.args.ranges, tt.args.ns); got != tt.want {
				t.Errorf("RangeReplace() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHasAnyPrefix(t *testing.T) {
	type args struct {
		s         string
		prefixGrp []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HasAnyPrefix(tt.args.s, tt.args.prefixGrp...); got != tt.want {
				t.Errorf("HasAnyPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHasAnySuffix(t *testing.T) {
	type args struct {
		s         string
		suffixGrp []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HasAnySuffix(tt.args.s, tt.args.suffixGrp...); got != tt.want {
				t.Errorf("HasAnySuffix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainsAny(t *testing.T) {
	type args struct {
		s    string
		aims []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsAny(tt.args.s, tt.args.aims...); got != tt.want {
				t.Errorf("ContainsAny() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReplace1stOnAnyOf(t *testing.T) {
	type args struct {
		str  string
		new  string
		aims []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			args: args{
				str:  "this is a test string",
				new:  "A",
				aims: []string{"a"},
			},
			want: "this is A test string",
		},
		{
			args: args{
				str:  "this is a test string",
				new:  "A",
				aims: []string{"is", "this"},
			},
			want: "A is a test string",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Replace1stOnAnyOf(tt.args.str, tt.args.new, tt.args.aims...); got != tt.want {
				t.Errorf("Replace1stOnAnyOf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReplaceAllOnAnyOf(t *testing.T) {
	type args struct {
		str  string
		new  string
		aims []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			args: args{
				str:  "this is a test string",
				new:  "A",
				aims: []string{"a"},
			},
			want: "this is A test string",
		},
		{
			args: args{
				str:  "this is a test string",
				new:  "A",
				aims: []string{"is", "this", "str"},
			},
			want: "A A a test Aing",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReplaceAllOnAnyOf(tt.args.str, tt.args.new, tt.args.aims...); got != tt.want {
				t.Errorf("ReplaceAllOnAnyOf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTrimTailFromLast(t *testing.T) {
	type args struct {
		s    string
		mark string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TrimTailFromLast(tt.args.s, tt.args.mark); got != tt.want {
				t.Errorf("TrimTailFromLast() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTrimHeadToLast(t *testing.T) {
	type args struct {
		s    string
		mark string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TrimHeadToLast(tt.args.s, tt.args.mark); got != tt.want {
				t.Errorf("TrimHeadToLast() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSplitLn(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SplitLn(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SplitLn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHtmlTextContent(t *testing.T) {
	type args struct {
		htmlStr string
	}
	tests := []struct {
		name   string
		args   args
		wantRt []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRt := HtmlTextContent(tt.args.htmlStr); !reflect.DeepEqual(gotRt, tt.wantRt) {
				t.Errorf("HtmlTextContent() = %v, want %v", gotRt, tt.wantRt)
			}
		})
	}
}

func TestSplitPart(t *testing.T) {
	// fmt.Println(SplitPartTo[int]("abc.23.ss", ".", 2))
	// fmt.Println(SplitPartTo[int]("abc.23.ss", ".", 6))
	fmt.Println(SplitPartTo[int]("abc.23.ss", ".", 1))
}

func TestSortPaths(t *testing.T) {
	paths := SortPaths(
		"B.m.1.u",
		"A.z",
		"B.m.1.t",
		"B.m.1.t.W",
		"A.b.1",
		"B.m.2.T",
		"A.Z",
		"b.0.A",
		"B.m.3.t",
	)
	for _, path := range paths {
		fmt.Println(path)
	}
}

func TestReversePath(t *testing.T) {
	path := "B.m.1.t.W"
	fmt.Println(ReversePath(path))
}

func TestModifyOriginOrIP(t *testing.T) {

	s := `
	http://127.0.0.1:3000/api/
	https://127.0.0.1:3000/api/
	http://localhost:3001/api/
	https://localhost:3001/api/
	localhost:3002/api/
	127.0.0.1:3002/api/
	`

	// fmt.Println(ModifyOriginIP(s, "localhost", "my.example.com", "", -1, true, true, true))
	// fmt.Println(ModifyOriginIP(s, "localhost", "my.example.com", "", -1, true, true, false))
	// fmt.Println(ModifyOriginIP(s, "localhost", "test://my.example.com", "", -1, false, true, false))
	fmt.Println(ModifyOriginIP(s, "localhost", "my.example.com:1234", "", -1, true, false, true))
}

func TestScanLineEx(t *testing.T) {

	s := `http://127.0.0.1:3000/api/
https://127.0.0.1:3000/api/
http://localhost:3001/api/
https://localhost:3001/api/
localhost:3002/api/
127.0.0.1:3002/api/`

	rt, err := ScanLineEx(strings.NewReader(s), 2, 2, "******", func(line string, cache []string) (bool, string) {
		// fmt.Println(line, "--->", cache)
		return true, strings.ToUpper(line)
	})

	if err == nil {
		fmt.Println(rt)
	}
}

func TestTrimTo(t *testing.T) {

	str := "                  \"asn_hasLevel\": {"

	s1 := TrimTailFromFirst(str, "\"")
	fmt.Println(s1, len(s1))

	s2 := TrimTailFromLast(str, "\"")
	fmt.Println(s2, len(s2))

	s3 := TrimHeadToFirst(str, "\"")
	fmt.Println(s3, len(s3))

	s4 := TrimHeadToLast(str, "\"")
	fmt.Println(s4, len(s4))
}
