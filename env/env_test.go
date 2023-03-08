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

func TestChunkToMap(t *testing.T) {
	type args struct {
		fPath       string
		markStart   string
		markEnd     string
		sep         string
		env         bool
		pathVal2abs bool
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
				fPath:       "./variables.sh",
				markStart:   "###export",
				markEnd:     "###",
				sep:         "=",
				env:         true,
				pathVal2abs: true,
			},
			want: map[string]string{"$A": "~/Desktop/", "$B": "defabcghighi", "$C": "ghi"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := ChunkToMap(tt.args.fPath, tt.args.markStart, tt.args.markEnd, tt.args.sep, tt.args.env, tt.args.pathVal2abs)
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
		wantValStr string
	}{
		// TODO: Add test cases.
		{
			name: "OK",
			args: args{
				str: "$TEST_ROOT/path/a/b/c",
				m:   map[string]string{"TEST_ROOT": "root"},
			},
			wantValStr: "root/path/a/b/c",
		},
		{
			name: "OK",
			args: args{
				str: "$TEST_ROOT/$TEST_PATH/a/b/c",
				m:   map[string]string{"TEST_ROOT": "root", "TEST_PATH": "path"},
			},
			wantValStr: "root/path/a/b/c",
		},
		{
			name: "OK",
			args: args{
				str: "$TEST_ROOT/$TEST_PATH/a/b/c",
				m:   nil,
			},
			wantValStr: "root/path/a/b/c",
		},
	}
	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if i == 2 {
				os.Setenv("TEST_ROOT", "root")
				os.Setenv("TEST_PATH", "path")
			}

			if gotValStr := EnvValued(tt.args.str, tt.args.m); gotValStr != tt.wantValStr {
				t.Errorf("EnvValued() = %v, want %v", gotValStr, tt.wantValStr)
			}
		})
	}
}
