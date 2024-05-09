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

	ColorPrintf("original color - %s - %s\n", B("hello"), R(CB(1.23)))

	fPrintf := FnFixedPrintf()
	for i := range 3 {
		s := strings.Repeat("-", i)
		fPrintf("%v %s\n", time.Now().Format("2006-01-02 15:04:05"), s)
		time.Sleep(time.Duration(1000) * time.Millisecond)
	}
	fmt.Println("Done")

	fcPrintf := FnFixedColorPrintf()
	for i := range 5 {
		s := strings.Repeat("-", i)
		fcPrintf("%v %s\n", time.Now().Format("2006-01-02 15:04:05"), G(s))
		time.Sleep(time.Duration(1000) * time.Millisecond)
	}
	fmt.Println("Done")
}
