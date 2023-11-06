package main

import (
	"fmt"

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

}
