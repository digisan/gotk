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

	span := 2
	limit := 3

	for i := 0; i < 100; i++ {
		canAccess := CheckAccess("user", span, limit)
		fmt.Println(canAccess, i)
		time.Sleep(time.Duration(100 * time.Millisecond))

		if i == 40 {
			fmt.Println("Sleeping... Second")
			time.Sleep(time.Duration(3 * time.Second))

			SetAccessCleanupPeriod(CleanOption{period: 4})
		}
		if i == 70 {
			fmt.Println("Sleeping... Second")
			time.Sleep(time.Duration(3 * time.Second))

			SetAccessCleanupPeriod(CleanOption{period: 3})
		}
	}

	for i := 0; i < 10; i++ {

		n := 0
		smAccess.Range(func(key, value any) bool {
			fmt.Println(key, value)
			n++
			return true
		})
		fmt.Println("smAccess:", n)

		n = 0
		smFrequent.Range(func(key, value any) bool {
			fmt.Println(key, value)
			n++
			return true
		})
		fmt.Println("smFrequent:", n)

		time.Sleep(time.Duration(2 * time.Second))
	}
}
