package nettool

import (
	"fmt"
	"testing"
)

func TestLocalIP(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
		{
			want: "192.168.31.227",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LocalIP(); got != tt.want {
				t.Errorf("LocalIP() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPublicIP(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
		{
			want: "115.70.114.243",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PublicIP(); got != tt.want {
				t.Errorf("PublicIP() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChangeLocalPort(t *testing.T) {
	fmt.Println(ChangeLocalUrlPort(3000, 5000, false, false, "./test-sample.ejs"))
}

func TestIPLoc2Pub(t *testing.T) {
	fmt.Println(ChangeLocalIP(false, true, false, "134.22.33.211", "./test-sample.ejs"))
}

func TestLhTo127(t *testing.T) {
	fmt.Println(LocalhostToIP127(false, true, "./test-sample.ejs"))
}

func Test127ToLh(t *testing.T) {
	fmt.Println(IP127ToLocalhost(false, true, "./test-sample.ejs"))
}
