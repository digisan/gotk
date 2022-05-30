package strs

import (
	"reflect"
	"testing"
	// . "github.com/digisan/go-generics/v2"
)

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
		{
			name: "",
			args: args{
				s:   "abcdeabcdabe  ab",
				sub: "ab",
			},
			wantStarts: []int{0, 5, 9, 14},
			wantEnds:   []int{2, 7, 11, 16},
		},
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
		{
			name: "",
			args: args{
				s:      "abcdefghijklmn",
				ranges: [][2]int{{2, 4}, {7, 9}, {12, 13}},
				ns:     []string{"XXX", "YYY", "ZZZZZZ"},
			},
			want: "abXXXefgYYYjklZZZZZZn",
		},
		{
			name: "",
			args: args{
				s:      "abcdefghijklmn",
				ranges: [][2]int{{2, 5}, {6, 9}},
				ns:     []string{"", "", "ZZZZZZ"},
			},
			want: "abfjklmn",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RangeReplace(tt.args.s, tt.args.ranges, tt.args.ns); got != tt.want {
				t.Errorf("RangeReplace() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSplitPart(t *testing.T) {
	{
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
			{
				args: args{
					s:   "abc^def^ghi",
					sep: "^",
					idx: 0,
				},
				want: "abc",
			},
			{
				args: args{
					s:   "abc^def^",
					sep: "^",
					idx: 2,
				},
				want: "",
			},
			{
				args: args{
					s:   "^def^ghi",
					sep: "^",
					idx: 0,
				},
				want: "",
			},
			{
				args: args{
					s:   "abcdefghi",
					sep: "^",
					idx: 0,
				},
				want: "abcdefghi",
			},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := SplitPart(tt.args.s, tt.args.sep, tt.args.idx); got != tt.want {
					t.Errorf("SplitPart() = %v, want %v", got, tt.want)
				}
			})
		}
	}

	{
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
			{
				args: args{
					s:   "abc^def^ghi",
					sep: "^",
					idx: 1,
				},
				want: "ghi",
			},
			{
				args: args{
					s:   "abc^def^",
					sep: "^",
					idx: 1,
				},
				want: "",
			},
			{
				args: args{
					s:   "^def^ghi",
					sep: "^",
					idx: 3,
				},
				want: "",
			},
			{
				args: args{
					s:   "abcdefghi",
					sep: "^",
					idx: 1,
				},
				want: "abcdefghi",
			},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := SplitPartFromLast(tt.args.s, tt.args.sep, tt.args.idx); got != tt.want {
					t.Errorf("SplitPart() = %v, want %v", got, tt.want)
				}
			})
		}
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
		{
			args: args{
				s: "abc\ndef",
			},
			want: []string{"abc", "def"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SplitLn(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SplitLn() = %v, want %v", got, tt.want)
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
		{
			args: args{
				s:    "abcd#efgh#i",
				mark: "#",
			},
			want: "abcd#efgh",
		},
		{
			args: args{
				s:    "abcd#",
				mark: "#",
			},
			want: "abcd",
		},
		{
			args: args{
				s:    "abcd",
				mark: "#",
			},
			want: "abcd",
		},
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
		{
			args: args{
				s:    "abcd#efgh#i",
				mark: "#",
			},
			want: "i",
		},
		{
			args: args{
				s:    "abcd#",
				mark: "#",
			},
			want: "",
		},
		{
			args: args{
				s:    "abcd",
				mark: "#",
			},
			want: "abcd",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TrimHeadToLast(tt.args.s, tt.args.mark); got != tt.want {
				t.Errorf("TrimHeadToLast() = %v, want %v", got, tt.want)
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
		{
			args: args{
				s:      "abc",
				length: 5,
			},
			want: "abc",
		},
		{
			args: args{
				s:      "abc",
				length: 2,
			},
			want: "ab",
		},
		{
			args: args{
				s:      "abc",
				length: 3,
			},
			want: "abc",
		},
		{
			args: args{
				s:      "",
				length: 3,
			},
			want: "",
		},
		{
			args: args{
				s:      "",
				length: 0,
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Maxlen(tt.args.s, tt.args.length); got != tt.want {
				t.Errorf("Maxlen() = %v, want %v", got, tt.want)
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
		{
			args: args{
				s:   "1-2-3-4-5",
				sep: "-",
				idx: 0,
			},
			want: 1.0,
		},
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
		{
			args: args{
				s:   "t-true-FALSE-false-True",
				sep: "-",
				idx: 0,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SplitPartToBool(tt.args.s, tt.args.sep, tt.args.idx); got != tt.want {
				t.Errorf("SplitPartToBool() = %v, want %v", got, tt.want)
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
		{
			args: args{
				s:   "1-2-3-4-5",
				sep: "-",
				idx: 1,
			},
			want: 5.0,
		},
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
		{
			args: args{
				s:   "T-f-true-FALSE-F-0",
				sep: "-",
				idx: 2,
			},
			want: false,
		},
		{
			args: args{
				s:   "T-f-true-FALSE-F-0",
				sep: "-",
				idx: 6,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SplitPartFromLastToBool(tt.args.s, tt.args.sep, tt.args.idx); got != tt.want {
				t.Errorf("SplitPartFromLastToBool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReplaceFirstOnAnyOf(t *testing.T) {
	type args struct {
		s    string
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
				s:    "abc AdefD def abc",
				new:  "xxx",
				aims: []string{"def", "ab"},
			},
			want: "xxxc AdefD def abc",
		},
		{
			args: args{
				s:    "abc def AabcD def",
				new:  "xxx",
				aims: []string{"a", "abc"},
			},
			want: "xxx def AabcD def",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReplaceFirstOnAnyOf(tt.args.s, tt.args.new, tt.args.aims...); got != tt.want {
				t.Errorf("ReplaceFirstOnAnyOf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReplaceAllOnAnyOf(t *testing.T) {
	type args struct {
		s    string
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
				s:    "abc def AabcD def",
				new:  "xxx",
				aims: []string{"abc", "def"},
			},
			want: "xxx xxx AxxxD xxx",
		},
		{
			args: args{
				s:    "abc def AabD def yyyy",
				new:  "xxx",
				aims: []string{"a", "abc", "def", "yy"},
			},
			want: "xxxbc xxx AxxxbD xxx xxxxxx",
		},
		{
			args: args{
				s:    "1234567",
				aims: []string{"12", "34"},
				new:  "1234",
			},
			want: "12341234567",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReplaceAllOnAnyOf(tt.args.s, tt.args.new, tt.args.aims...); got != tt.want {
				t.Errorf("ReplaceAllOnAnyOf() = %v, want %v", got, tt.want)
			}
		})
	}
}
