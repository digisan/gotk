package src

import (
	"fmt"
	"runtime"
	"testing"
)

type HeaderName string

const (
	DataElement                    HeaderName = "Data Element"                      // A
	DataElementDefinition          HeaderName = "Data Element Definition"           // B
	DataElementReferences          HeaderName = "Data Element References"           // C
	DataElementComments            HeaderName = "Data Element Comments"             // D
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
	C bool = false
)

func TestValuesFromConsts(t *testing.T) {
	_, fSrcPath, _, _ := runtime.Caller(0)

	values, err := ValuesFromConsts[HeaderName](fSrcPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, val := range values {
		fmt.Printf("%v\n", val)
	}

	fmt.Println("----------------------------------")

	values1, err1 := ValuesFromConsts[uint](fSrcPath)
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	for _, val := range values1 {
		fmt.Printf("%v\n", val)
	}
}
