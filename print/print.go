package print

import "fmt"

func FnFixedPrintf() func(format string, a ...any) (n int, err error) {
	fmt.Print("\033[s")
	return func(format string, a ...any) (n int, err error) {
		fmt.Print("\033[u\033[k")
		return fmt.Printf(format, a...)
	}
}
