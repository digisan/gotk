package tpt

import (
	"fmt"
	"image"
	"testing"
)

func TestMkSet(t *testing.T) {
	pts := []image.Point{
		{X: 1, Y: 2},
		{X: 1, Y: 3},
		{X: 3, Y: 2},
		{X: 1, Y: 2},
		{X: 1, Y: 3},
		{X: 1, Y: 4},
	}
	fmt.Println(MkSet(pts...))
}
