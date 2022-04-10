package misc

import (
	"testing"
	"time"
)

func TestTrackTime(t *testing.T) {
	type args struct {
		start time.Time
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			TrackTime(tt.args.start)
		})
	}
}
