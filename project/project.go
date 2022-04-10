package project

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"

	. "github.com/digisan/gotk/strs"
	tk "github.com/digisan/gotk/track"
)

func PrjName(dfltName string) (string, bool) {
	const check = "/.git"
NEXT:
	for i := 1; i < 64; i++ {
		for _, ln := range SplitLn(tk.TrackCaller(i)) {
			if strings.HasPrefix(ln, "/") {
				ln = TrimTailFromLast(ln, ":")
			AGAIN:
				dir := filepath.Dir(ln)
				if dir == "/" {
					continue NEXT
				}
				_, err := os.Stat(dir + check)
				if os.IsNotExist(err) {
					ln = dir
					goto AGAIN
				} else {
					return filepath.Base(dir), true
				}
			}
		}
	}
	return dfltName, false
}

func GitTag() (tag string, err error) {
	defer func() {
		if r := recover(); r != nil {
			tag, err = "", fmt.Errorf("%v", r)
		}
	}()

	// run git
	cmd := exec.Command("bash", "-c", "git describe --tags")
	output, err := cmd.Output()
	if outstr := strings.Trim(string(output), " \n\t"); outstr != "" {
		return strings.Split(outstr, "-")[0], nil
	}
	return "", nil
}

func GitVer(dfltVer string) (string, bool) {
	tag, err := GitTag()
	if err != nil {
		return dfltVer, false
	}
	if r := regexp.MustCompile(`^v\d+\.\d+\.\d+$`); r.MatchString(tag) {
		return tag, true
	}
	return dfltVer, false
}
