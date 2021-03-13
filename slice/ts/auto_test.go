package ts

import (
	"reflect"
	"testing"
)

func TestReorder(t *testing.T) {
	type args struct {
		arr     []string
		indices []int
	}
	tests := []struct {
		name       string
		args       args
		wantOrders []string
	}{
		// TODO: Add test cases.
		{
			name: "OK",
			args: args{
				arr:     []string{"1", "2", "3", "4", "5"},
				indices: []int{4, 3, 0},
			},
			wantOrders: []string{"5", "4", "1"},
		},
		{
			name: "OK",
			args: args{
				arr:     []string{"1", "2", "3", "4", "5"},
				indices: []int{3},
			},
			wantOrders: []string{"4"},
		},
		{
			name: "OK",
			args: args{
				arr:     []string{},
				indices: []int{3},
			},
			wantOrders: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOrders := Reorder(tt.args.arr, tt.args.indices); !reflect.DeepEqual(gotOrders, tt.wantOrders) {
				t.Errorf("Reorder() = %v, want %v", gotOrders, tt.wantOrders)
			}
		})
	}
}
