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

// XxxGen :
func XxxGen(T, outgofile string) {
	_, err := io.FileLineScan("./xxx.txt", func(line string) (bool, string) {
		if fp, ok := mTypFP[T]; ok {
			return true, sReplaceAll(sReplaceAll(line, "Xxx", fp), "xxx", T)
		}
		panic(T + " is not supported for xxx")
	}, outgofile)
	if err != nil {
		panic(err.Error())
	}
}
