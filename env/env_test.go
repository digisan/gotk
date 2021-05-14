package env

import (
	"fmt"
	"os"
	"reflect"
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
			if got := Chunk2Map(tt.args.filepath, tt.args.markstart, tt.args.markend, tt.args.sep, tt.args.env); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Chunk2Map() = %v, want %v", got, tt.want)
			}
		})
	}

	fmt.Println()
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		fmt.Println(pair[0], pair[1])
	}

}
