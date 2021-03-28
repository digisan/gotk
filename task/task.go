package task

import "reflect"

func funcWrap(fn interface{}) func(...interface{}) interface{} {
	if reflect.TypeOf(fn).Kind() != reflect.Func {
		panic("fn argument is not a function")
	}
	return func(args ...interface{}) interface{} {
		vArgs := make([]reflect.Value, len(args))
		for n, v := range args {
			vArgs[n] = reflect.ValueOf(v)
		}
		var rt []interface{}
		for _, v := range reflect.ValueOf(fn).Call(vArgs) {
			rt = append(rt, v.Interface())
		}
		return rt
	}
}

// Task : only accept one param.
// Define " cP4f, cR4f := make(chan interface{}), make(chan interface{}) "
// Then   " Task(f, cP4f, cR4f) "
// Then   " go func() { -> for { -> select { -> case rf := <-cR4f: ... "
func Task(fn interface{}, cParam, cRet chan interface{}) {
	svrFn := funcWrap(fn)
	go func() {
		for param := range cParam {
			if ss, ok := param.([]string); ok && len(ss) == 1 {
				if ss[0] == "task exit" {
					return
				}
			}
			cRet <- svrFn(param)
		}
	}()
}
