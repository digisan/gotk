package structtool

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"unicode"

	. "github.com/digisan/go-generics/v2"
)

// get all fields
func Fields(obj any) (fields []string) {
	if reflect.ValueOf(obj).Kind() == reflect.Ptr {
		ptr := reflect.ValueOf(obj).Elem()
		obj = ptr.Interface()
	}
	typ := reflect.TypeOf(obj)
	// fmt.Println("Type:", typ.Name(), "Kind:", typ.Kind())
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		fields = append(fields, field.Name)
	}
	return
}

// get only exported field value
func FieldValue(obj any, field string) any {
	if len(field) > 0 && unicode.IsUpper(rune(field[0])) {
		if reflect.ValueOf(obj).Kind() == reflect.Ptr {
			ptr := reflect.ValueOf(obj).Elem()
			obj = ptr.Interface()
		}
		r := reflect.ValueOf(obj)
		f := reflect.Indirect(r).FieldByName(field)
		return f.Interface()
	}
	return errors.New(fmt.Sprintf("'[%s] is unexported'", field))
}

// get all tags
func Tags(obj any, tag string, exclTags ...string) (tags []string) {
	if NotIn(tag, "json", "validate") {
		panic("tag must be [json, validate]")
	}
	if reflect.ValueOf(obj).Kind() == reflect.Ptr {
		ptr := reflect.ValueOf(obj).Elem()
		obj = ptr.Interface()
	}
	typ := reflect.TypeOf(obj)
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		tag := field.Tag.Get(tag)
		// fmt.Printf("%d. %v (%v), tag: '%v'\n", i+1, field.Name, field.Type.Name(), tag)
		tags = append(tags, strings.Split(tag, ",")...)
	}
	tags = Settify(tags...)
	Filter(&tags,
		func(i int, e string) bool {
			return len(e) > 0 && NotIn(e, exclTags...)
		},
	)
	return
}

// get all validator tags
func ValidatorTags(obj any, exclTags ...string) (tags []string) {
	return Tags(obj, "validate", exclTags...)
}

// get all json tags
func JsonTags(obj any, exclTags ...string) (tags []string) {
	return Tags(obj, "json", exclTags...)
}
