package env

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestSplitN(t *testing.T) {
	s := "abc"
	ss := strings.SplitN(s, "#", 2)
	fmt.Println(len(ss), len(ss[0]), ss)

	s = "abc#"
	ss = strings.SplitN(s, "#", 2)
	fmt.Println(len(ss), len(ss[0]), len(ss[1]), ss)

	s = "abc#1"
	ss = strings.SplitN(s, "#", 2)
	fmt.Println(len(ss), len(ss[0]), len(ss[1]), ss)
}

func TestChunk2Map(t *testing.T) {
	type args struct {
		filepath     string
		markstart    string
		markend      string
		sep          string
		env          bool
		val4path2abs bool
	}
	tests := []struct {
		name string
		args args
		want map[string]string
	}{
		// TODO: Add test cases.
		{
			name: "OK",
			args: args{
				filepath:     "./variables.sh",
				markstart:    "###export",
				markend:      "###",
				sep:          "=",
				env:          true,
				val4path2abs: true,
			},
			want: map[string]string{"$A": "~/Desktop/", "$B": "defabcghighi", "$C": "ghi"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := Chunk2Map(tt.args.filepath, tt.args.markstart, tt.args.markend, tt.args.sep, tt.args.env, tt.args.val4path2abs)
			fmt.Println(m)
		})
	}

	fmt.Println(`----------------------------------`)
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		fmt.Println(pair[0], pair[1])
	}

}

func TestEnvValued(t *testing.T) {
	type args struct {
		str string
		m   map[string]string
	}
	tests := []struct {
		name       string
		args       args
		wantValstr string
	}{
		// TODO: Add test cases.
		{
			name: "OK",
			args: args{
				str: "$TEST_ROOT/path/a/b/c",
				m:   map[string]string{"TEST_ROOT": "root"},
			},
			wantValstr: "root/path/a/b/c",
		},
		{
			name: "OK",
			args: args{
				str: "$TEST_ROOT/$TEST_PATH/a/b/c",
				m:   map[string]string{"TEST_ROOT": "root", "TEST_PATH": "path"},
			},
			wantValstr: "root/path/a/b/c",
		},
		{
			name: "OK",
			args: args{
				str: "$TEST_ROOT/$TEST_PATH/a/b/c",
				m:   nil,
			},
			wantValstr: "root/path/a/b/c",
		},
	}
	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if i == 2 {
				os.Setenv("TEST_ROOT", "root")
				os.Setenv("TEST_PATH", "path")
			}

			if gotValstr := EnvValued(tt.args.str, tt.args.m); gotValstr != tt.wantValstr {
				t.Errorf("EnvValued() = %v, want %v", gotValstr, tt.wantValstr)
			}
		})
	}
}
