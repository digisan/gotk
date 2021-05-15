package env

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestChunk2Map(t *testing.T) {
	type args struct {
		filepath  string
		markstart string
		markend   string
		sep       string
		env       bool
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
				filepath:  "./variables.md",
				markstart: "```export",
				markend:   "```",
				sep:       "=",
				env:       true,
			},
			want: map[string]string{"$A": "abcghi", "$B": "defabcghighi", "$C": "ghi"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Chunk2Map(tt.args.filepath, tt.args.markstart, tt.args.markend, tt.args.sep, tt.args.env)
		})
	}

	fmt.Println()
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotValstr := EnvValued(tt.args.str, tt.args.m); gotValstr != tt.wantValstr {
				t.Errorf("EnvValued() = %v, want %v", gotValstr, tt.wantValstr)
			}
		})
	}
}
