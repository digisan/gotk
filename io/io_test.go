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
				filepath: "../slice/T1.template",
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

func TestFileIsEmpty(t *testing.T) {
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
			name: "FileIsEmpty",
			args: args{filename: "./io.go"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FileIsEmpty(tt.args.filename); got != tt.want {
				t.Errorf("FileIsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDirIsEmpty(t *testing.T) {
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
			name: "FileIsEmpty",
			args: args{dirname: "./"},
			want: false,
		},
		{
			name: "FileIsEmpty",
			args: args{dirname: "../net"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DirIsEmpty(tt.args.dirname); got != tt.want {
				t.Errorf("DirIsEmpty() = %v, want %v", got, tt.want)
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
				filename: "./testout.txt",
				data:     []byte("TestMustWriteFile"),
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
				filename: "./testout1.txt",
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

func TestFileDirCount(t *testing.T) {
	type args struct {
		dirname   string
		recursive bool
	}
	tests := []struct {
		name          string
		args          args
		wantFileCount int
		wantDirCount  int
		wantErr       bool
	}{
		// TODO: Add test cases.
		{
			name: "OK",
			args: args{
				dirname:   "../",
				recursive: false,
			},
			wantFileCount: 5,
			wantDirCount:  5,
			wantErr:       false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotFileCount, gotDirCount, err := FileDirCount(tt.args.dirname, tt.args.recursive)
			if (err != nil) != tt.wantErr {
				t.Errorf("FileDirCount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotFileCount != tt.wantFileCount {
				t.Errorf("FileDirCount() gotFileCount = %v, want %v", gotFileCount, tt.wantFileCount)
			}
			if gotDirCount != tt.wantDirCount {
				t.Errorf("FileDirCount() gotDirCount = %v, want %v", gotDirCount, tt.wantDirCount)
			}
		})
	}
}
