package main

import (
	"reflect"
	"strconv"
	"testing"
)

func fizzBuzz_old(n int) []string {
	var ss []string
	for i := 1; i <= n; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			ss = append(ss, "FizzBuzz")
		case i%3 == 0:
			ss = append(ss, "Fizz")
		case i%5 == 0:
			ss = append(ss, "Buzz")
		default:
			ss = append(ss, strconv.Itoa(i))
		}
	}
	return ss
}

func fizzBuzz(n int) []string {
	var ss []string
	for i := 1; i <= n; i++ {
		s := ""
		switch {
		case i%3 == 0 && i%5 == 0:
			s = "FizzBuzz"
		case i%3 == 0:
			s = "Fizz"
		case i%5 == 0:
			s = "Buzz"
		default:
			s = strconv.Itoa(i)
		}
		ss = append(ss, s)
	}
	return ss
}

func Test_fizzBuzz(t *testing.T) {
	tests := []struct {
		in  int
		out []string
	}{
		{
			15,
			[]string{
				"1",
				"2",
				"Fizz",
				"4",
				"Buzz",
				"Fizz",
				"7",
				"8",
				"Fizz",
				"Buzz",
				"11",
				"Fizz",
				"13",
				"14",
				"FizzBuzz",
			},
		},
	}
	for _, tt := range tests {
		got := fizzBuzz(tt.in)
		if !reflect.DeepEqual(got, tt.out) {
			t.Errorf("got: %v, want: %v", got, tt.out)
		}
	}
}
