package queue

// Queue :
type Queue []interface{}

// Enqueue :
func (q *Queue) Enqueue(items ...interface{}) (n int) {
	*q = append(*q, items...)
	return len(items)
}

func (q *Queue) len() int {
	return len(*q)
}

// Len :
func (q *Queue) Len() int {
	return len(*q)
}

// Dequeue :
func (q *Queue) Dequeue() (interface{}, bool) {
	if q.len() > 0 {
		first := (*q)[0]
		*q = (*q)[1:]
		return first, true
	}
	return nil, false
}

// Peek :
func (q *Queue) Peek() (interface{}, bool) {
	if q.len() > 0 {
		return (*q)[0], true
	}
	return nil, false
}

// Clear :
func (q *Queue) Clear() (Queue, int) {
	return Queue{}, q.len()
}

// Copy :
func (q *Queue) Copy() Queue {
	tmp := make([]interface{}, q.len())
	copy(tmp, *q)
	return Queue(tmp)
}

// Sprint :
func (q *Queue) Sprint(sep string) string {
	sb := sBuilder{}
	for _, ele := range *q {
		sb.WriteString(fSf("%v", ele))
		sb.WriteString(sep)
	}
	return sTrimRight(sb.String(), sep)
}
