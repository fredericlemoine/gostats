package gostats

import (
	"testing"
)

func TestMedian_float(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		a    []float64
		want float64
	}{
		{name: "median", a: []float64{1, 11, 15, 19, 20, 24, 28, 34, 37, 47, 50, 61}, want: 26},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Median_float(tt.a)
			// TODO: update the condition below to compare got with tt.want.
			if got != tt.want {
				t.Errorf("Median_float() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQ1_float(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		a    []float64
		want float64
	}{
		{name: "median", a: []float64{1, 11, 15, 19, 20, 24, 28, 34, 37, 47, 50, 61}, want: 18},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Q1_float(tt.a)
			// TODO: update the condition below to compare got with tt.want.
			if got != tt.want {
				t.Errorf("Median_float() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQ3_float(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		a    []float64
		want float64
	}{
		{name: "median", a: []float64{1, 11, 15, 19, 20, 24, 28, 34, 37, 47, 50, 61}, want: 39.5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Q3_float(tt.a)
			// TODO: update the condition below to compare got with tt.want.
			if got != tt.want {
				t.Errorf("Median_float() = %v, want %v", got, tt.want)
			}
		})
	}
}
