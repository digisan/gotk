package strs

import (
	"fmt"
	"reflect"
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
				ignoreCase: true,
				ignoreSpace: false,
				s:           "A",
				group:       []string{"a ", "b", "c", "c"},
			},
			want: false,
		},
		{
			name: "",
			args: args{
				ignoreCase: false,
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
		ignoreCase bool
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
				ignoreCase: true,
				ignoreSpace: true,
				s:           "a",
				group:       []string{"A", "B", "C"},
			},
			want: false,
		},
		{
			name: "",
			args: args{
				ignoreCase: false,
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

func TestReplaceFirstOnAnyOf(t *testing.T) {
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReplaceFirstOnAnyOf(tt.args.str, tt.args.new, tt.args.aims...); got != tt.want {
				t.Errorf("ReplaceFirstOnAnyOf() = %v, want %v", got, tt.want)
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReplaceAllOnAnyOf(tt.args.str, tt.args.new, tt.args.aims...); got != tt.want {
				t.Errorf("ReplaceAllOnAnyOf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSplitPart(t *testing.T) {
	type args struct {
		s   string
		sep string
		idx int
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
			if got := SplitPart(tt.args.s, tt.args.sep, tt.args.idx); got != tt.want {
				t.Errorf("SplitPart() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSplitPartToNum(t *testing.T) {
	type args struct {
		s   string
		sep string
		idx int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SplitPartToNum(tt.args.s, tt.args.sep, tt.args.idx); got != tt.want {
				t.Errorf("SplitPartToNum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSplitPartToBool(t *testing.T) {
	type args struct {
		s   string
		sep string
		idx int
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
			if got := SplitPartToBool(tt.args.s, tt.args.sep, tt.args.idx); got != tt.want {
				t.Errorf("SplitPartToBool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSplitPartFromLast(t *testing.T) {
	type args struct {
		s   string
		sep string
		idx int
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
			if got := SplitPartFromLast(tt.args.s, tt.args.sep, tt.args.idx); got != tt.want {
				t.Errorf("SplitPartFromLast() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSplitPartFromLastToNum(t *testing.T) {
	type args struct {
		s   string
		sep string
		idx int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SplitPartFromLastToNum(tt.args.s, tt.args.sep, tt.args.idx); got != tt.want {
				t.Errorf("SplitPartFromLastToNum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSplitPartFromLastToBool(t *testing.T) {
	type args struct {
		s   string
		sep string
		idx int
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
			if got := SplitPartFromLastToBool(tt.args.s, tt.args.sep, tt.args.idx); got != tt.want {
				t.Errorf("SplitPartFromLastToBool() = %v, want %v", got, tt.want)
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
		htmlstr string
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
			if gotRt := HtmlTextContent(tt.args.htmlstr); !reflect.DeepEqual(gotRt, tt.wantRt) {
				t.Errorf("HtmlTextContent() = %v, want %v", gotRt, tt.wantRt)
			}
		})
	}
}
