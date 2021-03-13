package stack

import (
	"reflect"
	"testing"
)

func TestStack_Push(t *testing.T) {
	type args struct {
		items []interface{}
	}
	tests := []struct {
		name  string
		stk   *Stack
		args  args
		wantN int
	}{
		// TODO: Add test cases.
		{
			name:  "OK",
			stk:   &Stack{1, 2, 3},
			args:  args{items: []interface{}{4, 5}},
			wantN: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotN := tt.stk.Push(tt.args.items...); gotN != tt.wantN {
				t.Errorf("Stack.Push() = %v, want %v", gotN, tt.wantN)
			}
			fPln(*tt.stk)
		})
	}
}

func TestStack_len(t *testing.T) {
	tests := []struct {
		name string
		stk  *Stack
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.stk.len(); got != tt.want {
				t.Errorf("Stack.len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStack_Len(t *testing.T) {
	tests := []struct {
		name string
		stk  *Stack
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.stk.Len(); got != tt.want {
				t.Errorf("Stack.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStack_Pop(t *testing.T) {
	tests := []struct {
		name  string
		stk   *Stack
		want  interface{}
		want1 bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.stk.Pop()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Stack.Pop() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Stack.Pop() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestStack_Peek(t *testing.T) {
	tests := []struct {
		name  string
		stk   *Stack
		want  interface{}
		want1 bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.stk.Peek()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Stack.Peek() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Stack.Peek() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestStack_Clear(t *testing.T) {
	tests := []struct {
		name  string
		stk   *Stack
		want  Stack
		want1 int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.stk.Clear()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Stack.Clear() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Stack.Clear() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestStack_Copy(t *testing.T) {
	tests := []struct {
		name string
		stk  *Stack
		want Stack
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.stk.Copy(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Stack.Copy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStack_Sprint(t *testing.T) {
	type args struct {
		sep string
	}
	tests := []struct {
		name string
		stk  *Stack
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.stk.Sprint(tt.args.sep); got != tt.want {
				t.Errorf("Stack.Sprint() = %v, want %v", got, tt.want)
			}
		})
	}
}
