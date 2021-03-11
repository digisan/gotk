package slice

import "testing"

func TestXxxGen(t *testing.T) {
	type args struct {
		T         string
		outgofile string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "OK",
			args: args{
				T:         "string",
				outgofile: "./string.go",
			},
		},
		{
			name: "OK",
			args: args{
				T:         "int",
				outgofile: "./int.go",
			},
		},
		{
			name: "OK",
			args: args{
				T:         "interface{}",
				outgofile: "./object.go",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			XxxGen(tt.args.T, tt.args.outgofile)
		})
	}
}
