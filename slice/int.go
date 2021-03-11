package slice

// IntIn : if arr has element e, return true. otherwise false
func IntIn(e int, arr ...int) bool {
	return IntIdxOf(e, arr...) != -1
}

// IntNotIn : if arr does NOT have element e, return true. otherwise false
func IntNotIn(e int, arr ...int) bool {
	return !IntIn(e, arr...)
}

// IntIdxOf : returns the index of the first instance of e in set, or -1 if e is not present in set
func IntIdxOf(e int, arr ...int) int {
	for i, ele := range arr {
		if ele == e {
			return i
		}
	}
	return -1
}

// IntLastIdxOf : returns the index of the last instance of e in set, or -1 if e is not present in set
func IntLastIdxOf(e int, arr ...int) int {
	n := len(arr)
	for i := n - 1; i >= 0; i-- {
		if arr[i] == e {
			return i
		}
	}
	return -1
}