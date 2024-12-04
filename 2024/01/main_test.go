package main

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	input string
	want  []int
}

func TestGetDistances(t *testing.T) {
	testCases := []testCase{
		{
			input: `1 6
2 5
3 4`,
			want: []int{3, 3, 3},
		},
		{
			input: `3   4
4   3
2   5
1   3
3   9
3   3`,
			want: []int{2, 1, 0, 1, 2, 5},
		},
	}

	for _, tescCase := range testCases {
		leftList, rightList := getInput(tescCase.input)
		slices.Sort(leftList)
		slices.Sort(rightList)
		t.Log("leftList", leftList)
		t.Log("rightList", rightList)
		distances := getDistances(leftList, rightList)
		assert.Equal(t, tescCase.want, distances, "should be equal")
	}
}
func TestGetSimilarities(t *testing.T) {
	testCases := []testCase{
		{
			input: `1 6
2 5
3 4`,
			want: []int{0, 0, 0},
		},
		{
			input: `3   4
4   3
2   5
1   3
3   9
3   3`,
			want: []int{9, 4, 0, 0, 9, 9},
		},
	}

	for _, tescCase := range testCases {
		leftList, rightList := getInput(tescCase.input)
		t.Log("leftList", leftList)
		t.Log("rightList", rightList)
		distances := getSimilarities(leftList, rightList)
		assert.Equal(t, tescCase.want, distances, "should be equal")
	}
}
