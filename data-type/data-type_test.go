package datatype

import (
	"fmt"
	"os"
	"testing"
)

func TestTypeCheck(t *testing.T) {
	fmt.Println(IsJSON([]byte("")))
	fmt.Println(IsXML([]byte("")))
	fmt.Println(IsCSV([]byte("")))
	fmt.Println(IsTOML([]byte("")))
	fmt.Println(IsYAML([]byte("doe: \"a deer, a female deer\"")))
}

func TestTxtType(t *testing.T) {
	str := "abc"
	fmt.Println(DataType(str))

	data := []byte{88, 89}
	fmt.Println(DataType(data))

	f, err := os.Open("./data/sample.csv")
	if err == nil {
		fmt.Println("sample type:", DataType(f))
	}
	csv := make([]byte, 100)
	n, err := f.Read(csv)
	fmt.Println(n, err, string(csv))
}

func TestListTypes(t *testing.T) {
	fmt.Println(SupportedDataTypes())
}
