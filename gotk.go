package gotk

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"math"
	"reflect"
	"strconv"
	"time"

	"github.com/digisan/gotk/iter"
)

// TrackTime : defer TrackTime(time.Now())
func TrackTime(start time.Time) {
	elapsed := time.Since(start)
	fmt.Printf("Took %s\n", elapsed)
}

// IsXML : Check str is valid XML
func IsXML(str string) bool {
	return xml.Unmarshal([]byte(str), new(interface{})) == nil
}

// IsJSON : Check str is valid JSON
func IsJSON(str string) bool {
	return json.Unmarshal([]byte(str), new(interface{})) == nil
}

// IsNumeric : Check str is valid numeric style
func IsNumeric(str string) bool {
	_, err := strconv.ParseFloat(str, 64)
	return err == nil
}

// IsContInts : check ints is continuous int slice
func IsContInts(ints []int) (ok bool, minIfOk int, maxIfOk int) {
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
