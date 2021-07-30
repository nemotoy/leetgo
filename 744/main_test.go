package tree

import (
	"fmt"
	"reflect"
	"testing"
)

/*
	## summary
*/
func nextGreatestLetter(letters []byte, target byte) byte {
	var count [26]bool
	for _, r := range letters {
		count[r-'a'] = true
	}
	for {
		target++
		if target > 'z' {
			target = 'a'
		}
		if count[target-'a'] {
			return target
		}
	}
}

func nextGreatestLetter2(letters []byte, target byte) byte {
	for _, letter := range letters {
		if letter > target {
			return letter
		}
	}
	return letters[0]
}

func TestNextGreatestLetter(t *testing.T) {
	tests := []struct {
		in  []byte
		in2 byte
		out byte
	}{
		{[]byte{'c', 'f', 'j'}, 'a', 'c'},
		{[]byte{'c', 'f', 'j'}, 'c', 'f'},
		{[]byte{'c', 'f', 'j'}, 'd', 'f'},
		{[]byte{'c', 'f', 'j'}, 'g', 'j'},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", 'a'), func(t *testing.T) {
			got := nextGreatestLetter2(tt.in, tt.in2)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
