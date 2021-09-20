package tu8i

import (
	"fmt"
	"testing"
)

func TestMap2KVs(t *testing.T) {
	m := make(map[byte]int)
	m[3] = 5
	m[1] = 8
	m[0] = 2
	m[7] = 3
	m[233] = 6
	ks, vs := Map2KVs(m, func(i, j byte) bool { return i < j }, nil)
	fmt.Println(ks)
	fmt.Println(vs)

	ks, vs = Map2KVs(m, nil, func(i, j int) bool { return i < j })
	fmt.Println(ks)
	fmt.Println(vs)
}
