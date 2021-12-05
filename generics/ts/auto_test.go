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
		setC []string
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
				setA: []string{"a", "b", "c"},
				setB: []string{"b", "a", "c"},
				setC: []string{"c", "b", "a"},
			},
			want: true,
		},
		{
			name: "OK",
			args: args{
				setA: []string{"a", "b", "c"},
				setB: []string{"b", "a", "c"},
				setC: []string{"c", "b", "d"},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Equal(tt.args.setA, tt.args.setB, tt.args.setC); got != tt.want {
				t.Errorf("Equal() = %v, want %v", got, tt.want)
			}
			fmt.Println(tt.args.setA, tt.args.setB, tt.args.setC)
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

func Test_minus(t *testing.T) {
	type args struct {
		setA []string
		setB []string
	}
	tests := []struct {
		name    string
		args    args
		wantSet []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSet := minus(tt.args.setA, tt.args.setB); !reflect.DeepEqual(gotSet, tt.wantSet) {
				t.Errorf("minus() = %v, want %v", gotSet, tt.wantSet)
			}
		})
	}
}

func TestMinus(t *testing.T) {
	type args struct {
		setA      []string
		setOthers [][]string
	}
	tests := []struct {
		name    string
		args    args
		wantSet []string
	}{
		// TODO: Add test cases.
		{
			name: "OK",
			args: args{
				setA:      []string{"a", "b", "c"},
				setOthers: [][]string{{"b", "B"}, {"c", "C"}},
			},
			wantSet: []string{"a"},
		},
		{
			name: "OK",
			args: args{
				setA:      []string{"a", "b", "c"},
				setOthers: [][]string{{"B"}, {"c", "C"}, {"a"}},
			},
			wantSet: []string{"b"},
		},
		{
			name: "OK",
			args: args{
				setA:      []string{"a", "b", "c"},
				setOthers: [][]string{{"B", "b"}, {"c", "C"}, {"a"}},
			},
			wantSet: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSet := Minus(tt.args.setA, tt.args.setOthers...); !reflect.DeepEqual(gotSet, tt.wantSet) {
				t.Errorf("Minus() = %v, want %v", gotSet, tt.wantSet)
			}
		})
	}
}

func TestDelEle(t *testing.T) {
	type args struct {
		arr []string
		i   int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "DelEle",
			args: args{
				arr: []string{"1", "2", "3", "4", "5"},
				i:   0,
			},
		},
		{
			name: "DelEle",
			args: args{
				arr: []string{"1", "2", "3", "4", "5"},
				i:   1,
			},
		},
		{
			name: "DelEle",
			args: args{
				arr: []string{"1", "2", "3", "4", "5"},
				i:   4,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DelEle(&tt.args.arr, tt.args.i)
			// DelEleUnordered(&tt.args.arr, tt.args.i)
			fmt.Println(tt.args.arr)
		})
	}
}

func TestMapMerge(t *testing.T) {
	type args struct {
		ms []map[string]string
	}
	tests := []struct {
		name string
		args args
		want map[string][]string
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				ms: []map[string]string{
					{"id_1": "val_1"},
					{"id_2": "val_2", "id_1": "val_1"},
					{"id_1": "val_3"},
				},
			},
			want: map[string][]string{
				"id_1": {"val_1", "val_3"},
				"id_2": {"val_2"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MapMerge(tt.args.ms...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MapMerge() = %v, want %v", got, tt.want)
			}
		})
	}
}
