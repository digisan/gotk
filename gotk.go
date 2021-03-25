package gotk

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"strconv"
	"time"
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
