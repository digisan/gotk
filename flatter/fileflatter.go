package flatter

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/BurntSushi/toml"
	. "github.com/digisan/go-generics"
	dt "github.com/digisan/gotk/data-type"
	"github.com/digisan/gotk/strs"
	"gopkg.in/yaml.v3"
)

func flatContent(data []byte, half bool) (map[string]any, error) {

	var (
		object = make(map[string]any)
		err    error
	)

	switch dt.DataType(data) {
	case dt.JSON:
		if err = json.Unmarshal(data, &object); err != nil {
			return nil, err
		}

	case dt.TOML:
		if _, err = toml.Decode(string(data), &object); err != nil {
			return nil, err
		}

	case dt.YAML:
		if err = yaml.Unmarshal(data, &object); err != nil {
			return nil, err
		}

	case dt.XML:
		if err = xml.Unmarshal(data, &object); err != nil {
			return nil, err
		}

	case dt.CSV:
		records, err := csv.NewReader(bytes.NewReader(data)).ReadAll()
		if err != nil {
			return nil, err
		}
		mIColHdr := make(map[int]string)
		for i, row := range records {
			// fmt.Printf("%+v\n", row)
			switch i {
			case 0: // headers
				for iCol, hdr := range row {
					object[hdr] = []any{}
					mIColHdr[iCol] = hdr
				}
			default: // rows
				for iCol, item := range row {
					hdr := mIColHdr[iCol]
					object[hdr] = append(object[hdr].([]any), item)
				}
			}
		}

	default:
		return nil, errors.New("unknown data type in [flatContent]")
	}

	if half {
		return MapNestedToHalfFlat(object)
	}

	return MapNestedToFlat(object), nil
}

func FlatContent[T dt.Block](data T, half bool) (map[string]any, error) {
	var in any = data
	switch TypeOf(data) {
	case "string":
		return flatContent(StrToConstBytes(in.(string)), half)
	case "[]uint8":
		return flatContent(in.([]byte), half)
	case "*os.File":
		f := in.(*os.File)
		defer func() { f.Seek(0, io.SeekStart) }()
		bytes, err := io.ReadAll(f)
		if err != nil {
			return nil, err
		}
		return flatContent(bytes, half)
	}
	panic("shouldn't be here")
}

func PrintFlat(m map[string]any) {
	ks, _ := MapToKVs(m, nil, nil)
	ks = strs.SortPaths(ks...)
	fmt.Println("\n------------------------------------")
	for _, k := range ks {
		keyLen := len(k)
		keySpace := (keyLen/10 + 1) * 10
		fmtStr := fmt.Sprintf("%%-%dv %%v\n", keySpace)
		fmt.Printf(fmtStr, k+":", m[k])
	}
	fmt.Println("------------------------------------")
}

func CsvRowsAsMaps(fm map[string]any, indices ...int) (rt []map[string]any) {

	keys, _ := MapToKVs(fm, nil, nil)

	idxSuffix := func(key string) int {
		return strs.SplitPartFromLastTo[int](key, ".", 1)
	}

	fillVal := func(p, idx int) {
		rt[p] = make(map[string]any)
		for _, key := range keys {
			suffix := fmt.Sprintf(".%d", idx)
			if strings.HasSuffix(key, suffix) {
				newKey := strings.TrimSuffix(key, suffix)
				rt[p][newKey] = fm[key]
			}
		}
	}

	if len(indices) == 0 {
		suffixIndices := make([]int, 0, len(keys))
		for _, key := range keys {
			suffixIndices = append(suffixIndices, idxSuffix(key))
		}
		indices = IterToSlc(Max(suffixIndices...) + 1)
	}

	rt = make([]map[string]any, len(indices))
	for i, idx := range indices {
		fillVal(i, idx)
	}

	return
}
