package nettool

import (
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
