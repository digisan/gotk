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

func TestModifyOriginOrIP(t *testing.T) {

	s := `
	http://127.0.0.1:3000/api/
	https://127.0.0.1:3000/api/
	http://localhost:3001/api/
	https://localhost:3001/api/
	localhost:3002/api/
	127.0.0.1:3002/api/
	`

	// fmt.Println(ModifyOriginIP(s, "localhost", "my.example.com", "", -1, true, true, true))
	// fmt.Println(ModifyOriginIP(s, "localhost", "my.example.com", "", -1, true, true, false))
	// fmt.Println(ModifyOriginIP(s, "localhost", "test://my.example.com", "", -1, false, true, false))
	fmt.Println(ModifyOriginIP(s, "localhost", "my.example.com:1234", "", -1, true, false, true))
}
