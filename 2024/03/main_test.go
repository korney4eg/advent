package main

import "testing"

func TestFindCorrectInstructions(t *testing.T) {
	testCases := []struct {
		input string
		want  []string
	}{
		{
			input: "xmul(2,4)%",
			want:  []string{"mul(2,4)"},
		},
		{
			input: "+mul(32,64]",
			want:  []string{},
		},
		{
			input: "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))",
			want:  []string{"mul(2,4)", "mul(5,5)", "mul(11,8)", "mul(8,5)"},
		},
	}
	for _, tc := range testCases {
		got := findCorrectInstructions(tc.input)
		if len(got) != len(tc.want) {
			t.Fatalf("findCorrectInstructions(%q) = %v; want %v", tc.input, got, tc.want)
		}
		for i := range got {
			if got[i] != tc.want[i] {
				t.Fatalf("findCorrectInstructions(%q) = %v; want %v", tc.input, got, tc.want)
			}
		}
	}
}

func TestMulAll(t *testing.T) {
	testCases := []struct {
		input string
		want  int
	}{
		{
			input: "xmul(2,4)%",
			want:  8,
		},
		{
			input: "+mul(32,64]",
			want:  0,
		},
		{
			input: "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))",
			want:  161,
		},
	}
	for _, tc := range testCases {
		muls := findCorrectInstructions(tc.input)
		got := mulAll(muls)
		if got != tc.want {
			t.Fatalf("MulAll(%q) = %v; want %v", tc.input, got, tc.want)
		}
	}
}

func TestFindEnabledInstructions(t *testing.T) {
	testCases := []struct {
		input string
		want  []string
	}{
		{
			input: "xmul(2,4)%",
			want:  []string{"mul(2,4)"},
		},
		{
			input: "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)",
			want:  []string{"mul(2,4)"},
		},
		{
			input: "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))",
			want:  []string{"mul(2,4)", "mul(8,5)"},
		},
	}
	for _, tc := range testCases {
		got := findEnabledInstructions(tc.input)
		if len(got) != len(tc.want) {
			t.Fatalf("findEnabledInstructions(%q) = %v; want %v", tc.input, got, tc.want)
		}
		for i := range got {
			if got[i] != tc.want[i] {
				t.Fatalf("findEnabledInstructions(%q) = %v; want %v", tc.input, got, tc.want)
			}
		}
	}
}
func TestEnabledMulAll(t *testing.T) {
	testCases := []struct {
		input string
		want  int
	}{
		{
			input: "xmul(2,4)%",
			want:  8,
		},
		{
			input: "+mul(32,64]",
			want:  0,
		},
		{
			input: "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))",
			want:  48,
		},
	}
	for _, tc := range testCases {
		muls := findEnabledInstructions(tc.input)
		got := mulAll(muls)
		if got != tc.want {
			t.Fatalf("MulAll(%q) = %v; want %v", tc.input, got, tc.want)
		}
	}
}
