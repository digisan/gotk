package stack

// Stack :
type Stack []interface{}

// Push :
func (stk *Stack) Push(items ...interface{}) (n int) {
	*stk = append(*stk, items...)
	return len(items)
}

func (stk *Stack) len() int {
	return len(*stk)
}

// Len :
func (stk *Stack) Len() int {
	return len(*stk)
}

// Pop :
func (stk *Stack) Pop() (interface{}, bool) {
	if stk.len() > 0 {
		last := (*stk)[stk.len()-1]
		*stk = (*stk)[:stk.len()-1]
		return last, true
	}
	return nil, false
}

// Peek :
func (stk *Stack) Peek() (interface{}, bool) {
	if stk.len() > 0 {
		return (*stk)[stk.len()-1], true
	}
	return nil, false
}

// Clear :
func (stk *Stack) Clear() (Stack, int) {
	return Stack{}, stk.len()
}

// Copy :
func (stk *Stack) Copy() Stack {
	tmp := make([]interface{}, stk.len())
	copy(tmp, *stk)
	return Stack(tmp)
}

// Sprint :
func (stk *Stack) Sprint(sep string) string {
	sb := sBuilder{}
	for _, ele := range *stk {
		sb.WriteString(fSf("%v", ele))
		sb.WriteString(sep)
	}
	return sTrimRight(sb.String(), sep)
}
