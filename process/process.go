package process

import (
	"fmt"
	"log"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

	. "github.com/digisan/go-generics/v2"
	fd "github.com/digisan/gotk/file-dir"
	"github.com/digisan/gotk/strs"
)

var (
	fSf    = fmt.Sprintf
	sTrim  = strings.Trim
	sSplit = strings.Split
)

// GetRunningPID:
func GetRunningPID(pathOfExe string) (pidGrp []int) {
	absPath, err := fd.AbsPath(pathOfExe, true)
	if err != nil {
		log.Fatalf("%v", err)
	}

	dir, exe := filepath.Dir(absPath), filepath.Base(absPath)

	// ps -Af | grep ***
	out, err := exec.Command("/bin/sh", "-c", "ps -Af | grep "+exe).CombinedOutput()
	if fSf("%v", err) == "exit status 1" {
		return
	}
	if err != nil {
		log.Fatalf("%v", err)
	}

	pidGrpGrep := []int{}
	strs.StrLineScan(string(out), func(ln string) (bool, string) {
		I := 0
		for _, seg := range sSplit(ln, " ") {
			if len(seg) > 0 {
				I++
			}
			if I == 2 {
				pid, err := strconv.Atoi(seg)
				if err != nil {
					log.SetFlags(log.Lshortfile | log.LstdFlags)
					log.Fatalln(err)
				}
				pidGrpGrep = append(pidGrpGrep, pid)
				break
			}
		}
		return true, ""
	})

	// pgrep ***
	if len(exe) > 15 {
		exe = exe[:15]
	}
	out, err = exec.Command("/bin/sh", "-c", "pgrep "+exe).CombinedOutput()
	if fSf("%v", err) == "exit status 1" {
		return
	}
	if err != nil {
		log.Fatalf("%v", err)
	}
	pidGrpPGrep := sSplit(sTrim(string(out), " \t\r\n"), "\n")

	// check... dirOfExe & in pgrep
	for _, pid := range pidGrpGrep {
		out, err = exec.Command("/bin/sh", "-c", "pwdx "+fmt.Sprint(pid)).CombinedOutput()
		if fSf("%v", err) == "exit status 1" {
			return
		}
		if err != nil {
			log.Fatalf("%v", err)
		}

		procPath := sSplit(sTrim(string(out), " \t\r\n"), ": ")[1]
		if dir == procPath && In(fmt.Sprint(pid), pidGrpPGrep...) {
			pidGrp = append(pidGrp, pid)
		}
	}

	return
}

// ExistRunningPS:
func ExistRunningPS(pathOfExe string) bool {
	return len(GetRunningPID(pathOfExe)) > 0
}
