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
	fmt.Println(ChangeLocalUrlPort("./test-sample.ejs", 3000, 5000, false, false))
}

func TestIPLoc2Pub(t *testing.T) {
	fmt.Println(LocIP2PubIP("./test-sample.ejs", true, false))
}

func TestLhTo127(t *testing.T) {
	fmt.Println(LocalhostToIP127("./test-sample.ejs", true, false))
}

func Test127ToLh(t *testing.T) {
	fmt.Println(IP127ToLocalhost("./test-sample.ejs", false, false))
}
