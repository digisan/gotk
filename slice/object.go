package slice

// ObjIn : if arr has element e, return true. otherwise false
func ObjIn(e interface{}, arr ...interface{}) bool {
	return ObjIdxOf(e, arr...) != -1
}

// ObjNotIn : if arr does NOT have element e, return true. otherwise false
func ObjNotIn(e interface{}, arr ...interface{}) bool {
	return !ObjIn(e, arr...)
}

// ObjIdxOf : returns the index of the first instance of e in set, or -1 if e is not present in set
func ObjIdxOf(e interface{}, arr ...interface{}) int {
	for i, ele := range arr {
		if ele == e {
			return i
		}
	}
	return -1
}

// ObjLastIdxOf : returns the index of the last instance of e in set, or -1 if e is not present in set
func ObjLastIdxOf(e interface{}, arr ...interface{}) int {
	n := len(arr)
	for i := n - 1; i >= 0; i-- {
		if arr[i] == e {
			return i
		}
	}
	return -1
}