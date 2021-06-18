package main

import (
	"fmt"
	"testing"
)

/*
	## summary
*/
func addBinary(a string, b string) string {
	ret, sum := []byte{}, 0
	for i, j := len(a)-1, len(b)-1; i >= 0 || j >= 0 || sum != 0; i, j = i-1, j-1 {
		if i >= 0 {
			sum += int(a[i] - '0')
		}
		if j >= 0 {
			sum += int(b[j] - '0')
		}

		ret = append(ret, byte(sum)%2+'0')
		sum /= 2
	}
	for i, j := 0, len(ret)-1; i < j; i, j = i+1, j-1 {
		ret[i], ret[j] = ret[j], ret[i]
	}
	return string(ret)
}

func TestAddBinary(t *testing.T) {
	tests := []struct {
		in  string
		in2 string
		out string
	}{
		{
			"11", "1", "100",
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := addBinary(tt.in, tt.in2)
			if got != tt.out {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
