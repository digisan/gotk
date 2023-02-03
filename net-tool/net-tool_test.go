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

func TestReplacePort4LocalUrl(t *testing.T) {
	fmt.Println(ReplacePort4LocalUrl(-1, 5000, false, false, "./test-sample.ejs"))
}

func TestLoc127To(t *testing.T) {
	fmt.Println(Loc127To(false, true, "134.22.33.211", false, true, "./test-sample.ejs"))
}

func TestLocIP2PubIP(t *testing.T) {
	fmt.Println(LocIP2PubIP(false, true, "./test-sample.ejs"))
}

func TestModifyOriginOrIP(t *testing.T) {
	// fmt.Println(ModifyOriginOrIP("localhost", "127.0.0.1", false, true, false, "./test-sample.ejs"))
	// fmt.Println(ModifyOriginOrIP("127.0.0.1", "localhost", false, true, false, "./test-sample.ejs"))
	// fmt.Println(ModifyOriginOrIP("localhost", "my.example.com", false, false, true, "./test-sample.ejs"))
	// fmt.Println(ModifyOriginOrIP("localhost", "my.example.com", false, false, false, "./test-sample.ejs"))
	fmt.Println(ModifyOriginOrIP("localhost", "my.example.com", true, false, false, "./test-sample.ejs"))
}
