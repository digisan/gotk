package slice

import "github.com/digisan/gotk/io"

// type -> func prefix
var mTypFP = map[string]string{
	"int":         "Int",
	"int8":        "I8",
	"int16":       "I16",
	"int32":       "I32",
	"int64":       "I64",
	"float32":     "F32",
	"float64":     "F64",
	"bool":        "Bool",
	"uint":        "Uint",
	"uint8":       "U8",
	"uint16":      "U16",
	"uint32":      "U32",
	"uint64":      "U64",
	"string":      "Str",
	"interface{}": "Obj",
}

// T1FuncGen :
func T1FuncGen(Tx, pkgname, outgofile string) {

	if !io.FileExists(outgofile) || io.FileIsEmpty(outgofile) {
		io.MustWriteFile(outgofile, []byte("package "+pkgname+"\n"))
	}

	src, err := io.FileLineScan("./T1.template", func(line string) (bool, string) {
		line = sTrimRight(line, " \t")
		if fp, ok := mTypFP[Tx]; ok {
			return true, sReplaceAll(sReplaceAll(line, "Xxx", fp), "xxx", Tx)
		}
		panic(Tx + " is not supported for xxx")
	}, "")
	if err != nil {
		panic(err.Error())
	}
	io.MustAppendFile(outgofile, []byte(src), true)
	io.MustAppendFile(outgofile, []byte{}, true)
}
