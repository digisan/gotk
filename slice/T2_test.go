package slice

import "testing"

func TestT2Gen(t *testing.T) {
	type args struct {
		Tx        string
		Ty        string
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
				Tx:        "int",
				Ty:        "string",
				pkgname:   "slice",
				outgofile: "auto.go",
			},
		},
		{
			name: "OK",
			args: args{
				Tx:        "string",
				Ty:        "string",
				pkgname:   "slice",
				outgofile: "auto.go",
			},
		},
		{
			name: "OK",
			args: args{
				Tx:        "string",
				Ty:        "int",
				pkgname:   "slice",
				outgofile: "auto.go",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			T2FuncGen(tt.args.Tx, tt.args.Ty, tt.args.pkgname, tt.args.outgofile)
		})
	}
}
