package gotk

import (
	"testing"
	"time"
)

func TestTrackTime(t *testing.T) {
	type args struct {
		start time.Time
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			TrackTime(tt.args.start)
		})
	}
}

func TestIsXML(t *testing.T) {
	type args struct {
		str string
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
			if got := IsXML(tt.args.str); got != tt.want {
				t.Errorf("IsXML() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsJSON(t *testing.T) {
	type args struct {
		str string
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
			if got := IsJSON(tt.args.str); got != tt.want {
				t.Errorf("IsJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsNumeric(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "OK",
			args: args{"123"},
			want: true,
		},
		{
			name: "OK",
			args: args{".123"},
			want: true,
		},
		{
			name: "OK",
			args: args{"a123"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsNumeric(tt.args.s); got != tt.want {
				t.Errorf("IsNumeric() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsContInts(t *testing.T) {
	type args struct {
		ints []int
	}
	tests := []struct {
		name    string
		args    args
		wantOk  bool
		wantMin int
		wantMax int
	}{
		// TODO: Add test cases.
		{
			name:    "OK",
			args:    args{[]int{1, 2, 3, 4, 5}},
			wantOk:  true,
			wantMin: 1,
			wantMax: 5,
		},
		{
			name:    "OK",
			args:    args{[]int{1, 2, 3, 4, 6}},
			wantOk:  false,
			wantMin: 1,
			wantMax: 6,
		},
		{
			name:    "OK",
			args:    args{[]int{5, 4, 3, 2, 1}},
			wantOk:  true,
			wantMin: 1,
			wantMax: 5,
		},
		{
			name:    "OK",
			args:    args{[]int{6, 4, 3, 2, 1}},
			wantOk:  false,
			wantMin: 1,
			wantMax: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOk, gotMin, gotMax := IsContInts(tt.args.ints)
			if gotOk != tt.wantOk {
				t.Errorf("IsContInts() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
			if gotMin != tt.wantMin {
				t.Errorf("IsContInts() gotMin = %v, want %v", gotMin, tt.wantMin)
			}
			if gotMax != tt.wantMax {
				t.Errorf("IsContInts() gotMax = %v, want %v", gotMax, tt.wantMax)
			}
		})
	}
}

func TestIsCSV(t *testing.T) {
	type args struct {
		str string
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
				str: "a,b,c,d\n1,2,3,4",
			},
			want: true,
		},
		{
			name: "OK1",
			args: args{
				str: "a,b,c,d,e\n1,2,3,4",
			},
			want: false,
		},
		{
			name: "OK1",
			args: args{
				str: "a,b,c\n1,2,3,4",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsCSV(tt.args.str); got != tt.want {
				t.Errorf("IsCSV() = %v, want %v", got, tt.want)
			}
		})
	}
}
