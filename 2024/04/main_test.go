package main

import "testing"

//	func TestFindSmallField(t *testing.T) {
//		smallField := `XMAS
//
// XMAS
// XMAS
// XMAS`
//
//		f := field{}
//		f.read(smallField)
//		t.Log(f.board)
//		got := findInSmallField(f.board, 0)
//		if got != 4 {
//			t.Error("Expected 4", got)
//		}
//	}
func TestFindInField(t *testing.T) {
	testCases := []struct {
		input       string
		wantNorm    int
		wantDiag    int
		wantRotated int
	}{
		{
			input: `XMAS
XMAS
XMAS
XMAS`,
			wantNorm: 4,
			wantDiag: 2,
		},
		{
			input: `SAMX
XAAS
XMXS
XMSS`,
			wantNorm: 1,
		},
		{
			input: `SAMX
XAAS
XMMS
XMSX`,
			wantNorm: 1,
			wantDiag: 1,
		},
		{
			input: `XXXX
MMMM
AAAA
SSSS`,
			wantRotated: 4,
			wantDiag:    2,
		},
	}
	for _, tc := range testCases {
		norm := findInField(tc.input)
		diags := countInDiagonals(tc.input)
		rotated := findInField(rotate(tc.input))
		if norm != tc.wantNorm {
			t.Errorf("Expected %d, got %d", tc.wantNorm, norm)
		}
		if diags != tc.wantDiag {
			t.Errorf("Expected %d, got %d", tc.wantDiag, diags)
		}
		if rotated != tc.wantRotated {
			t.Errorf("Expected %d, got %d", tc.wantRotated, rotated)
		}
	}
}
