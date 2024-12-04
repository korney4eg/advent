package main

import "testing"

func TestIsLevelSafe(t *testing.T) {
	testCases := []struct {
		input string
		want  bool
	}{
		{
			input: "7 6 4 2 1",
			want:  true,
		},
		{
			input: "1 2 7 8 9",
			want:  false,
		},
		{
			input: "1 3 2 4 5",
			want:  false,
		},
	}
	for _, tc := range testCases {
		rep := getInput(tc.input)[0]
		t.Log(rep)
		got, ops := rep.isLevelSafe()
		t.Log("ops: ", ops)
		if got != tc.want {
			t.Errorf("isLevelSafe(%q) = %t; want %t", tc.input, got, tc.want)
		}
	}
}

func TestCountSafeLevels(t *testing.T) {
	testCases := []struct {
		input string
		want  int
	}{
		{
			input: `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`,
			want: 2,
		},
	}
	for _, tc := range testCases {
		reports := getInput(tc.input)
		t.Log(reports)
		got, ops := countSafeLevels(reports)
		t.Log("ops: ", ops)
		if got != tc.want {
			t.Errorf("countSafeLevels(%v) = %v; want %v", tc.input, got, tc.want)
		}
	}
}

func TestIsLevelSafeWithDamper(t *testing.T) {
	testCases := []struct {
		input string
		want  bool
	}{
		{
			input: "1 3 2 4 5",
			want:  true,
		},
		{
			input: "1 2 7 8 9",
			want:  false,
		},
		{
			input: "8 6 4 4 1",
			want:  true,
		},
		{
			input: "88 90 85 83 80 79 77 71",
			want:  false,
		},
	}
	for _, tc := range testCases {
		rep := getInput(tc.input)[0]
		t.Log(rep)
		got, ops := rep.isLevelSafeWithDamper()
		t.Log("ops: ", ops)
		if got != tc.want {
			t.Errorf("isLevelSafe(%q) = %t; want %t", tc.input, got, tc.want)
		}
	}
}

func TestCountSafeLevelsWithDampers(t *testing.T) {
	testCases := []struct {
		input string
		want  int
	}{
		{
			input: `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`,
			want: 4,
		},
		{
			input: `19 20 21 23 24 25 28 26
56 58 60 63 66 69 69
3 6 7 8 11 15
50 53 55 58 63
39 41 42 45 42 44 46
22 25 27 26 25`,
			want: 4,
		},
	}
	for _, tc := range testCases {
		reports := getInput(tc.input)
		t.Log(reports)
		got, ops := countSafeLevelsWithDamper(reports)
		t.Log("ops: ", ops)
		if got != tc.want {
			t.Errorf("countSafeLevels(%v) = %v; want %v", tc.input, got, tc.want)
		}
	}
}
