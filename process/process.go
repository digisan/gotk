package process

import (
	"log"
	"os/exec"
	"path/filepath"

	"github.com/digisan/gotk/io"
)

// GetRunningPID:
func GetRunningPID(pathOfExe string) (pidGrp []string) {
	abspath, err := filepath.Abs(pathOfExe)
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

	pidGrpTemp := []string{}
	io.StrLineScan(string(out), func(ln string) (bool, string) {
		I := 0
		for _, seg := range sSplit(ln, " ") {
			if len(seg) > 0 {
				I++
			}
			if I == 2 {
				pidGrpTemp = append(pidGrpTemp, seg)
				break
			}
		}
		return true, ""
	}, "")

	for _, pid := range pidGrpTemp {
		out, err := exec.Command("/bin/sh", "-c", "pwdx "+pid).CombinedOutput()
		if fSf("%v", err) == "exit status 1" {
			return
		}
		if err != nil {
			log.Fatalf("%v", err)
		}

		procpath := sSplit(sTrim(string(out), " \t\r\n"), ": ")[1]
		if dir == procpath {
			pidGrp = append(pidGrp, pid)
		}
	}

	//

	// dir, exe := filepath.Dir(abspath), filepath.Base(abspath)
	// out, err := exec.Command("/bin/sh", "-c", "pgrep "+exe).CombinedOutput()
	// if fSf("%v", err) == "exit status 1" {
	// 	return
	// }
	// if err != nil {
	// 	log.Fatalf("%v", err)
	// }

	// outstr := sTrim(string(out), " \t\r\n")
	// for _, pid := range sSplit(outstr, "\n") {
	// 	out, err := exec.Command("/bin/sh", "-c", "pwdx "+pid).CombinedOutput()
	// 	if fSf("%v", err) == "exit status 1" {
	// 		return
	// 	}
	// 	if err != nil {
	// 		log.Fatalf("%v", err)
	// 	}

	// 	outstr := sTrim(string(out), " \t\r\n")
	// 	procpath := sSplit(outstr, ": ")[1]
	// 	if dir == procpath {
	// 		pidGrp = append(pidGrp, pid)
	// 	}
	// }

	return
}

// ExistRunningPS:
func ExistRunningPS(pathOfExe string) bool {
	return len(GetRunningPID(pathOfExe)) > 0
}
