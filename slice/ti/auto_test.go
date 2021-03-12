package ti

import (
	"reflect"
	"testing"
)

func TestIn(t *testing.T) {
	type args struct {
		e   int
		arr []int
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
				e:   1,
				arr: []int{1, 2, 3},
			},
			want: true,
		},
		{
			name: "OK",
			args: args{
				e:   1,
				arr: []int{2, 3},
			},
			want: false,
		},
		{
			name: "OK",
			args: args{
				e:   1,
				arr: []int{},
			},
			want: false,
		},
		{
			name: "OK",
			args: args{
				e:   1,
				arr: nil,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := In(tt.args.e, tt.args.arr...); got != tt.want {
				t.Errorf("In() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNotIn(t *testing.T) {
	type args struct {
		e   int
		arr []int
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
				e:   1,
				arr: []int{1, 2, 3},
			},
			want: false,
		},
		{
			name: "OK",
			args: args{
				e:   1,
				arr: []int{2, 3},
			},
			want: true,
		},
		{
			name: "OK",
			args: args{
				e:   1,
				arr: []int{},
			},
			want: true,
		},
		{
			name: "OK",
			args: args{
				e:   1,
				arr: nil,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NotIn(tt.args.e, tt.args.arr...); got != tt.want {
				t.Errorf("NotIn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIdxOf(t *testing.T) {
	type args struct {
		e   int
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "OK",
			args: args{
				e:   1,
				arr: []int{1, 2, 3},
			},
			want: 0,
		},
		{
			name: "OK",
			args: args{
				e:   1,
				arr: []int{2, 3},
			},
			want: -1,
		},
		{
			name: "OK",
			args: args{
				e:   1,
				arr: []int{},
			},
			want: -1,
		},
		{
			name: "OK",
			args: args{
				e:   1,
				arr: nil,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IdxOf(tt.args.e, tt.args.arr...); got != tt.want {
				t.Errorf("IdxOf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLastIdxOf(t *testing.T) {
	type args struct {
		e   int
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "OK",
			args: args{
				e:   1,
				arr: []int{1, 2, 1, 3},
			},
			want: 2,
		},
		{
			name: "OK",
			args: args{
				e:   1,
				arr: []int{2, 3},
			},
			want: -1,
		},
		{
			name: "OK",
			args: args{
				e:   1,
				arr: []int{},
			},
			want: -1,
		},
		{
			name: "OK",
			args: args{
				e:   1,
				arr: nil,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LastIdxOf(tt.args.e, tt.args.arr...); got != tt.want {
				t.Errorf("LastIdxOf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMkSet(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name    string
		args    args
		wantSet []int
	}{
		// TODO: Add test cases.
		{
			name: "OK",
			args: args{
				arr: []int{1, 2, 3, 4, 3, 2, 1},
			},
			wantSet: []int{1, 2, 3, 4},
		},
		{
			name: "OK",
			args: args{
				arr: []int{},
			},
			wantSet: []int{},
		},
		{
			name: "OK",
			args: args{
				arr: nil,
			},
			wantSet: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSet := MkSet(tt.args.arr...); !reflect.DeepEqual(gotSet, tt.wantSet) {
				t.Errorf("MkSet() = %v, want %v", gotSet, tt.wantSet)
			}
		})
	}
}

func TestSuperset(t *testing.T) {
	type args struct {
		setA []int
		setB []int
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
				setA: []int{1, 2, 3, 4},
				setB: []int{2, 4},
			},
			want: true,
		},
		{
			name: "OK",
			args: args{
				setA: []int{1, 2, 3, 4},
				setB: []int{2, 4, 1, 3},
			},
			want: false,
		},
		{
			name: "OK",
			args: args{
				setA: []int{1},
				setB: []int{2, 4, 1, 3},
			},
			want: false,
		},
		{
			name: "OK",
			args: args{
				setA: []int{1},
				setB: []int{},
			},
			want: true,
		},
		{
			name: "OK",
			args: args{
				setA: []int{1},
				setB: nil,
			},
			want: true,
		},
		{
			name: "OK",
			args: args{
				setA: nil,
				setB: nil,
			},
			want: false,
		},
		{
			name: "OK",
			args: args{
				setA: nil,
				setB: []int{},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Superset(tt.args.setA, tt.args.setB); got != tt.want {
				t.Errorf("Superset() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubset(t *testing.T) {
	type args struct {
		setA []int
		setB []int
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
				setA: []int{1, 2, 3, 4},
				setB: []int{2, 4},
			},
			want: false,
		},
		{
			name: "OK",
			args: args{
				setA: []int{1, 2, 3, 4},
				setB: []int{2, 4, 1, 3},
			},
			want: false,
		},
		{
			name: "OK",
			args: args{
				setA: []int{1},
				setB: []int{2, 4, 1, 3},
			},
			want: true,
		},
		{
			name: "OK",
			args: args{
				setA: []int{1},
				setB: []int{},
			},
			want: false,
		},
		{
			name: "OK",
			args: args{
				setA: []int{1},
				setB: nil,
			},
			want: false,
		},
		{
			name: "OK",
			args: args{
				setA: nil,
				setB: nil,
			},
			want: false,
		},
		{
			name: "OK",
			args: args{
				setA: nil,
				setB: []int{},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Subset(tt.args.setA, tt.args.setB); got != tt.want {
				t.Errorf("Subset() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEqual(t *testing.T) {
	type args struct {
		setA []int
		setB []int
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
				setA: []int{1, 2, 3, 4},
				setB: []int{2, 4},
			},
			want: false,
		},
		{
			name: "OK",
			args: args{
				setA: []int{1, 2, 3, 4},
				setB: []int{2, 4, 1, 3},
			},
			want: true,
		},
		{
			name: "OK",
			args: args{
				setA: []int{1},
				setB: []int{2, 4, 1, 3},
			},
			want: false,
		},
		{
			name: "OK",
			args: args{
				setA: []int{},
				setB: []int{},
			},
			want: true,
		},
		{
			name: "OK",
			args: args{
				setA: []int{1},
				setB: nil,
			},
			want: false,
		},
		{
			name: "OK",
			args: args{
				setA: nil,
				setB: nil,
			},
			want: true,
		},
		{
			name: "OK",
			args: args{
				setA: nil,
				setB: []int{},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Equal(tt.args.setA, tt.args.setB); got != tt.want {
				t.Errorf("Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_union(t *testing.T) {
	type args struct {
		setA []int
		setB []int
	}
	tests := []struct {
		name    string
		args    args
		wantSet []int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSet := union(tt.args.setA, tt.args.setB); !reflect.DeepEqual(gotSet, tt.wantSet) {
				t.Errorf("union() = %v, want %v", gotSet, tt.wantSet)
			}
		})
	}
}

func TestUnion(t *testing.T) {
	type args struct {
		sets [][]int
	}
	tests := []struct {
		name    string
		args    args
		wantSet []int
	}{
		// TODO: Add test cases.
		{
			name: "OK",
			args: args{
				sets: [][]int{{}},
			},
			wantSet: []int{},
		},
		{
			name: "OK",
			args: args{
				sets: [][]int{},
			},
			wantSet: nil,
		},
		{
			name: "OK",
			args: args{
				sets: [][]int{{}, nil, {1, 2, 3}, {2, 3, 4}, {3, 4, 5}},
			},
			wantSet: []int{1, 2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSet := Union(tt.args.sets...); !reflect.DeepEqual(gotSet, tt.wantSet) {
				t.Errorf("Union() = %v, want %v", gotSet, tt.wantSet)
			}
		})
	}
}

func Test_intersect(t *testing.T) {
	type args struct {
		setA []int
		setB []int
	}
	tests := []struct {
		name    string
		args    args
		wantSet []int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSet := intersect(tt.args.setA, tt.args.setB); !reflect.DeepEqual(gotSet, tt.wantSet) {
				t.Errorf("intersect() = %v, want %v", gotSet, tt.wantSet)
			}
		})
	}
}

func TestIntersect(t *testing.T) {
	type args struct {
		sets [][]int
	}
	tests := []struct {
		name    string
		args    args
		wantSet []int
	}{
		// TODO: Add test cases.
		{
			name: "OK",
			args: args{
				sets: [][]int{{}},
			},
			wantSet: []int{},
		},
		{
			name: "OK",
			args: args{
				sets: [][]int{},
			},
			wantSet: nil,
		},
		{
			name: "OK",
			args: args{
				sets: [][]int{{}, nil, {1, 2, 3}, {2, 3, 4}, {3, 4, 5}},
			},
			wantSet: nil,
		},
		{
			name: "OK",
			args: args{
				sets: [][]int{{}, {1, 2, 3}, {2, 3, 4}, {3, 4, 5}},
			},
			wantSet: []int{},
		},
		{
			name: "OK",
			args: args{
				sets: [][]int{{1, 2, 3}, {2, 3, 4}, {3, 4, 5}},
			},
			wantSet: []int{3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSet := Intersect(tt.args.sets...); !reflect.DeepEqual(gotSet, tt.wantSet) {
				t.Errorf("Intersect() = %v, want %v", gotSet, tt.wantSet)
			}
		})
	}
}

func TestFilterModify(t *testing.T) {
	type args struct {
		arr      []int
		filter   func(i int, e int) bool
		modifier func(i int, e int) int
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
				arr:      []int{10, 20, 30, 40, 50, 60, 70},
				filter:   func(i int, e int) bool { return e > 30 },
				modifier: func(i int, e int) int { return e + 1 },
			},
			wantR: []int{41, 51, 61, 71},
		},
		{
			name: "OK",
			args: args{
				arr:      []int{10, 20, 30, 40, 50, 60, 70},
				filter:   nil,
				modifier: func(i int, e int) int { return e + 1 },
			},
			wantR: []int{11, 21, 31, 41, 51, 61, 71},
		},
		{
			name: "OK",
			args: args{
				arr:      []int{10, 20, 30, 40, 50, 60, 70},
				filter:   func(i int, e int) bool { return e > 30 },
				modifier: nil,
			},
			wantR: []int{40, 50, 60, 70},
		},
		{
			name: "OK",
			args: args{
				arr:      []int{10, 20, 30, 40, 50, 60, 70},
				filter:   nil,
				modifier: nil,
			},
			wantR: []int{10, 20, 30, 40, 50, 60, 70},
		},
		{
			name: "OK",
			args: args{
				arr:      nil,
				filter:   nil,
				modifier: nil,
			},
			wantR: nil,
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
