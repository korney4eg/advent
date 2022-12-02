package main

import "testing"
	type testCase struct {
		me             string
		enemy          string
		expectedResult int
	}

func TestGetVictoryPoints(t *testing.T)

	cases := []*testCase{
		&testCase{
			me:             "p",
			enemy:          "p",
			expectedResult: 3,
		},
	}

	for _, c := range cases {
		result := GetVictoryPoints(c.enemy, c.me)
		if result != c.expectedResult {
			t.Errorf("expected: %d\n     got:%d\n", c.expectedResult, result)
		}
	}

}
