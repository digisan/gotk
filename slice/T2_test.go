package slice

import "testing"

func TestT2FuncGen(t *testing.T) {
	type args struct {
		Tx     string
		Ty     string
		pkgdir string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		// {
		// 	name: "OK",
		// 	args: args{
		// 		Tx:     "int",
		// 		Ty:     "string",
		// 		pkgdir: "./",
		// 	},
		// },
		// {
		// 	name: "OK",
		// 	args: args{
		// 		Tx:     "int",
		// 		Ty:     "int",
		// 		pkgdir: "./",
		// 	},
		// },
		// {
		// 	name: "OK",
		// 	args: args{
		// 		Tx:     "string",
		// 		Ty:     "string",
		// 		pkgdir: "./",
		// 	},
		// },
		// {
		// 	name: "OK",
		// 	args: args{
		// 		Tx:     "string",
		// 		Ty:     "int",
		// 		pkgdir: "./",
		// 	},
		// },
		// {
		// 	name: "OK",
		// 	args: args{
		// 		Tx:     "string",
		// 		Ty:     "bool",
		// 		pkgdir: "./",
		// 	},
		// },
		{
			name: "OK",
			args: args{
				Tx:     "interface{}",
				Ty:     "interface{}",
				pkgdir: "./",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			T2FuncGen(tt.args.Tx, tt.args.Ty, tt.args.pkgdir)
		})
	}
}
