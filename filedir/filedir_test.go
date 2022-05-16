package filedir

import (
	"fmt"
	"reflect"
	"testing"
	// . "github.com/digisan/go-generics/v2"
)

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

func TestMergeDir(t *testing.T) {
	type args struct {
		destdir string
		move    bool
		srcdirs []string
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
				destdir: "~/Desktop/mergetest0",
				move:    false,
				srcdirs: []string{"./", "../io"},
			},
		},
		{
			name: "OK",
			args: args{
				destdir: "./merge4",
				move:    true,
				srcdirs: []string{"./merge1", "./merge3"},
			},
		},
	}

	oc := func(existing, incoming []byte) (bool, []byte) {
		return true, append(existing, incoming...)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := MergeDir(tt.args.destdir, tt.args.move, oc, tt.args.srcdirs...); (err != nil) != tt.wantErr {
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
		path string
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
				path: "~/Desktop/testout.txt",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Remove(tt.args.path); (err != nil) != tt.wantErr {
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
		fpath    string
		newtail  string
		fromlast int
		keepext  bool
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
				fpath:    "a/b/c/d.txt",
				newtail:  "D",
				fromlast: 1,
				keepext:  true,
			},
			want: "a/b/c/D.txt",
		},
		{
			name: "",
			args: args{
				fpath:    "a/b/c/d.txt",
				newtail:  "D",
				fromlast: 1,
				keepext:  false,
			},
			want: "a/b/c/D",
		},
		{
			name: "",
			args: args{
				strict:   true,
				cp:       true,
				fpath:    "./test/v.go",
				newtail:  "C/D",
				fromlast: 2,
				keepext:  true,
			},
			want: "./C/D.go",
		},
		{
			name: "",
			args: args{
				fpath:    "a/b/c/d.txt",
				newtail:  "D.json",
				fromlast: 2,
				keepext:  false,
			},
			want: "a/b/D.json",
		},
		{
			name: "",
			args: args{
				fpath:    "/d.txt",
				newtail:  "D",
				fromlast: 1,
				keepext:  true,
			},
			want: "/D.txt",
		},
		{
			name: "",
			args: args{
				fpath:    "/d.txt",
				newtail:  "D.json",
				fromlast: 1,
				keepext:  false,
			},
			want: "/D.json",
		},
		{
			name: "move",
			args: args{
				fpath:    "./var.go",
				newtail:  "test/v",
				fromlast: 1,
				keepext:  true,
				cp:       true,
				mv:       true,
			},
			want: "./test/v.go",
		},
		// {
		// 	name: "strict",
		// 	args: args{
		// 		strict: true,
		// 		fpath:    "./Var.go",
		// 		newtail:  "test/v",
		// 		fromlast: 1,
		// 		keepext:  true,
		// 		cp:       true,
		// 		mv:       true,
		// 	},
		// 	want: "./test/v.go",
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ChangeFilePath(tt.args.strict, tt.args.fpath, tt.args.newtail, tt.args.fromlast, tt.args.keepext, tt.args.cp, tt.args.mv); got != tt.want {
				t.Errorf("ChangeFilePath() = %v, want %v", got, tt.want)
			}
		})
	}
}
