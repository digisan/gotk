package src

import (
	_ "embed"
	"fmt"
	"testing"
)

//go:embed const_test.go
var src string

type HeaderName string

const (
	// DataElement                    HeaderName = "Data Element"                      // A
	// DataElementDefinition          HeaderName = "Data Element Definition"           // B
	// DataElementReferences          HeaderName = "Data Element References"           // C
	// DataElementComments            HeaderName = "Data Element Comments"             // D
	ValueDomain                    HeaderName = "Value Domain"                      // E
	ValueDomainDefinition          HeaderName = "Value Domain Definition"           // F
	ValueDomainReferences          HeaderName = "Value Domain References"           // G
	ValueDomainComments            HeaderName = "Value Domain Comments"             // H
	ValueDomainPermissibleValues   HeaderName = "Value Domain Permissible Values"   // I
	ValueDomainSupplementaryValues HeaderName = "Value Domain Supplementary Values" // J
	ValueDomainFormat              HeaderName = "Value Domain Format"               // K
	DataElementConcept             HeaderName = "Data Element Concept"              // L
	DataElementConceptDefinition   HeaderName = "Data Element Concept Definition"   // M
	DataElementConceptReferences   HeaderName = "Data Element Concept References"   // N
	DataElementConceptComments     HeaderName = "Data Element Concept Comments"     // O
	ObjectClass                    HeaderName = "Object Class"                      // P
	ObjectClassDefinition          HeaderName = "Object Class Definition"           // Q
	ObjectClassReferences          HeaderName = "Object Class References"           // R
	ObjectClassComments            HeaderName = "Object Class Comments"             // S
	Property                       HeaderName = "Property"                          // T
	PropertyDefinition             HeaderName = "Property Definition"               // U
	PropertyReferences             HeaderName = "Property References"               // V
	PropertyComments               HeaderName = "Property Comments"                 // W
	DataType                       HeaderName = "Data Type"                         // X
	DataTypeDefinition             HeaderName = "Data Type Definition"              // Y
	DataTypeReferences             HeaderName = "Data Type References"              // Z
	DataTypeComments               HeaderName = "Data Type Comments"                // AA
)

const (
	A uint = 1
	B uint = 2
	C bool = true
)

func TestValuesFromConsts(t *testing.T) {

	values, consts, err := ValuesFromConsts[HeaderName](src)
	if err != nil {
		fmt.Println(err)
		return
	}
	for i, val := range values {
		fmt.Printf("%v - %v\n", consts[i], val)
	}

	fmt.Println("----------------------------------")

	values1, consts1, err1 := ValuesFromConsts[bool](src)
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	for i, val := range values1 {
		fmt.Printf("%v - %v\n", consts1[i], val)
	}
}

func TestMapFromConsts(t *testing.T) {

	m, err := MapFromConsts[uint](src)
	if err != nil {
		fmt.Println(err)
		return
	}
	for k, v := range m {
		fmt.Println(k, ":", v)
	}

}
