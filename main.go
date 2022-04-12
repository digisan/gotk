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

	fmt.Println(track.TrackCaller(0))
	fmt.Println("----------------")

}
