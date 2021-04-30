package slice

import (
	"os"

	"github.com/digisan/gotk/io"
)

// T2FuncGen :
func T2FuncGen(Tx, Ty, pkgdir string) {

	pkgname1, ok := mTypPkg[Tx]
	if !ok {
		panic(Tx + " is not supported for T<xxx>")
	}
	pkgname2, ok := mTypPkg[Ty]
	if !ok {
		panic(Ty + " is not supported for T<yyy>")
	}
	if Tx == Ty {
		pkgname2 = ""
	} else {
		pkgname2 = sTrimPrefix(pkgname2, "t")
	}

	pkgdir = sTrimSuffix(pkgdir, "/") + "/"
	if !io.DirExists(pkgdir) {
		if err := os.MkdirAll(pkgdir, os.ModePerm); err != nil {
			panic(err.Error())
		}
	}

	pkgname := pkgname1 + pkgname2
	outgofile := pkgdir + pkgname + "/auto.go"

	// if !io.FileExists(outgofile) || io.FileIsEmpty(outgofile) {
	// 	io.MustWriteFile(outgofile, []byte("package "+pkgname))
	// }
	if empty, err := io.FileIsEmpty(outgofile); err != nil || empty {
		io.MustWriteFile(outgofile, []byte("package "+pkgname))
	}

	flagXeqY, flagXneY := true, true

	src, err := io.FileLineScan("./T2.template", func(line string) (bool, string) {
		line = sTrimRight(line, " \t")

		switch {
		case sHasSuffix(line, `[S@x==y]`):
			flagXeqY, flagXneY = true, false
		case sHasSuffix(line, `[E@x==y]`):
			flagXeqY, flagXneY = true, true
		case sHasSuffix(line, `[S@x!=y]`):
			flagXeqY, flagXneY = false, true
		case sHasSuffix(line, `[E@x!=y]`):
			flagXeqY, flagXneY = true, true
		}

		if sHasPrefix(sTrimLeft(line, " \t"), "// [") {
			return false, ""
		}

		line = sReplaceAll(line, "yyy", Ty)
		line = sReplaceAll(line, "xxx", Tx)

		if (Tx == Ty && flagXeqY) || (Tx != Ty && flagXneY) {
			return true, line
		}
		return false, ""

	}, "")

	if err != nil {
		panic(err.Error())
	}

	io.MustAppendFile(outgofile, []byte(""), true)
	io.MustAppendFile(outgofile, []byte(src), true)
}
