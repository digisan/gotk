package strings

import (
	"reflect"
	"testing"
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
