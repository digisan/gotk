package to

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

AGAIN:
	for i, a := range setA {
		for j, b := range setB {
			if a == b {
				setA = append(setA[:i], setA[i+1:]...)
				setB = append(setB[:j], setB[j+1:]...)
				goto AGAIN
			}
		}
	}
	return len(setA) == 0 && len(setB) == 0
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