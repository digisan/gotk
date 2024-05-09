package main

import (
	"fmt"
	"strings"
	"time"

	. "github.com/digisan/gotk/print"
	. "github.com/digisan/gotk/project"
	"github.com/digisan/gotk/track"
)

func main() {

	pn, ok := PrjName("TestProject")
	fmt.Println(pn, ok)

	pt, err := GitTag()
	fmt.Println(pt, err)

	pv, ok := GitVer("v0.0.0")
	fmt.Println(pv, ok)

	// fmt.Println(track.CallerDescription(0))
	// fmt.Println("----------------")

	fmt.Println(track.CallerFn(true, 0))
	fmt.Println("----------------")

	fmt.Println(track.CallerFn(true, 1))
	fmt.Println("----------------")

	// track.TrackAll()

	FixedPrintf := FnFixedPrintf()
	for i := range 30 {
		s := strings.Repeat("-", i)
		FixedPrintf("%v %s\n", time.Now().Format("2006-01-02 15:04:05"), s)
		time.Sleep(time.Duration(1000) * time.Millisecond)
	}
	fmt.Println("Done")

}
