package tsi

import (
	"reflect"
	"testing"
)

func TestFilterMap(t *testing.T) {
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

func TestMap2KVs(t *testing.T) {
	type args struct {
		m          map[string]int
		less4key   func(i string, j string) bool
		less4value func(i int, j int) bool
	}
	tests := []struct {
		name       string
		args       args
		wantKeys   []string
		wantValues []int
	}{
		// TODO: Add test cases.
		{
			name: "OK",
			args: args{
				m:          map[string]int{"1": 1, "12": 2, "123": 3, "234": 2},
				less4key:   func(i string, j string) bool { return len(i) > len(j) },
				less4value: func(i, j int) bool { return i < j },
			},
			wantKeys:   []string{"1", "234", "12", "123"},
			wantValues: []int{1, 2, 2, 3},
		},
		{
			name: "OK",
			args: args{
				m:        map[string]int{"1": 1, "12": 2, "123": 3, "234": 2},
				less4key: func(i string, j string) bool { return len(i) > len(j) },
			},
			wantKeys:   []string{"123", "234", "12", "1"},
			wantValues: []int{3, 2, 2, 1},
		},
		{
			name: "OK",
			args: args{
				m:          map[string]int{"1": 1, "12": 2, "123": 3, "234": 2},
				less4value: func(i, j int) bool { return i < j },
			},
			wantKeys:   []string{"1", "12", "234", "123"},
			wantValues: []int{1, 2, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotKeys, gotValues := Map2KVs(tt.args.m, tt.args.less4key, tt.args.less4value)
			if !reflect.DeepEqual(gotKeys, tt.wantKeys) {
				t.Errorf("Map2KVs() gotKeys = %v, want %v", gotKeys, tt.wantKeys)
			}
			if !reflect.DeepEqual(gotValues, tt.wantValues) {
				t.Errorf("Map2KVs() gotValues = %v, want %v", gotValues, tt.wantValues)
			}
		})
	}
}
