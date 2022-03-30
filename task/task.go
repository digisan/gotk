package task

import (
	"fmt"
	"reflect"
)

func funcWrap(fn any) func(...any) any {
	if reflect.TypeOf(fn).Kind() != reflect.Func {
		panic("[fn] argument is not a function")
	}
	return func(args ...any) (rt any) {
		defer func() {
			if r := recover(); r != nil {
				rt = fmt.Errorf("%v", r)
			}
		}()

		vArgs := make([]reflect.Value, len(args))
		for n, v := range args {
			vArgs[n] = reflect.ValueOf(v)
		}
		var rtVal []any
		for _, v := range reflect.ValueOf(fn).Call(vArgs) {
			rtVal = append(rtVal, v.Interface())
		}
		return rtVal
	}
}

// Task : only accept one param.
// Define " cP4f, cR4f := make(chan any), make(chan any) "
// Then   " Task(true, f, cP4f, cR4f) "
// Then   " go func() { -> for { -> select { -> case rf := <-cR4f: ... "
// Please see "task_test.go"
func Task(async bool, fn any, cParam <-chan any, cRet chan<- any) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
		}
	}()
	svrFn := funcWrap(fn)
	go func() {
		if async {
			for param := range cParam {
				go func(param any) { cRet <- svrFn(param) }(param)
			}
		} else {
			for param := range cParam {
				cRet <- svrFn(param)
			}
		}
	}()
	return nil
}
