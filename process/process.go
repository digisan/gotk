package process

import (
	"log"
	"os/exec"
	"path/filepath"
)

// GetRunningPID:
func GetRunningPID(pathOfExe string) (pidGrp []string) {
	abspath, err := filepath.Abs(pathOfExe)
	if err != nil {
		log.Fatalf("%v", err)
	}

	dir, exe := filepath.Dir(abspath), filepath.Base(abspath)
	out, err := exec.Command("/bin/sh", "-c", "pgrep "+exe).CombinedOutput()
	if fSf("%v", err) == "exit status 1" {
		return
	}
	if err != nil {
		log.Fatalf("%v", err)
	}

	outstr := sTrim(string(out), " \t\r\n")
	for _, pid := range sSplit(outstr, "\n") {
		out, err := exec.Command("/bin/sh", "-c", "pwdx "+pid).CombinedOutput()
		if fSf("%v", err) == "exit status 1" {
			return
		}
		if err != nil {
			log.Fatalf("%v", err)
		}

		outstr := sTrim(string(out), " \t\r\n")
		procpath := sSplit(outstr, ": ")[1]
		if dir == procpath {
			pidGrp = append(pidGrp, pid)
		}
	}
	return
}

// ExistRunningPS:
func ExistRunningPS(pathOfExe string) bool {
	return len(GetRunningPID(pathOfExe)) > 0
}
