package tso

import (
	"fmt"
	"testing"
)

func TestMap2KVs(t *testing.T) {
	m := make(map[string]interface{})
	m["3"] = 5
	m["1"] = 8
	m["0"] = 2
	m["7"] = 3
	m["233"] = 6
	ks, vs := Map2KVs(m, func(i, j string) bool { return i < j }, nil)
	fmt.Println(ks)
	fmt.Println(vs)

	ks, vs = Map2KVs(m, nil, func(i, j interface{}) bool { return i.(int) < j.(int) })
	fmt.Println(ks)
	fmt.Println(vs)
}
