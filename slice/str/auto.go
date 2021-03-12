package str

// In : if arr has element e, return true. otherwise false
func In(e string, arr ...string) bool {
	return IdxOf(e, arr...) != -1
}

// NotIn : if arr does NOT have element e, return true. otherwise false
func NotIn(e string, arr ...string) bool {
	return !In(e, arr...)
}

// IdxOf : returns the index of the first instance of e in slice, or -1 if e is not present in slice
func IdxOf(e string, arr ...string) int {
	for i, ele := range arr {
		if ele == e {
			return i
		}
	}
	return -1
}

// LastIdxOf : returns the index of the last instance of e in slice, or -1 if e is not present in slice
func LastIdxOf(e string, arr ...string) int {
	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] == e {
			return i
		}
	}
	return -1
}

// MkSet : remove repeated elements in arr
func MkSet(arr ...string) (set []string) {
	if arr == nil {
		return nil
	}
	m := make(map[string]struct{})
	for _, ele := range arr {
		if _, ok := m[ele]; !ok {
			set = append(set, ele)
			m[ele] = struct{}{}
		}
	}
	if len(set) == 0 {
		return []string{}
	}
	return
}

// Superset :
func Superset(setA, setB []string) bool {
NEXT_B:
	for _, b := range setB {
		for _, a := range setA {
			if b == a {
				continue NEXT_B
			}
		}
		return false
	}
	return len(setA) > len(setB)
}

// Subset :
func Subset(setA, setB []string) bool {
	return Superset(setB, setA)
}

// Equal :
func Equal(setA, setB []string) bool {
	if (setA == nil && setB != nil) || (setA != nil && setB == nil) {
		return false
	}
	return !Superset(setA, setB) && !Subset(setA, setB)
}

// union :
func union(setA, setB []string) (set []string) {
	if setA == nil && setB == nil {
		return nil
	}
	if setA == nil && setB != nil {
		return setB
	}
	if setA != nil && setB == nil {
		return setA
	}

	m := make(map[string]struct{})
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
		return []string{}
	}
	return
}

// Union :
func Union(sets ...[]string) (set []string) {
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
func intersect(setA, setB []string) (set []string) {
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
		return []string{}
	}
	return
}

// Intersect :
func Intersect(sets ...[]string) (set []string) {
	if len(sets) == 0 {
		return nil
	}
	set = sets[0]
	for _, s := range sets[1:] {
		set = intersect(set, s)
	}
	return set
}

// FilterModify : Filter & Modify []string slice, return []string slice
func FilterModify(arr []string, filter func(i int, e string) bool, modifier func(i int, e string) string) (r []string) {
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