package tsi

// FilterModify : Filter & Modify []string slice, return []int slice
func FilterModify(arr []string, filter func(i int, e string) bool, modifier func(i int, e string) int) (r []int) {
	if modifier == nil {
		panic("modifier cannot be nil")
	}

	switch {
	case filter != nil:
		for i, e := range arr {
			if filter(i, e) {
				r = append(r, modifier(i, e))
			}
		}
	default:
		for i, e := range arr {
			r = append(r, modifier(i, e))
		}
	}
	return
}

var (
	FM = FilterModify
)