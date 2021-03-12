package strint

import (
	"reflect"
	"testing"
)

func TestFilterModify(t *testing.T) {
	type args struct {
		arr      []string
		filter   func(i int, e string) bool
		modifier func(i int, e string) int
	}
	tests := []struct {
		name  string
		args  args
		wantR []int
	}{
		// TODO: Add test cases.
		{
			name: "OK",
			args: args{
				arr:      []string{"1", "20", "300", "4000", "50000", "600000", "7000000"},
				filter:   func(i int, e string) bool { return len(e) > 3 },
				modifier: func(i int, e string) int { return i },
			},
			wantR: []int{3, 4, 5, 6},
		},
		{
			name: "OK",
			args: args{
				arr:      []string{"1", "20", "300", "4000", "50000", "600000", "7000000"},
				filter:   nil,
				modifier: func(i int, e string) int { return i },
			},
			wantR: []int{0, 1, 2, 3, 4, 5, 6},
		},
		{
			name: "SHOULD PANIC",
			args: args{
				arr:      []string{"1", "20", "300", "4000", "50000", "600000", "7000000"},
				filter:   func(i int, e string) bool { return len(e) > 3 },
				modifier: nil,
			},
			wantR: []int{0, 1, 2, 3, 4, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotR := FM(tt.args.arr, tt.args.filter, tt.args.modifier); !reflect.DeepEqual(gotR, tt.wantR) {
				t.Errorf("FM() = %v, want %v", gotR, tt.wantR)
			}
		})
	}
}
