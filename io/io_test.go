package io

import (
	"fmt"
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
				filename: "~/Desktop/gotk/io/io.go",
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
				dirname: "~/Desktop/gotk/io",
			},
			want: true,
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
		exctypes  []string
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
				exctypes:  []string{"go"},
			},
			wantFileCount: 4,
			wantDirCount:  5,
			wantErr:       false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotFileCount, gotDirCount, err := FileDirCount(tt.args.dirname, tt.args.recursive, tt.args.exctypes...)
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

func TestAbsPath(t *testing.T) {
	type args struct {
		path  string
		check bool
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
				path:  "~/Desktop/",
				check: true,
			},
			want:    "/home/qmiao/Desktop",
			wantErr: false,
		},
		{
			name: "OK",
			args: args{
				path:  "~/Desktop1",
				check: true,
			},
			want:    "/home/qmiao/Desktop1",
			wantErr: true,
		},
		{
			name: "OK",
			args: args{
				path:  "~/Desktop1",
				check: false,
			},
			want:    "/home/qmiao/Desktop1",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AbsPath(tt.args.path, tt.args.check)
			if (err != nil) != tt.wantErr {
				t.Errorf("AbsPath() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("AbsPath() = %v, want %v", got, tt.want)
			}
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

func TestFilesAllExist(t *testing.T) {
	type args struct {
		filenames []string
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
				filenames: []string{
					"/home/qmiao/Desktop/nats-stream.service",
					"~/Desktop/nats-stream.service",
					"../../nats-stream.service",
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FilesAllExist(tt.args.filenames); got != tt.want {
				t.Errorf("FilesExist() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDirsAllExist(t *testing.T) {
	type args struct {
		dirnames []string
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
				dirnames: []string{
					"~/Desktop/",
					"~/Desktop/gotk",
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DirsAllExist(tt.args.dirnames); got != tt.want {
				t.Errorf("DirsExist() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFileIsEmpty(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "OK",
			args:    args{filename: "./io.go"},
			want:    false,
			wantErr: false,
		},
		{
			name:    "OK",
			args:    args{filename: "./io.txt"},
			want:    true,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FileIsEmpty(tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("FileIsEmpty() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
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
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "OK",
			args:    args{dirname: "./"},
			want:    false,
			wantErr: false,
		},
		{
			name:    "OK",
			args:    args{dirname: "../ioio"},
			want:    true,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DirIsEmpty(tt.args.dirname)
			if (err != nil) != tt.wantErr {
				t.Errorf("DirIsEmpty() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DirIsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRelPath(t *testing.T) {
	type args struct {
		path  string
		check bool
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
				path:  "~/Desktop/gotk/task/task.go",
				check: false,
			},
			want:    "",
			wantErr: false,
		},
		{
			name: "OK",
			args: args{
				path:  "~/Desktop/gotk/README1.md",
				check: true,
			},
			want:    "",
			wantErr: false,
		},
		{
			name: "OK",
			args: args{
				path:  "./var.go",
				check: false,
			},
			want:    "",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RelPath(tt.args.path, tt.args.check)
			fmt.Println(got, err)
			// if (err != nil) != tt.wantErr {
			// 	t.Errorf("RelPath() error = %v, wantErr %v", err, tt.wantErr)
			// 	return
			// }
			// if got != tt.want {
			// 	t.Errorf("RelPath() = %v, want %v", got, tt.want)
			// }
		})
	}
	fmt.Println("finish")
}
