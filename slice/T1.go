package slice

import (
	"os"

	"github.com/digisan/gotk/io"
)

// type -> package name
var mTypPkg = map[string]string{
	"int":         "int",
	"int8":        "i8",
	"int16":       "i16",
	"int32":       "i32",
	"int64":       "i64",
	"float32":     "f32",
	"float64":     "f64",
	"bool":        "bool",
	"uint":        "uint",
	"uint8":       "u8",
	"uint16":      "u16",
	"uint32":      "u32",
	"uint64":      "u64",
	"string":      "str",
	"interface{}": "obj",
}

// T1FuncGen :
func T1FuncGen(Tx, pkgdir string) {

	pkgname, ok := mTypPkg[Tx]
	if !ok {
		panic(Tx + " is not supported for T<xxx>")
	}

	pkgdir = sTrimSuffix(pkgdir, "/") + "/"
	if !io.DirExists(pkgdir) {
		if err := os.MkdirAll(pkgdir, os.ModePerm); err != nil {
			panic(err.Error())
		}
	}

	outgofile := pkgdir + pkgname + "/auto.go"

	if !io.FileExists(outgofile) || io.FileIsEmpty(outgofile) {
		io.MustWriteFile(outgofile, []byte("package "+pkgname))
	}

	src, err := io.FileLineScan("./T1.template", func(line string) (bool, string) {
		line = sTrimRight(line, " \t")
		return true, sReplaceAll(line, "xxx", Tx)
	}, "")

	if err != nil {
		panic(err.Error())
	}

	io.MustAppendFile(outgofile, []byte(""), true)
	io.MustAppendFile(outgofile, []byte(src), true)
}
