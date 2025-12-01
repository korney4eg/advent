package main

import "testing"

func TestProcess(t *testing.T) {
	dial := make([]int, 100)
	testCases := []struct {
		input int
		want  int
	}{
		{
			input: 49,
			want:  99,
		},
		{
			input: 50,
			want:  0,
		},
		{
			input: 150,
			want:  0,
		},
		{
			input: -50,
			want:  0,
		},
		{
			input: -52,
			want:  98,
		},
	}
	for _, tc := range testCases {
		if got := process(dial, 50, []int{tc.input})[0]; got != tc.want {
			t.Errorf("for input %d. Wanted %d, got %d", tc.input, tc.want, got)
		}
	}
}
func TestProcess2(t *testing.T) {
	dial := make([]int, 100)
	testCases := []struct {
		input int
		want  int
	}{
		{
			input: 49,
			want:  99,
		},
		{
			input: 50,
			want:  0,
		},
		{
			input: 150,
			want:  0,
		},
		{
			input: -50,
			want:  0,
		},
		{
			input: -52,
			want:  98,
		},
	}
	for _, tc := range testCases {
		if got := process(dial, 50, []int{tc.input})[0]; got != tc.want {
			t.Errorf("for input %d. Wanted %d, got %d", tc.input, tc.want, got)
		}
	}
}
