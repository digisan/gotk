package to

import "sort"

// In : if arr has element e, return true. otherwise false
func In(e interface{}, arr ...interface{}) bool {
	return IdxOf(e, arr...) != -1
}

// NotIn : if arr does NOT have element e, return true. otherwise false
func NotIn(e interface{}, arr ...interface{}) bool {
	return !In(e, arr...)
}

// IdxOf : returns the index of the first instance of e in slice, or -1 if e is not present in slice
func IdxOf(e interface{}, arr ...interface{}) int {
	for i, ele := range arr {
		if ele == e {
			return i
		}
	}
	return -1
}

// LastIdxOf : returns the index of the last instance of e in slice, or -1 if e is not present in slice
func LastIdxOf(e interface{}, arr ...interface{}) int {
	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] == e {
			return i
		}
	}
	return -1
}

// MkSet : remove repeated elements in arr
func MkSet(arr ...interface{}) (set []interface{}) {
	if arr == nil {
		return nil
	}
	m := make(map[interface{}]struct{})
	for _, ele := range arr {
		if _, ok := m[ele]; !ok {
			set = append(set, ele)
			m[ele] = struct{}{}
		}
	}
	if len(set) == 0 {
		return []interface{}{}
	}
	return
}

// Superset :
func Superset(setA, setB []interface{}) bool {
NEXT_B:
	for _, b := range setB {
		for _, a := range setA {
			if a == b {
				continue NEXT_B
			}
		}
		return false
	}
	return len(setA) > len(setB)
}

// Subset :
func Subset(setA, setB []interface{}) bool {
	return Superset(setB, setA)
}

// Equal :
func Equal(setA, setB []interface{}) bool {
	if (setA == nil && setB != nil) || (setA != nil && setB == nil) {
		return false
	}
	if len(setA) != len(setB) {
		return false
	}

	tmpA := make([]interface{}, len(setA))
	tmpB := make([]interface{}, len(setB))
	copy(tmpA, setA)
	copy(tmpB, setB)

AGAIN:
	for i, a := range tmpA {
		for j, b := range tmpB {
			if a == b {
				tmpA = append(tmpA[:i], tmpA[i+1:]...)
				tmpB = append(tmpB[:j], tmpB[j+1:]...)
				goto AGAIN
			}
		}
	}
	return len(tmpA) == 0 && len(tmpB) == 0
}

// SuperEq :
func SuperEq(setA, setB []interface{}) bool {
	return Superset(setA, setB) || Equal(setA, setB)
}

// SubEq :
func SubEq(setA, setB []interface{}) bool {
	return Subset(setA, setB) || Equal(setA, setB)
}

// union :
func union(setA, setB []interface{}) (set []interface{}) {
	if setA == nil && setB == nil {
		return nil
	}
	if setA == nil && setB != nil {
		return setB
	}
	if setA != nil && setB == nil {
		return setA
	}

	m := make(map[interface{}]struct{})
	for _, a := range setA {
		if _, ok := m[a]; !ok {
			set = append(set, a)
			m[a] = struct{}{}
		}
	}
	for _, b := range setB {
		if _, ok := m[b]; !ok {
			set = append(set, b)
			m[b] = struct{}{}
		}
	}
	if set == nil {
		return []interface{}{}
	}
	return
}

// Union :
func Union(sets ...[]interface{}) (set []interface{}) {
	if len(sets) == 0 {
		return nil
	}
	set = sets[0]
	for _, s := range sets[1:] {
		set = union(set, s)
	}
	return set
}

// intersect :
func intersect(setA, setB []interface{}) (set []interface{}) {
	if setA == nil || setB == nil {
		return nil
	}

	copyA, copyB := make([]interface{}, len(setA)), make([]interface{}, len(setB))
	copy(copyA, setA)
	copy(copyB, setB)

AGAIN:
	for i, a := range copyA {
		for j, b := range copyB {
			if a == b {
				set = append(set, a)
				copyA = append(copyA[:i], copyA[i+1:]...)
				copyB = append(copyB[:j], copyB[i+j:]...)
				goto AGAIN
			}
		}
	}
	if set == nil {
		return []interface{}{}
	}
	return
}

// Intersect :
func Intersect(sets ...[]interface{}) (set []interface{}) {
	if len(sets) == 0 {
		return nil
	}
	set = sets[0]
	for _, s := range sets[1:] {
		set = intersect(set, s)
	}
	return set
}

func minus(setA, setB []interface{}) (set []interface{}) {
	if setA == nil {
		return nil
	}
	set = make([]interface{}, 0)

NEXT_A:
	for _, a := range setA {
		for _, b := range setB {
			if a == b {
				continue NEXT_A
			}
		}
		set = append(set, a)
	}
	return
}

func Minus(setA []interface{}, setOthers ...[]interface{}) (set []interface{}) {
	return minus(setA, Union(setOthers...))
}

// Reorder : any index must less than len(arr)
func Reorder(arr []interface{}, indices []int) (orders []interface{}) {
	if arr == nil || indices == nil {
		return nil
	}
	if len(arr) == 0 || len(indices) == 0 {
		return []interface{}{}
	}
	for _, idx := range indices {
		orders = append(orders, arr[idx])
	}
	return orders
}

// Reverse : [1,2,3] => [3,2,1]
func Reverse(arr []interface{}) []interface{} {
	indices := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		indices[i] = len(arr) - 1 - i
	}
	return Reorder(arr, indices)
}

// FilterModify : Filter & Modify []interface{} slice, return []interface{} slice
func FilterModify(arr []interface{}, filter func(i int, e interface{}) bool, modifier func(i int, e interface{}) interface{}) (r []interface{}) {
	switch {
	case filter != nil && modifier != nil:
		for i, e := range arr {
			if filter(i, e) {
				r = append(r, modifier(i, e))
			}
		}
	case filter != nil && modifier == nil:
		for i, e := range arr {
			if filter(i, e) {
				r = append(r, e)
			}
		}
	case filter == nil && modifier != nil:
		for i, e := range arr {
			r = append(r, modifier(i, e))
		}
	default:
		return arr
	}
	return
}

var (
	FM = FilterModify
)

// Map2KVs : map to key slice & value slice
func Map2KVs(m map[interface{}]interface{}, less4key func(i interface{}, j interface{}) bool, less4value func(i interface{}, j interface{}) bool) (keys []interface{}, values []interface{}) {

	if m == nil {
		return nil, nil
	}
	if len(m) == 0 {
		return []interface{}{}, []interface{}{}
	}

	type kv struct {
		key   interface{}
		value interface{}
	}

	kvSlc := []kv{}
	for k, v := range m {
		kvSlc = append(kvSlc, kv{key: k, value: v})
	}

	switch {
	case less4key != nil && less4value == nil:
		sort.SliceStable(kvSlc, func(i, j int) bool { return less4key(kvSlc[i].key, kvSlc[j].key) })

	case less4key == nil && less4value != nil:
		sort.SliceStable(kvSlc, func(i, j int) bool { return less4value(kvSlc[i].value, kvSlc[j].value) })

	case less4key != nil && less4value != nil:
		sort.SliceStable(kvSlc, func(i, j int) bool {
			if kvSlc[i].value == kvSlc[j].value {
				return less4key(kvSlc[i].key, kvSlc[j].key)
			}
			return less4value(kvSlc[i].value, kvSlc[j].value)
		})

	default:
		// do not sort
	}

	for _, kvEle := range kvSlc {
		keys = append(keys, kvEle.key)
		values = append(values, kvEle.value)
	}
	return
}
