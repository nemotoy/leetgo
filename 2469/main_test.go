package main

import (
	"reflect"
	"testing"
)

func convertTemperature(celsius float64) []float64 {
	var kelvin, fahrenheit float64
	kelvin = celsius + 273.15
	fahrenheit = celsius*1.80 + 32.00
	return []float64{kelvin, fahrenheit}
}

func TestConvertTemperature(t *testing.T) {
	tests := []struct {
		in  float64
		out []float64
	}{
		{
			36.50, []float64{309.65000, 97.70000},
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := convertTemperature(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
