package track

import (
	"fmt"
	"testing"
)

func Test_trackCaller(t *testing.T) {
	fmt.Println(track(0))
	fmt.Println("----------------")
}

func TestTrackCaller(t *testing.T) {
	fmt.Println(TrackCaller(0))
	fmt.Println("----------------")
}

func TestCallerSrc(t *testing.T) {
	fmt.Println(CallerSrc())
	fmt.Println("----------------")
}

func TestCaller(t *testing.T) {
	fmt.Println(Caller(true))
	fmt.Println(Caller(false))
}
