package tu8

import (
	"reflect"
	"testing"
)

func TestIn(t *testing.T) {
	type args struct {
		e   byte
		arr []byte
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
				e:   12,
				arr: []uint8{11, 12, 13},
			},
			want: true,
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
		e   byte
		arr []byte
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
			if got := NotIn(tt.args.e, tt.args.arr...); got != tt.want {
				t.Errorf("NotIn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIdxOf(t *testing.T) {
	type args struct {
		e   byte
		arr []byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
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
		e   byte
		arr []byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
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
		arr []byte
	}
	tests := []struct {
		name    string
		args    args
		wantSet []byte
	}{
		// TODO: Add test cases.
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
		setA []byte
		setB []byte
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
			if got := Superset(tt.args.setA, tt.args.setB); got != tt.want {
				t.Errorf("Superset() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubset(t *testing.T) {
	type args struct {
		setA []byte
		setB []byte
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
			if got := Subset(tt.args.setA, tt.args.setB); got != tt.want {
				t.Errorf("Subset() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEqual(t *testing.T) {
	type args struct {
		setA []byte
		setB []byte
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
			if got := Equal(tt.args.setA, tt.args.setB); got != tt.want {
				t.Errorf("Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_union(t *testing.T) {
	type args struct {
		setA []byte
		setB []byte
	}
	tests := []struct {
		name    string
		args    args
		wantSet []byte
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
		sets [][]byte
	}
	tests := []struct {
		name    string
		args    args
		wantSet []byte
	}{
		// TODO: Add test cases.
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
		setA []byte
		setB []byte
	}
	tests := []struct {
		name    string
		args    args
		wantSet []byte
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
		sets [][]byte
	}
	tests := []struct {
		name    string
		args    args
		wantSet []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSet := Intersect(tt.args.sets...); !reflect.DeepEqual(gotSet, tt.wantSet) {
				t.Errorf("Intersect() = %v, want %v", gotSet, tt.wantSet)
			}
		})
	}
}
