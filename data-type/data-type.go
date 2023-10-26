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
	. "github.com/digisan/go-generics/v2"
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

// IsHTML: check str is valid HTML
func IsHTML(data []byte) bool {
	z := html.NewTokenizer(bytes.NewReader(data))
	for {
		switch z.Next() {
		case html.ErrorToken:
			return false
		case html.StartTagToken, html.EndTagToken:
			return true
		}
	}
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
