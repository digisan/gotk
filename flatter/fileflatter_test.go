package flatter

import (
	"fmt"
	"os"
	"testing"

	dt "github.com/digisan/gotk/data-type"
)

func TestFlatContent(t *testing.T) {

	f, err := os.Open("../data-for-test/sample.csv")
	if err == nil {
		fmt.Println("sample type:", dt.DataType(f))
	}
	fmt.Println()

	fm, err := FlatContent(f, false)
	if err != nil {
		fmt.Println(err)
		return
	}
	for k, v := range fm {
		fmt.Println(k, v)
	}

	fmt.Println()

	fmt.Printf("%+v\n", CsvRowsAsMaps(fm, 1))

	fmt.Println()

	fmt.Printf("%+v\n", CsvRowsAsMaps(fm, 0, 1))

	// fmt.Println()

	// fmt.Printf("%+v", CsvRowsAsMaps(fm)[1])

	// PrintFlat(fm)
}
