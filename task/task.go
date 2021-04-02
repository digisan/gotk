package task

import (
	"fmt"
	"reflect"
)

func funcWrap(fn interface{}) func(...interface{}) interface{} {
	if reflect.TypeOf(fn).Kind() != reflect.Func {
		panic("[fn] argument is not a function")
	}
	return func(args ...interface{}) (rt interface{}) {
		defer func() {
			if r := recover(); r != nil {
				rt = fmt.Errorf("%v", r)
			}
		}()

		vArgs := make([]reflect.Value, len(args))
		for n, v := range args {
			vArgs[n] = reflect.ValueOf(v)
		}
		var rtVal []interface{}
		for _, v := range reflect.ValueOf(fn).Call(vArgs) {
			rtVal = append(rtVal, v.Interface())
		}
		return rtVal
	}
}

// Task : only accept one param.
// Define " cP4f, cR4f := make(chan interface{}), make(chan interface{}) "
// Then   " Task(true, f, cP4f, cR4f) "
// Then   " go func() { -> for { -> select { -> case rf := <-cR4f: ... "
// Please see "task_test.go"
func Task(async bool, fn interface{}, cParam <-chan interface{}, cRet chan<- interface{}) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
		}
	}()
	svrFn := funcWrap(fn)
	go func() {
		if async {
			for param := range cParam {
				go func(param interface{}) { cRet <- svrFn(param) }(param)
			}
		} else {
			for param := range cParam {
				cRet <- svrFn(param)
			}
		}
	}()
	return nil
}
