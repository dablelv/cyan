package util

import "testing"

func TestIsReservedIP(t *testing.T) {
	type args struct {
		ip string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"loopback", args{"127.0.0.1"}, 1},
		{"multicast", args{"10.255.255.255"}, 1},
		{"A private", args{"10.0.0.1"}, 1},
		{"B private", args{"172.16.0.1"}, 1},
		{"C private", args{"192.168.0.1"}, 1},
		{"public", args{"1.0.0.1"}, 0},
		{"invalid ipv4", args{"192.168.0.256"}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsReservedIP(tt.args.ip); got != tt.want {
				t.Errorf("IsReservedIP() = %v, want %v", got, tt.want)
			}
		})
	}
}
