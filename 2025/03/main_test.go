package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// func TestFindJoltage(t *testing.T) {
// 	testCases := []struct {
// 		input string
// 		want  string
// 	}{
// 		{
// 			input: "987654321111111",
// 			want:  "98",
// 		},
// 		{
// 			input: "811111111111119",
// 			want:  "89",
// 		},
// 		{
// 			input: "234234234234278",
// 			want:  "78",
// 		},
// 		{
// 			input: "818181911112111",
// 			want:  "92",
// 		},
// 	}
// 	for _, tc := range testCases {
// 		got := findJoltage(tc.input, 2)
// 		assert.Equal(t, tc.want, got, "joltage check")
// 	}
// }

func TestFindJoltageLong(t *testing.T) {
	testCases := []struct {
		input string
		want  string
	}{
		{
			input: "987654321111111",
			want:  "987654321111",
		},
		{
			input: "811111111111119",
			want:  "811111111119",
		},
		{
			input: "234234234234278",
			want:  "434234234278",
		},
		{
			input: "818181911112111",
			want:  "888911112111",
		},
	}
	for _, tc := range testCases {
		got := findJoltage(tc.input, 12)
		assert.Equal(t, tc.want, got, "joltage check")
	}

}
