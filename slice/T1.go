package slice

import (
	"os"

	"github.com/digisan/gotk/io"
)

// type -> package name
var mTypPkg = map[string]string{
	"int":         "ti",
	"int8":        "ti8",
	"int16":       "ti16",
	"int32":       "ti32",
	"rune":        "ti32",
	"int64":       "ti64",
	"float32":     "tf32",
	"float64":     "tf64",
	"bool":        "tb",
	"uint":        "tu",
	"uint8":       "tu8",
	"byte":        "tu8",
	"uint16":      "tu16",
	"uint32":      "tu32",
	"uint64":      "tu64",
	"complex64":   "tc64",
	"complex128":  "tc128",
	"string":      "ts",
	"interface{}": "to",
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

	// if !io.FileExists(outgofile) || io.FileIsEmpty(outgofile) {
	// 	io.MustWriteFile(outgofile, []byte("package "+pkgname))
	// }
	if empty, err := io.FileIsEmpty(outgofile); err != nil || empty {
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
