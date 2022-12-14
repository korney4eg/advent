package signals

import "testing"

func TestIsRightOrder(t *testing.T) {
	testCases := []struct {
		signal1        string
		signal2        string
		expectedResult bool
	}{
		{
			signal1:        `[1,1,3,1,1]`,
			signal2:        `[1,1,5,1,1]`,
			expectedResult: true,
		},
		{
			signal1:        `[[1],[2,3,4]]`,
			signal2:        `[[1],4]`,
			expectedResult: true,
		},
		{
			signal1:        `[9]`,
			signal2:        `[[8,7,6]]`,
			expectedResult: false,
		},
		{
			signal1:        `[[4,4],4,4]`,
			signal2:        `[[4,4],4,4,4]`,
			expectedResult: true,
		},
		{
			signal1:        `[7,7,7,7]`,
			signal2:        `[7,7,7]`,
			expectedResult: false,
		},
		{
			signal1:        `[]`,
			signal2:        `[3]`,
			expectedResult: true,
		},
		{
			signal1:        `[[[]]]`,
			signal2:        `[[]]`,
			expectedResult: false,
		},
		{
			signal1:        `[1,[2,[3,[4,[5,6,7]]]],8,9]`,
			signal2:        `[1,[2,[3,[4,[5,6,0]]]],8,9]`,
			expectedResult: false,
		},
	}
	for _, testCase := range testCases {
		result := IsRightOrder(testCase.signal1, testCase.signal2)
		if result != testCase.expectedResult {
			t.Errorf("\nsignal1:%v\nsignal2:%v\nexpected: %v, got: %v\n", testCase.signal1, testCase.signal2, testCase.expectedResult, result)
		}
	}
}
