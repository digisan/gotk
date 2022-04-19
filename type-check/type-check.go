package typecheck

import (
	"encoding/csv"
	"encoding/json"
	"encoding/xml"
	"math"
	"reflect"
	"strconv"
	"strings"

	"github.com/digisan/gotk/iter"
)

// IsXML : Check str is valid XML
func IsXML(str string) bool {
	return xml.Unmarshal([]byte(str), new(any)) == nil
}

// IsJSON : Check str is valid JSON
func IsJSON(str string) bool {
	return json.Unmarshal([]byte(str), new(any)) == nil
}

// IsCSV : Check str is valid CSV
func IsCSV(str string) bool {
	records, err := csv.NewReader(strings.NewReader(str)).ReadAll()
	return err == nil && len(records) > 0
}

// IsNumeric : Check str is valid numeric style
func IsNumeric(str string) bool {
	_, err := strconv.ParseFloat(str, 64)
	return err == nil
}

// IsContInts : check ints is continuous int slice
func IsContInts(ints ...int) (ok bool, minIfOk int, maxIfOk int) {
	if len(ints) == 0 {
		return false, math.MinInt32, math.MaxInt32
	}
	if len(ints) == 1 {
		return true, ints[0], ints[0]
	}

	s, e := ints[0], ints[len(ints)-1]
	if s < e {
		return reflect.DeepEqual(iter.Iter2Slc(s, e+1), ints), s, e
	}
	return reflect.DeepEqual(iter.Iter2Slc(s, e-1), ints), e, s
}

// nil-pointer could be non-nil any
func IsInterfaceNil(i any) bool {
	return i == nil || reflect.ValueOf(i).IsNil()
}
