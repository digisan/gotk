package filedir

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestChangeFileName(t *testing.T) {
	fPath := "./a/b/c/d.txt"
	fmt.Println(fPath)
	fmt.Println(ChangeFileName(fPath, "prefix-", "-crop"))

	fmt.Println("------------------------------------")

	fPath = "./a/b/c/d"
	fmt.Println(fPath)
	fmt.Println(ChangeFileName(fPath, "prefix-", "-crop"))
}

func TestParent(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "OK",
			args: args{
				path: "./io.go",
			},
			want: "io",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Parent(tt.args.path); got != tt.want {
				t.Errorf("Parent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGrandParent(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "OK",
			args: args{
				path: "./io.go",
			},
			want: "gotk",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GrandParent(tt.args.path); got != tt.want {
				t.Errorf("GrandParent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAncestorList(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
		{
			name: "OK",
			args: args{path: "./io.go"},
			want: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AncestorList(tt.args.path); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AncestorList() = %v, want %v", got, tt.want)
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

func TestFilesAllExist(t *testing.T) {
	type args struct {
		paths []string
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
				paths: []string{
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
			if got := AllFilesExist(tt.args.paths...); got != tt.want {
				t.Errorf("FilesExist() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDirsAllExist(t *testing.T) {
	type args struct {
		paths []string
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
				paths: []string{
					"~/Desktop/",
					"~/Desktop/gotk",
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AllDirsExist(tt.args.paths...); got != tt.want {
				t.Errorf("DirsExist() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAllExistAsWhole(t *testing.T) {
	type args struct {
		paths []string
	}
	tests := []struct {
		name  string
		args  args
		want  bool
		want1 string
	}{
		// TODO: Add test cases.
		{
			args: args{
				paths: []string{
					"./a.txt",
					"./c/c.txt",
					"./b.txt",
					"./d.txt",
				},
			},
			want:  false,
			want1: "./a.txt",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, got1 := AllExistAsWhole(tt.args.paths...); got != tt.want || got1 != tt.want1 {
				t.Errorf("AllExistAsWhole() = %v %v, want %v %v", got, got1, tt.want, tt.want1)
			}
		})
	}
}

func TestFileIsEmpty(t *testing.T) {
	type args struct {
		fName string
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
			args:    args{fName: "./io.go"},
			want:    false,
			wantErr: false,
		},
		{
			name:    "OK",
			args:    args{fName: "./io.txt"},
			want:    true,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IsFileEmpty(tt.args.fName)
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
			got, err := IsDirEmpty(tt.args.dirname)
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

func TestFileDir(t *testing.T) {
	type args struct {
		dirname   string
		recursive bool
		exctypes  []string
	}
	tests := []struct {
		name      string
		args      args
		wantFiles []string
		wantDirs  []string
		wantErr   bool
	}{
		// TODO: Add test cases.
		{
			name: "OK",
			args: args{
				dirname:   "../",
				recursive: true,
				exctypes:  []string{"test"},
			},
			wantFiles: []string{}, // 3,
			wantDirs:  []string{}, // 9,
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotFiles, gotDirs, err := WalkFileDir(tt.args.dirname, tt.args.recursive, tt.args.exctypes...)
			fmt.Println(err)
			fmt.Println("----------")
			fmt.Println(gotFiles)
			fmt.Println("----------")
			fmt.Println(gotDirs)
		})
	}
}

func TestFileExists(t *testing.T) {
	type args struct {
		fName string
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
				fName: "./io.go",
			},
			want: true,
		},
		{
			name: "OK",
			args: args{
				fName: "~/Desktop/gotk/io/io.go",
			},
			want: true,
		},
		{
			name: "OK",
			args: args{
				fName: "./io.txt",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FileExists(tt.args.fName); got != tt.want {
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

func TestMergeDir(t *testing.T) {
	type args struct {
		destDir string
		move    bool
		srcDirs []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "OK",
			args: args{
				destDir: "~/Desktop/mergetest0",
				move:    false,
				srcDirs: []string{"./", "../io"},
			},
		},
		{
			name: "OK",
			args: args{
				destDir: "./merge4",
				move:    true,
				srcDirs: []string{"./merge1", "./merge3"},
			},
		},
	}

	oc := func(existing, incoming []byte) (bool, []byte) {
		return true, append(existing, incoming...)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := MergeDir(tt.args.destDir, tt.args.move, oc, tt.args.srcDirs...); (err != nil) != tt.wantErr {
				t.Errorf("MergeDir() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDotExt(t *testing.T) {
	type args struct {
		ext string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "OK",
			args: args{
				ext: ".txt",
			},
			want: ".txt",
		},
		{
			name: "OK",
			args: args{
				ext: "txt",
			},
			want: ".txt",
		},
		{
			name: "OK",
			args: args{
				ext: "  	",
			},
			want: "",
		},
		{
			name: "OK",
			args: args{
				ext: "",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DotExt(tt.args.ext); got != tt.want {
				t.Errorf("DotExt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemove(t *testing.T) {
	type args struct {
		path       string
		rmEmptyDir bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "OK",
			args: args{
				path:       "~/Desktop/test/a.txt",
				rmEmptyDir: true,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Remove(tt.args.path, tt.args.rmEmptyDir); (err != nil) != tt.wantErr {
				t.Errorf("Remove() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDirSize(t *testing.T) {
	sz, err := DirSize("../../", "k")
	fmt.Printf("%f %v", sz, err)
}

func TestSelfHash(t *testing.T) {
	fmt.Println(SelfMD5())
	fmt.Println(SelfSHA1())
	fmt.Println(SelfSHA256())
}

func TestChangePath(t *testing.T) {
	type args struct {
		strict   bool
		fPath    string
		newTail  string
		fromLast int
		keepExt  bool
		cp       bool
		mv       bool
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				fPath:    "a/b/c/d.txt",
				newTail:  "D",
				fromLast: 1,
				keepExt:  true,
			},
			want: "a/b/c/D.txt",
		},
		{
			name: "",
			args: args{
				fPath:    "a/b/c/d.txt",
				newTail:  "D",
				fromLast: 1,
				keepExt:  false,
			},
			want: "a/b/c/D",
		},
		{
			name: "",
			args: args{
				strict:   true,
				cp:       true,
				fPath:    "./test/v.go",
				newTail:  "C/D",
				fromLast: 2,
				keepExt:  true,
			},
			want: "./C/D.go",
		},
		{
			name: "",
			args: args{
				fPath:    "a/b/c/d.txt",
				newTail:  "D.json",
				fromLast: 2,
				keepExt:  false,
			},
			want: "a/b/D.json",
		},
		{
			name: "",
			args: args{
				fPath:    "/d.txt",
				newTail:  "D",
				fromLast: 1,
				keepExt:  true,
			},
			want: "/D.txt",
		},
		{
			name: "",
			args: args{
				fPath:    "/d.txt",
				newTail:  "D.json",
				fromLast: 1,
				keepExt:  false,
			},
			want: "/D.json",
		},
		{
			name: "move",
			args: args{
				fPath:    "./var.go",
				newTail:  "test/v",
				fromLast: 1,
				keepExt:  true,
				cp:       true,
				mv:       true,
			},
			want: "./test/v.go",
		},
		{
			name: "strict",
			args: args{
				strict:   true,
				fPath:    "./Var.go",
				newTail:  "test/v",
				fromLast: 1,
				keepExt:  true,
				cp:       true,
				mv:       true,
			},
			want: "./test/v.go",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ChangeFilePath(tt.args.strict, tt.args.fPath, tt.args.newTail, tt.args.fromLast, tt.args.keepExt, tt.args.cp, tt.args.mv); got != tt.want {
				t.Errorf("ChangeFilePath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRmFiles(t *testing.T) {
	RmFilesIn("~/Desktop/test", true, true, "jpg", "json")
}

func TestMustWriteFile(t *testing.T) {
	type args struct {
		fName string
		data  []byte
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "OK",
			args: args{
				fName: "~/Desktop/testout.txt",
				data:  []byte("TestMustWriteFile111"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MustWriteFile(tt.args.fName, tt.args.data)
		})
	}
}

func TestMustAppendFile(t *testing.T) {
	type args struct {
		fName   string
		data    []byte
		newline bool
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "OK",
			args: args{
				fName:   "~/Desktop/testout.txt",
				data:    []byte("TestMustAppendFile"),
				newline: true,
			},
		},
		{
			name: "OK",
			args: args{
				fName:   "./testout1.txt",
				data:    []byte("TestMustAppendFile"),
				newline: false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MustAppendFile(tt.args.fName, tt.args.data, tt.args.newline)
		})
	}
}

func TestMustCreateDirs(t *testing.T) {
	type args struct {
		dir []string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "OK",
			args: args{
				dir: []string{
					"./test/must/create/dir/",
					"./test/must1/create1/dir1/",
					"./test/must2/create2/dir2/",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MustCreateDirs(tt.args.dir...)
		})
	}
}

func TestFileLineScan(t *testing.T) {
	type args struct {
		fPath   string
		f       func(line string) (bool, string)
		outFile string
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
				fPath: "../generics/T1.template",
				f: func(line string) (bool, string) {
					return true, strings.ReplaceAll(strings.ReplaceAll(line, "Xxx", "Int"), "xxx", "int")
				},
				outFile: "./out.txt",
			},
			want:    "",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FileLineScan(tt.args.fPath, tt.args.f, tt.args.outFile)
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
