package structtool

import (
	"fmt"
	"testing"
)

type TEST struct {
	a int    `json:"email0" validate:"required,email,email-db"`
	A int    `json:"email1" validate:"email1,email-db"`
	s string `json:"email2" validate:"email2,email-db"`
	S string `json:"email3" validate:"email3,email-db"`
}

func TestFields(t *testing.T) {

	test := &TEST{
		a: 1,
		A: 12,
		s: "s",
		S: "SS",
	}

	for _, f := range Fields(test) {
		fmt.Println(f, FieldValue(test, f))
	}

	fmt.Println()

	for _, f := range Fields(*test) {
		fmt.Println(f, FieldValue(*test, f))
	}

	fmt.Println()

	fmt.Println(ValidatorTags(test, "validate", "email"))
	fmt.Println(Tags(*test, "validate", "required", "email-db"))

	fmt.Println()

	fmt.Println(JsonTags(test))
	fmt.Println(Tags(*test, "json", "email3"))
}
