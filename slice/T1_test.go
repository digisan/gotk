package slice

import "testing"

func TestXxxGen(t *testing.T) {
	type args struct {
		Tx        string
		pkgname   string
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
				Tx:        "string",
				pkgname:   "slice",
				outgofile: "./auto.go",
			},
		},
		{
			name: "OK",
			args: args{
				Tx:        "int",
				pkgname:   "slice",
				outgofile: "./auto.go",
			},
		},
		{
			name: "OK",
			args: args{
				Tx:        "interface{}",
				pkgname:   "slice",
				outgofile: "./auto.go",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			T1FuncGen(tt.args.Tx, tt.args.pkgname, tt.args.outgofile)
		})
	}
}
