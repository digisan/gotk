package track

import (
	"fmt"
	"testing"
	"time"
)

func Test_trackCaller(t *testing.T) {

	// fmt.Println("track0: ----------------")
	// fmt.Println(track(0))

	// fmt.Println("track1: ----------------")
	// fmt.Println(track(1))

	fmt.Println("TrackCaller0: ----------------")
	fmt.Println(CallerDescription(0))

	fmt.Println("TrackCaller1: ----------------")
	fmt.Println(CallerDescription(1))

	// fmt.Println("CallerSrc: ----------------")
	// fmt.Println(CallerSrc())

	// fmt.Println("Caller: ----------------")
	// fmt.Println(Caller(true))
	// fmt.Println(Caller(false))
}

func TestRecord(t *testing.T) {

	span := 1
	limit := 5

	for i := 0; i < 100; i++ {
		canAccess := CheckAccess("user", span, limit)
		fmt.Println(canAccess)
		time.Sleep(time.Duration(100 * time.Millisecond))
	}
}
