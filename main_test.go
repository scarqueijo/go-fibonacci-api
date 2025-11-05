package main

import "testing"

func TestFibonacci(t *testing.T) {
	tests := []struct {
		input int
		want  []int
	}{
		{1, []int{1}},
		{2, []int{1, 1}},
		{5, []int{1, 1, 2, 3, 5}},
		{7, []int{1, 1, 2, 3, 5, 8, 13}},
	}

	for _, tt := range tests {
		got := Fibonacci(tt.input)
		for i := range got {
			if got[i] != tt.want[i] {
				t.Errorf("Fibonacci(%d) = %v; want %v", tt.input, got, tt.want)
				break
			}
		}
	}
}
