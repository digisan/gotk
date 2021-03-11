package slice

// StrIn : if arr has element e, return true. otherwise false
func StrIn(e string, arr ...string) bool {
	return StrIdxOf(e, arr...) != -1
}

// StrNotIn : if arr does NOT have element e, return true. otherwise false
func StrNotIn(e string, arr ...string) bool {
	return !StrIn(e, arr...)
}

// StrIdxOf : returns the index of the first instance of e in set, or -1 if e is not present in set
func StrIdxOf(e string, arr ...string) int {
	for i, ele := range arr {
		if ele == e {
			return i
		}
	}
	return -1
}

// StrLastIdxOf : returns the index of the last instance of e in set, or -1 if e is not present in set
func StrLastIdxOf(e string, arr ...string) int {
	n := len(arr)
	for i := n - 1; i >= 0; i-- {
		if arr[i] == e {
			return i
		}
	}
	return -1
}