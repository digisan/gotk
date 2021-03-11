package slice

import "github.com/digisan/gotk/io"

// T2FuncGen :
func T2FuncGen(Tx, Ty, pkgname, outgofile string) {

	if !io.FileExists(outgofile) || io.FileIsEmpty(outgofile) {
		io.MustWriteFile(outgofile, []byte("package "+pkgname+"\n"))
	}

	flagXeqY, flagXneY := true, true

	src, err := io.FileLineScan("./T2.template", func(line string) (bool, string) {
		line = sTrimRight(line, " \t")

		fpx, ok := mTypFP[Tx]
		if !ok {
			panic(Tx + " is incorrect type")
		}

		fsy, ok := mTypFP[Ty]
		if !ok {
			panic(Ty + " is incorrect type")
		}
		if Tx == Ty {
			fsy = ""
		}

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

		line = sReplaceAll(sReplaceAll(line, "Yyys", fsy), "yyy", Ty)
		line = sReplaceAll(sReplaceAll(line, "Xxx", fpx), "xxx", Tx)

		if (Tx == Ty && flagXeqY) || (Tx != Ty && flagXneY) {
			return true, line
		}
		return false, ""

	}, "")
	if err != nil {
		panic(err.Error())
	}
	io.MustAppendFile(outgofile, []byte(src), true)
	io.MustAppendFile(outgofile, []byte{}, true)

	return
}
