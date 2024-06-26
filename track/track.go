package track

import (
	"fmt"
	"log"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	. "github.com/digisan/go-generics"
	"github.com/digisan/gotk/strs"
)

// TrackTime : defer TrackTime(time.Now())
func TrackTime(start time.Time) {
	elapsed := time.Since(start)
	fmt.Printf("Took %s\n", elapsed)
}

func TrackCallerAll() {
	pc := make([]uintptr, 15)
	n := runtime.Callers(0, pc)
	frames := runtime.CallersFrames(pc[:n])
AGAIN:
	if frame, ok := frames.Next(); ok {
		fmt.Println(frame.File, frame.Line, frame.Function)
		goto AGAIN
	}
}

// running track(0) in caller(), return 1. caller.go, 2. caller.go line number, 3. caller package & name
func track(lvl int) (src string, line int, fn string) {
	pc := make([]uintptr, 15)
	n := runtime.Callers(lvl+2, pc)
	frames := runtime.CallersFrames(pc[:n])
	if frame, ok := frames.Next(); ok {
		return frame.File, frame.Line, frame.Function
	}
	return "", 0, ""
}

// running CallerDescription(0) in Caller(), return line1-Caller.go:line & line2-Caller
func CallerDescription(lvl int) string {
	file, line, fn := track(lvl + 1) // "+ 1" is compensating 'CallerDescription' wrapping one more level for 'track'
	return fmt.Sprintf("%s:%d\n%s", file, line, fn)
}

func CallerSrc() (dir, src string) {
	file, _, _ := track(1) // "1" is compensating 'CallerSrc' wrapping one more level for 'track'
	return filepath.Dir(file), filepath.Base(file)
}

// ancestor 0 is caller self, 1 is caller's parent ...
func CallerFn(fullPath bool, ancestor int) string {
	_, _, fn := track(1 + ancestor) // "1" is compensating 'CallerFn' wrapping one more level for 'track'
	if fullPath {
		return fn
	}
	return strs.TrimHeadToLast(fn, ".")
}

func FuncTrack(i any) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

///////////////////////////////////////////////////////////////////////////////////////////

const (
	SEP = "^"
)

var (
	smAccess   = &sync.Map{} // key - number
	smFrequent = &sync.Map{} // key - ts (@triggered)
)

func genKey(invoker string) string {
	return fmt.Sprintf("%s%s%s%s%d", invoker, SEP, CallerFn(true, 2), SEP, time.Now().Unix())
}

func parseKey(key string) (invoker, callPath string, ts int64) {
	if strings.Count(key, SEP) != 2 {
		log.Fatal("incorrect key format @SEP")
	}
	parts := strings.Split(key, SEP)
	ts, ok := AnyTryToType[int64](parts[2])
	if !ok {
		log.Fatal("incorrect key format @timestamp")
	}
	return parts[0], parts[1], ts
}

func genBatchKeys(invoker string, rng int) []string {
	keys := []string{}
	ts := time.Now().Unix()
	for i := -rng; i <= rng; i++ {
		keys = append(keys, fmt.Sprintf("%s%s%s%s%d", invoker, SEP, CallerFn(true, 2), SEP, ts+int64(i)))
	}
	return keys
}

func delRecord(prefix string) int {
	keys := []string{}
	smAccess.Range(func(key, value any) bool {
		if k := key.(string); strings.HasPrefix(k, prefix) {
			keys = append(keys, k)
		}
		return true
	})
	for _, key := range keys {
		smAccess.Delete(key)
	}
	smFrequent.Delete(prefix)
	return len(keys)
}

var (
	flag atomic.Int32 // store clean up monitor period (second)
	mtx  = sync.Mutex{}
)

func CheckAccess(invoker string, spanSec, accessLmt int) bool {

	key := genKey(invoker)
	_, _, tsNow := parseKey(key)
	if n, ok := smAccess.Load(key); ok {
		smAccess.Store(key, n.(int)+1)
	} else {
		smAccess.Store(key, 1)
	}

	// check smFrequent status...
	prefix := strs.TrimTailFromLast(key, SEP)

	if tsLast, ok := smFrequent.Load(prefix); ok {
		if tsNow-tsLast.(int64) <= int64(spanSec) {
			smFrequent.Store(prefix, tsNow) // update latest access timestamp
			return false
		} else {
			delRecord(prefix)      // after enough period, clear previous record
			smAccess.Store(key, 1) // but this time access still need to add 1
			return true
		}
	}

	// summarize total access number
	total := 0
	for _, key := range genBatchKeys(invoker, spanSec) {
		if n, ok := smAccess.Load(key); ok {
			total += n.(int)
		}
	}

	// check total access number with limit number
	if total > accessLmt {
		smFrequent.Store(prefix, tsNow)
		return false
	}

	// start single monitor for cleaning up
	mtx.Lock()
	defer mtx.Unlock()
	if flag.Load() == 0 {
		flag.Store(60) // default monitoring period
		go func() {
			fmt.Println("Starting Access Cleanup Monitor")
			for {
				period := flag.Load()
				time.Sleep(time.Duration(period) * time.Second)
				//
				tsNow := time.Now().Unix()
				smAccess.Range(func(key, value any) bool {
					if _, _, ts := parseKey(key.(string)); tsNow-ts > int64(period) {
						smAccess.Delete(key)
					}
					return true
				})
				smFrequent.Range(func(key, value any) bool {
					if tsNow-value.(int64) > int64(period) {
						smFrequent.Delete(key)
					}
					return true
				})
			}
		}()
	}

	return true
}

type CleanOption struct {
	period int32
}

func SetAccessCleanupPeriod(opt CleanOption) {
	if opt.period > 0 {
		flag.Store(opt.period)
	}
}
