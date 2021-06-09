package ts

import (
	"fmt"
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

func TestEqual(t *testing.T) {
	type args struct {
		setA []string
		setB []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "OK",
			args: args{
				setA: []string{"a", "b"},
				setB: []string{"b", "a"},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Equal(tt.args.setA, tt.args.setB); got != tt.want {
				t.Errorf("Equal() = %v, want %v", got, tt.want)
			}
			fmt.Println(tt.args.setA, tt.args.setB)
		})
	}

}

func TestReverse(t *testing.T) {
	type args struct {
		arr []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "OK",
			args: args{
				arr: []string{"a", "b", "c"},
			},
			want: []string{"c", "b", "a"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Reverse(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
