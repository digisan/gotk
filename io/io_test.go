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
				filepath: "../generics/T1.template",
				f: func(line string) (bool, string) {
					return true, sReplaceAll(sReplaceAll(line, "Xxx", "Int"), "xxx", "int")
				},
				outfile: "./out.txt",
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

func TestMustWriteFile(t *testing.T) {
	type args struct {
		filename string
		data     []byte
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "OK",
			args: args{
				filename: "~/Desktop/testout.txt",
				data:     []byte("TestMustWriteFile111"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MustWriteFile(tt.args.filename, tt.args.data)
		})
	}
}

func TestMustAppendFile(t *testing.T) {
	type args struct {
		filename string
		data     []byte
		newline  bool
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "OK",
			args: args{
				filename: "~/Desktop/testout.txt",
				data:     []byte("TestMustAppendFile"),
				newline:  true,
			},
		},
		{
			name: "OK",
			args: args{
				filename: "./testout1.txt",
				data:     []byte("TestMustAppendFile"),
				newline:  false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MustAppendFile(tt.args.filename, tt.args.data, tt.args.newline)
		})
	}
}

func TestMustCreateDir(t *testing.T) {
	type args struct {
		dir string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "OK",
			args: args{
				dir: "./test/must/create/dir/",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MustCreateDir(tt.args.dir)
		})
	}
}

