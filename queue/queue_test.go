package queue

import (
	"reflect"
	"testing"
)

func TestQueue_Enqueue(t *testing.T) {
	type args struct {
		items []interface{}
	}
	tests := []struct {
		name  string
		q     *Queue
		args  args
		wantN int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotN := tt.q.Enqueue(tt.args.items...); gotN != tt.wantN {
				t.Errorf("Queue.Enqueue() = %v, want %v", gotN, tt.wantN)
			}
		})
	}
}

func TestQueue_len(t *testing.T) {
	tests := []struct {
		name string
		q    *Queue
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.q.len(); got != tt.want {
				t.Errorf("Queue.len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueue_Len(t *testing.T) {
	tests := []struct {
		name string
		q    *Queue
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.q.Len(); got != tt.want {
				t.Errorf("Queue.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueue_Dequeue(t *testing.T) {
	tests := []struct {
		name  string
		q     *Queue
		want  interface{}
		want1 bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.q.Dequeue()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Queue.Dequeue() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Queue.Dequeue() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestQueue_Peek(t *testing.T) {
	tests := []struct {
		name  string
		q     *Queue
		want  interface{}
		want1 bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.q.Peek()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Queue.Peek() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Queue.Peek() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestQueue_Clear(t *testing.T) {
	tests := []struct {
		name  string
		q     *Queue
		want  Queue
		want1 int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.q.Clear()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Queue.Clear() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Queue.Clear() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestQueue_Copy(t *testing.T) {
	tests := []struct {
		name string
		q    *Queue
		want Queue
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.q.Copy(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Queue.Copy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueue_Sprint(t *testing.T) {
	type args struct {
		sep string
	}
	tests := []struct {
		name string
		q    *Queue
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.q.Sprint(tt.args.sep); got != tt.want {
				t.Errorf("Queue.Sprint() = %v, want %v", got, tt.want)
			}
		})
	}
}
