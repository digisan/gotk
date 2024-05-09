package print

import (
	"fmt"

	clr "github.com/gookit/color"
)

var (
	W = clr.FgWhite.Render
	R = clr.FgRed.Render
	G = clr.FgGreen.Render
	B = clr.FgLightBlue.Render
	C = clr.FgLightCyan.Render
	M = clr.FgMagenta.Render
	Y = clr.FgYellow.Render

	WB = clr.BgWhite.Render
	RB = clr.BgRed.Render
	GB = clr.BgGreen.Render
	BB = clr.BgLightBlue.Render
	CB = clr.BgLightCyan.Render
	MB = clr.BgMagenta.Render
	YB = clr.BgYellow.Render
)

func ColorPrintf(format string, a ...any) (n int, err error) {
	return fmt.Printf(format, a...)
}

func FnFixedPrintf() func(format string, a ...any) (n int, err error) {
	fmt.Print("\033[s")
	return func(format string, a ...any) (n int, err error) {
		fmt.Print("\033[u\033[k")
		return fmt.Printf(format, a...)
	}
}

func FnFixedColorPrintf() func(format string, a ...any) (n int, err error) {
	fmt.Print("\033[s")
	return func(format string, a ...any) (n int, err error) {
		fmt.Print("\033[u\033[k")
		return ColorPrintf(format, a...)
	}
}
