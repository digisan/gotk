package ti

// In : if arr has element e, return true. otherwise false
func In(e int, arr ...int) bool {
	return IdxOf(e, arr...) != -1
}

// NotIn : if arr does NOT have element e, return true. otherwise false
func NotIn(e int, arr ...int) bool {
	return !In(e, arr...)
}

// IdxOf : returns the index of the first instance of e in slice, or -1 if e is not present in slice
func IdxOf(e int, arr ...int) int {
	for i, ele := range arr {
		if ele == e {
			return i
		}
	}
	return -1
}

// LastIdxOf : returns the index of the last instance of e in slice, or -1 if e is not present in slice
func LastIdxOf(e int, arr ...int) int {
	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] == e {
			return i
		}
	}
	return -1
}

// MkSet : remove repeated elements in arr
func MkSet(arr ...int) (set []int) {
	if arr == nil {
		return nil
	}
	m := make(map[int]struct{})
	for _, ele := range arr {
		if _, ok := m[ele]; !ok {
			set = append(set, ele)
			m[ele] = struct{}{}
		}
	}
	if len(set) == 0 {
		return []int{}
	}
	return
}

// Superset :
func Superset(setA, setB []int) bool {
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
func Subset(setA, setB []int) bool {
	return Superset(setB, setA)
}

// Equal :
func Equal(setA, setB []int) bool {
	if (setA == nil && setB != nil) || (setA != nil && setB == nil) {
		return false
	}
	if len(setA) != len(setB) {
		return false
	}

	tmpA := make([]int, len(setA))
	tmpB := make([]int, len(setB))
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
func SuperEq(setA, setB []int) bool {
	return Superset(setA, setB) || Equal(setA, setB)
}

// SubEq :
func SubEq(setA, setB []int) bool {
	return Subset(setA, setB) || Equal(setA, setB)
}

// union :
func union(setA, setB []int) (set []int) {
	if setA == nil && setB == nil {
		return nil
	}
	if setA == nil && setB != nil {
		return setB
	}
	if setA != nil && setB == nil {
		return setA
	}

	m := make(map[int]struct{})
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
		return []int{}
	}
	return
}

// Union :
func Union(sets ...[]int) (set []int) {
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
func intersect(setA, setB []int) (set []int) {
	if setA == nil || setB == nil {
		return nil
	}
AGAIN:
	for i, a := range setA {
		for j, b := range setB {
			if a == b {
				set = append(set, a)
				setA = append(setA[:i], setA[i+1:]...)
				setB = append(setB[:j], setB[i+j:]...)
				goto AGAIN
			}
		}
	}
	if set == nil {
		return []int{}
	}
	return
}

// Intersect :
func Intersect(sets ...[]int) (set []int) {
	if len(sets) == 0 {
		return nil
	}
	set = sets[0]
	for _, s := range sets[1:] {
		set = intersect(set, s)
	}
	return set
}

// Reorder : any index must less than len(arr)
func Reorder(arr []int, indices []int) (orders []int) {
	if arr == nil || indices == nil {
		return nil
	}
	if len(arr) == 0 || len(indices) == 0 {
		return []int{}
	}
	for _, idx := range indices {
		orders = append(orders, arr[idx])
	}
	return orders
}

// FilterModify : Filter & Modify []int slice, return []int slice
func FilterModify(arr []int, filter func(i int, e int) bool, modifier func(i int, e int) int) (r []int) {
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