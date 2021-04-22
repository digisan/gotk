package process

import (
	"fmt"
	"testing"
	"time"
)

func TestGetRunningPID(t *testing.T) {
	fmt.Println(GetRunningPID("/home/qmiao/Desktop/n3/nats-streaming-server-v0.21.2-linux-amd64/nats-streaming-server"))
	fmt.Println(time.Now())
}
