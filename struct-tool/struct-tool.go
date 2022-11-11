package structtool

import (
	"fmt"
	"reflect"
	"strings"
	"unicode"

	. "github.com/digisan/go-generics/v2"
)

// get all fields
func Fields(object any) (fields []string) {
	if reflect.ValueOf(object).Kind() == reflect.Ptr {
		ptr := reflect.ValueOf(object).Elem()
		object = ptr.Interface()
	}
	typ := reflect.TypeOf(object)
	// fmt.Println("Type:", typ.Name(), "Kind:", typ.Kind())
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		fields = append(fields, field.Name)
	}
	return
}

// get only exported field value
func FieldValue(object any, field string) (any, error) {
	if len(field) > 0 && unicode.IsUpper(rune(field[0])) {
		if reflect.ValueOf(object).Kind() == reflect.Ptr {
			ptr := reflect.ValueOf(object).Elem()
			object = ptr.Interface()
		}
		r := reflect.ValueOf(object)
		f := reflect.Indirect(r).FieldByName(field)
		if f.Kind() == 0 {
			return nil, fmt.Errorf("field '%s' is NOT in struct '%v'", field, reflect.TypeOf(object))
		}
		return f.Interface(), nil
	}
	return nil, fmt.Errorf("field '%s' is NOT exported", field)
}

func Partial(object any, fields ...string) (map[string]any, error) {
	part := make(map[string]any)
	for _, field := range fields {
		v, err := FieldValue(object, field)
		if err != nil {
			return nil, err
		}
		part[field] = v
	}
	return part, nil
}

// get all tags
func Tags(object any, tag string, exclTags ...string) (tags []string) {
	if NotIn(tag, "json", "validate") {
		panic("tag must be [json, validate]")
	}
	if reflect.ValueOf(object).Kind() == reflect.Ptr {
		ptr := reflect.ValueOf(object).Elem()
		object = ptr.Interface()
	}
	typ := reflect.TypeOf(object)
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		tag := field.Tag.Get(tag)
		// fmt.Printf("%d. %v (%v), tag: '%v'\n", i+1, field.Name, field.Type.Name(), tag)
		tags = append(tags, strings.Split(tag, ",")...)
	}
	tags = Settify(tags...)
	FilterFast(&tags,
		func(i int, e string) bool {
			return len(e) > 0 && NotIn(e, exclTags...)
		},
	)
	return
}

// get all validator tags
func ValidatorTags(object any, exclTags ...string) (tags []string) {
	return Tags(object, "validate", exclTags...)
}

// get all json tags
func JsonTags(object any, exclTags ...string) (tags []string) {
	return Tags(object, "json", exclTags...)
}
