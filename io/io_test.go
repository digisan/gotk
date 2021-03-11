package io

import (
	"testing"
)

func TestFileLineScan(t *testing.T) {
	type args struct {
		filepath string
		f        func(line string) (bool, string)
		outfile  string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "OK",
			args: args{
				filepath: "../slice/xxx.txt",
				f: func(line string) (bool, string) {
					return true, sReplaceAll(sReplaceAll(line, "Xxx", "Int"), "xxx", "int")
				},
				outfile: "../slice/int.go",
			},
			want:    "",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FileLineScan(tt.args.filepath, tt.args.f, tt.args.outfile)
			if (err != nil) != tt.wantErr {
				t.Errorf("FileLineScan() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("FileLineScan() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFileExists(t *testing.T) {
	type args struct {
		filename string
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
				filename: "./io.go",
			},
			want: true,
		},
		{
			name: "OK",
			args: args{
				filename: "./io.txt",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FileExists(tt.args.filename); got != tt.want {
				t.Errorf("FileExists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDirExists(t *testing.T) {
	type args struct {
		dirname string
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
				dirname: "./io",
			},
			want: false,
		},
		{
			name: "OK",
			args: args{
				dirname: "../io",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DirExists(tt.args.dirname); got != tt.want {
				t.Errorf("DirExists() = %v, want %v", got, tt.want)
			}
		})
	}
}
