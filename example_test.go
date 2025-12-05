package main

import "testing"

func TestParseInput(t *testing.T) {
	testCases := []struct {
		input string
		want  string
	}{}
	for _, tc := range testCases {
		t.Log(tc.want)
	}
}
