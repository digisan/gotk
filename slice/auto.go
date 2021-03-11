package slice

// StrIn : if arr has element e, return true. otherwise false
func StrIn(e string, arr ...string) bool {
	return StrIdxOf(e, arr...) != -1
}

// StrNotIn : if arr does NOT have element e, return true. otherwise false
func StrNotIn(e string, arr ...string) bool {
	return !StrIn(e, arr...)
}

// StrIdxOf : returns the index of the first instance of e in slice, or -1 if e is not present in slice
func StrIdxOf(e string, arr ...string) int {
	for i, ele := range arr {
		if ele == e {
			return i
		}
	}
	return -1
}

// StrLastIdxOf : returns the index of the last instance of e in slice, or -1 if e is not present in slice
func StrLastIdxOf(e string, arr ...string) int {
	n := len(arr)
	for i := n - 1; i >= 0; i-- {
		if arr[i] == e {
			return i
		}
	}
	return -1
}

// IntIn : if arr has element e, return true. otherwise false
func IntIn(e int, arr ...int) bool {
	return IntIdxOf(e, arr...) != -1
}

// IntNotIn : if arr does NOT have element e, return true. otherwise false
func IntNotIn(e int, arr ...int) bool {
	return !IntIn(e, arr...)
}

// IntIdxOf : returns the index of the first instance of e in slice, or -1 if e is not present in slice
func IntIdxOf(e int, arr ...int) int {
	for i, ele := range arr {
		if ele == e {
			return i
		}
	}
	return -1
}

// IntLastIdxOf : returns the index of the last instance of e in slice, or -1 if e is not present in slice
func IntLastIdxOf(e int, arr ...int) int {
	n := len(arr)
	for i := n - 1; i >= 0; i-- {
		if arr[i] == e {
			return i
		}
	}
	return -1
}

// ObjIn : if arr has element e, return true. otherwise false
func ObjIn(e interface{}, arr ...interface{}) bool {
	return ObjIdxOf(e, arr...) != -1
}

// ObjNotIn : if arr does NOT have element e, return true. otherwise false
func ObjNotIn(e interface{}, arr ...interface{}) bool {
	return !ObjIn(e, arr...)
}

// ObjIdxOf : returns the index of the first instance of e in slice, or -1 if e is not present in slice
func ObjIdxOf(e interface{}, arr ...interface{}) int {
	for i, ele := range arr {
		if ele == e {
			return i
		}
	}
	return -1
}

// ObjLastIdxOf : returns the index of the last instance of e in slice, or -1 if e is not present in slice
func ObjLastIdxOf(e interface{}, arr ...interface{}) int {
	n := len(arr)
	for i := n - 1; i >= 0; i-- {
		if arr[i] == e {
			return i
		}
	}
	return -1
}

// IntsFMStr : Filter & Modify []int slice, return []string slice
func IntsFMStr(arr []int, filter func(i int) bool, modifier func(i int) string) (r []string) {

	if modifier == nil {
		panic("modifier cannot be nil")
	}
	switch {
	case filter != nil:
		for i := range arr {
			if filter(i) {
				r = append(r, modifier(i))
			}
		}
	default:
		for i := range arr {
			r = append(r, modifier(i))
		}
	}
	return
}

// StrsFM : Filter & Modify []string slice, return []string slice
func StrsFM(arr []string, filter func(i int) bool, modifier func(i int) string) (r []string) {
	switch {
	case filter != nil && modifier != nil:
		for i := range arr {
			if filter(i) {
				r = append(r, modifier(i))
			}
		}
	case filter != nil && modifier == nil:
		for i, e := range arr {
			if filter(i) {
				r = append(r, e)
			}
		}
	case filter == nil && modifier != nil:
		for i := range arr {
			r = append(r, modifier(i))
		}
	default:
		return arr
	}
	return

}

// StrsFMInt : Filter & Modify []string slice, return []int slice
func StrsFMInt(arr []string, filter func(i int) bool, modifier func(i int) int) (r []int) {

	if modifier == nil {
		panic("modifier cannot be nil")
	}
	switch {
	case filter != nil:
		for i := range arr {
			if filter(i) {
				r = append(r, modifier(i))
			}
		}
	default:
		for i := range arr {
			r = append(r, modifier(i))
		}
	}
	return
}
