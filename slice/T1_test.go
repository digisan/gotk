package slice

import "testing"

func TestT1FuncGen(t *testing.T) {
	type args struct {
		Tx     string
		pkgdir string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "OK",
			args: args{
				Tx:     "string",
				pkgdir: "./",
			},
		},
		{
			name: "OK",
			args: args{
				Tx:     "int",
				pkgdir: "./",
			},
		},
		{
			name: "OK",
			args: args{
				Tx:     "byte",
				pkgdir: "./",
			},
		},
		{
			name: "OK",
			args: args{
				Tx:     "rune",
				pkgdir: "./",
			},
		},
		{
			name: "OK",
			args: args{
				Tx:     "interface{}",
				pkgdir: "./",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			T1FuncGen(tt.args.Tx, tt.args.pkgdir)
		})
	}
}
