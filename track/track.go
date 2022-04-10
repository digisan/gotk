package track

import (
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"

	. "github.com/digisan/gotk/strs"
)

// running track(0) in caller(), return 1. caller.go, 2. caller.go line number, 3. caller package & name
func track(lvl int) (string, int, string) {
	pc := make([]uintptr, 15)
	n := runtime.Callers(lvl+2, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	return frame.File, frame.Line, frame.Function
}

// running TrackCaller(0) in Caller(), return line1-Caller.go:line & line2-Caller
func TrackCaller(lvl int) string {
	file, line, fn := track(lvl + 1)
	return fmt.Sprintf("\n%s:%d\n%s\n", file, line, fn)
}

func CallerSrc() (dir, src string) {
	file, _, _ := track(1)
	return filepath.Dir(file), filepath.Base(file)
}

func Caller(fullpath bool) string {
	_, _, fn := track(1)
	if fullpath {
		return fn
	}
	return TrimHeadToLast(fn, ".")
}

func ParentCaller(fullpath bool) string {
	_, _, fn := track(2)
	if fullpath {
		return fn
	}
	return TrimHeadToLast(fn, ".")
}

func FuncTrack(i any) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}
