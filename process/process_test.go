package process

import (
	"fmt"
	"testing"
	"time"
)

func TestGetRunningPID(t *testing.T) {
	fmt.Println(GetRunningPID("~/Desktop/OTF/otf-reader/cmd/otf-reader/otf-reader"))
	fmt.Println(time.Now())
}
