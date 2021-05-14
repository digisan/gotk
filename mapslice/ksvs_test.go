package mapslice

import (
	"reflect"
	"testing"
)

func TestMap2Slices(t *testing.T) {

	M := map[string]string{"ab": "12", "abc": "1", "a": "123"}

	type args struct {
		m         map[string]string
		orderType string
	}
	tests := []struct {
		name       string
		args       args
		wantKeys   []string
		wantValues []string
	}{
		// TODO: Add test cases.
		{
			name:       "OK",
			args:       args{m: M, orderType: "KL-DESC"},
			wantKeys:   []string{"abc", "ab", "a"},
			wantValues: []string{"1", "12", "123"},
		},
		{
			name:       "OK",
			args:       args{m: M, orderType: "KL-ASC"},
			wantKeys:   []string{"a", "ab", "abc"},
			wantValues: []string{"123", "12", "1"},
		},
		{
			name:       "OK",
			args:       args{m: M, orderType: "VL-ASC"},
			wantKeys:   []string{"abc", "ab", "a"},
			wantValues: []string{"1", "12", "123"},
		},
		{
			name:       "OK",
			args:       args{m: M, orderType: "VL-DESC"},
			wantKeys:   []string{"a", "ab", "abc"},
			wantValues: []string{"123", "12", "1"},
		},
		{
			name:       "OK1",
			args:       args{m: M, orderType: ""},
			wantKeys:   []string{"ab", "abc", "a"},
			wantValues: []string{"12", "1", "123"},
		},
		{
			name:       "OK",
			args:       args{m: map[string]string{}, orderType: ""},
			wantKeys:   []string{},
			wantValues: []string{},
		},
		{
			name:       "OK",
			args:       args{m: nil, orderType: ""},
			wantKeys:   nil,
			wantValues: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotKeys, gotValues := KsVs2Slc(tt.args.m, tt.args.orderType)
			if !reflect.DeepEqual(gotKeys, tt.wantKeys) {
				t.Errorf("KsVs2Slc() gotKeys = %v, want %v", gotKeys, tt.wantKeys)
			}
			if !reflect.DeepEqual(gotValues, tt.wantValues) {
				t.Errorf("KsVs2Slc() gotValues = %v, want %v", gotValues, tt.wantValues)
			}
		})
	}
}
