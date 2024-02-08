package datatype

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"encoding/xml"
	"io"
	"os"
	"strings"

	"github.com/BurntSushi/toml"
	. "github.com/digisan/go-generics"
	"golang.org/x/net/html"
	"gopkg.in/yaml.v3"
)

const (
	// Data Type
	JSON    = "json"
	XML     = "xml"
	CSV     = "csv"
	TOML    = "toml"
	YAML    = "yaml"
	HTML    = "html"
	TEXT    = "text"
	UNKNOWN = "unknown"
)

// IsXML: Check str is valid XML
func IsXML(data []byte) bool {
	return xml.Unmarshal(data, new(any)) == nil
}

// IsJSON: Check str is valid JSON
func IsJSON(data []byte) bool {
	return json.Unmarshal(data, new(any)) == nil
}

// IsCSV: Check str is valid CSV
func IsCSV(data []byte) bool {
	records, err := csv.NewReader(bytes.NewReader(data)).ReadAll()
	return err == nil && len(records) > 0
}

// IsTOML: check str is valid TOML
func IsTOML(data []byte) bool {
	m := map[string]any{}
	_, err := toml.Decode(ConstBytesToStr(data), &m)
	return err == nil && len(m) > 0
}

// IsYAML: check str is valid YAML
func IsYAML(data []byte) bool {
	m := map[string]any{}
	err := yaml.Unmarshal(data, &m)
	return err == nil && len(m) > 0
}

// IsValidHTML: check str is valid HTML. if plainTextIncl is true, plain text is valid HTML.
func IsValidHTML(data []byte, plainTextIncl bool) bool {

	isVoidElement := func(tagName string) bool {
		// List of known void (self-closing) HTML elements
		voidElements := map[string]bool{
			"area": true, "base": true, "br": true, "col": true, "embed": true,
			"hr": true, "img": true, "input": true, "link": true, "meta": true,
			"param": true, "source": true, "track": true, "wbr": true, "keygen": true,
		}
		return voidElements[tagName]
	}

	tokenizer := html.NewTokenizer(bytes.NewReader(data))
	openTags := make(map[string]int)
	flagTag := false
	for {
		switch tokenizer.Next() {
		case html.ErrorToken:
			if tokenizer.Err().Error() == "EOF" {
				// Make sure all previously opened tags are closed
				for _, count := range openTags {
					if count > 0 {
						return false
					}
				}
				return IF(plainTextIncl, true, flagTag)
			}
			return false

		case html.StartTagToken:
			flagTag = true
			token := tokenizer.Token()
			tagName := token.Data
			if !isVoidElement(tagName) {
				openTags[tagName]++
			}

		case html.EndTagToken:
			token := tokenizer.Token()
			tagName := token.Data
			if !isVoidElement(tagName) {
				openTags[tagName]--
			}
		}
	}
}

func IsHTML(data []byte) bool {
	return IsValidHTML(data, false)
}

func SupportedTypes() []string {
	return []string{JSON, XML, CSV, TOML, YAML, HTML, TEXT, UNKNOWN}
}

func IsSupportedType(tType string) bool {
	tType = strings.ToLower(tType)
	return In(tType, SupportedTypes()...)
}

type Block interface {
	string | []byte | *os.File
}

func dataType(data []byte) string {
	switch {
	case IsJSON(data):
		return JSON
	case IsTOML(data):
		return TOML
	case IsXML(data):
		return XML
	case IsYAML(data):
		return YAML
	case IsCSV(data):
		return CSV
	case IsHTML(data):
		return HTML
	default:
		return UNKNOWN
	}
}

func DataType[T Block](data T) string {
	var in any = data
	switch TypeOf(data) {
	case "string":
		return dataType(StrToConstBytes(in.(string)))
	case "[]uint8":
		return dataType(in.([]byte))
	case "*os.File":
		f := in.(*os.File)
		defer f.Seek(0, io.SeekStart)
		bytes, err := io.ReadAll(f)
		if err != nil {
			return err.Error()
		}
		return dataType(bytes)
	}
	panic("shouldn't be here")
}
