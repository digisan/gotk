package process

import (
	"log"
	"os/exec"
	"path/filepath"

	"github.com/digisan/gotk/io"
	"github.com/digisan/gotk/slice/ts"
)

// GetRunningPID:
func GetRunningPID(pathOfExe string) (pidGrp []string) {
	abspath, err := io.AbsPath(pathOfExe, true)
	if err != nil {
		log.Fatalf("%v", err)
	}

	dir, exe := filepath.Dir(abspath), filepath.Base(abspath)

	// ps -Af | grep ***
	out, err := exec.Command("/bin/sh", "-c", "ps -Af | grep "+exe).CombinedOutput()
	if fSf("%v", err) == "exit status 1" {
		return
	}
	if err != nil {
		log.Fatalf("%v", err)
	}

	pidGrpGrep := []string{}
	io.StrLineScan(string(out), func(ln string) (bool, string) {
		I := 0
		for _, seg := range sSplit(ln, " ") {
			if len(seg) > 0 {
				I++
			}
			if I == 2 {
				pidGrpGrep = append(pidGrpGrep, seg)
				break
			}
		}
		return true, ""
	}, "")

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
		out, err = exec.Command("/bin/sh", "-c", "pwdx "+pid).CombinedOutput()
		if fSf("%v", err) == "exit status 1" {
			return
		}
		if err != nil {
			log.Fatalf("%v", err)
		}

		procpath := sSplit(sTrim(string(out), " \t\r\n"), ": ")[1]
		if dir == procpath && ts.In(pid, pidGrpPGrep...) {
			pidGrp = append(pidGrp, pid)
		}
	}

	return

}

// ExistRunningPS:
func ExistRunningPS(pathOfExe string) bool {
	return len(GetRunningPID(pathOfExe)) > 0
}
