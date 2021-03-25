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

// IsXML :
func IsXML(str string) bool {
	return xml.Unmarshal([]byte(str), new(interface{})) == nil
}

// IsJSON :
func IsJSON(str string) bool {
	return json.Unmarshal([]byte(str), new(interface{})) == nil
}

// IsNumeric :
func IsNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}
