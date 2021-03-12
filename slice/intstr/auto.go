package intstr

// FilterModify : Filter & Modify []int slice, return []string slice
func FilterModify(arr []int, filter func(i int, e int) bool, modifier func(i int, e int) string) (r []string) {
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