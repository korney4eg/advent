package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsIdValid(t *testing.T) {
	testCases := []struct {
		input string
		want  bool
		n     int
	}{
		{
			input: "22",
			want:  false,
			n:     2,
		},
		{
			input: "11",
			want:  false,
			n:     2,
		},
		{
			input: "12",
			want:  true,
			n:     2,
		},
		{
			input: "111",
			want:  true,
			n:     2,
		},
		{
			input: "111",
			want:  false,
			n:     3,
		},
	}
	for _, tc := range testCases {
		got := isIdValid(tc.input, tc.n)
		assert.Equal(t, tc.want, got, "is id valid for", tc.input, tc.n)
	}
}

func TestGetNextInvalidId(t *testing.T) {
	testCases := []struct {
		input string
		want  string
		n     int
	}{
		{
			input: "1",
			want:  "11",
			n:     2,
		},
		{
			input: "10",
			want:  "11",
			n:     2,
		},
		{
			input: "11",
			want:  "22",
			n:     2,
		},
		{
			input: "22",
			want:  "33",
			n:     2,
		},
		{
			input: "95",
			want:  "99",
			n:     2,
		},
		{
			input: "99",
			want:  "1010",
			n:     2,
		},
		{
			input: "998",
			want:  "1010",
			n:     2,
		},
		{
			input: "1188511880",
			want:  "1188511885",
			n:     2,
		},
		{
			input: "222220",
			want:  "222222",
			n:     2,
		},
		{
			input: "446443",
			want:  "446446",
			n:     2,
		},
		{
			input: "38593856",
			want:  "38593859",
			n:     2,
		},
		{
			input: "17",
			want:  "22",
			n:     2,
		},
		{
			input: "79",
			want:  "88",
			n:     2,
		},
		{
			input: "99",
			want:  "111",
			n:     3,
		},
		{
			input: "280",
			want:  "333",
			n:     3,
		},
	}
	for _, tc := range testCases {
		got := getNextInvalidId(tc.input, tc.n)
		assert.Equal(t, tc.want, got, "next invalid id", tc.input, tc.n)
	}
}

func TestListAllInvalidIds(t *testing.T) {
	testCases := []struct {
		input string
		want  int
		n     int
	}{
		{
			input: "11-22",
			want:  2,
			n:     2,
		},
		{
			input: "95-115",
			want:  1,
			n:     2,
		},
		{
			input: "998-1012",
			want:  1,
			n:     2,
		},
		{
			input: "1188511880-1188511890",
			want:  1,
			n:     2,
		},
		{
			input: "222220-222224",
			want:  1,
			n:     2,
		},
		{
			input: "1698522-1698528",
			want:  0,
			n:     2,
		},
		{
			input: "446443-446449",
			want:  1,
			n:     2,
		},
		{
			input: "38593856-38593862",
			want:  1,
			n:     2,
		},
		{
			input: "95-115",
			want:  1,
			n:     3,
		},
	}
	for _, tc := range testCases {
		got := len(listAllInvalidIds(tc.input, tc.n))
		t.Log(listAllInvalidIds(tc.input, tc.n))
		assert.Equal(t, tc.want, got, "cound invalid ids", tc.input)
	}
}

func TestGetAllInvalidIds(t *testing.T) {
	testCases := []struct {
		input string
		want  []int
	}{
		{
			input: "95-115",
			want:  []int{99, 111},
		},
		{
			input: "79-106",
			want:  []int{88, 99},
		},
		{
			input: "280-392",
			want:  []int{333},
		},
	}
	for _, tc := range testCases {
		got := getAllInvalidIds(tc.input)
		assert.Equal(t, tc.want, got, "cound invalid ids", tc.input)
	}
}
