package integers

import "testing"

func assertCorrectResult(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got: %v, want %v", got, want)
	}
}

// test cases

type testCase struct {
	operator string
	input_a  int
	input_b  int
	result   int
}

var testCases = []testCase{
	{operator: "+", input_a: 3, input_b: 4, result: 7},
	{operator: "-", input_a: 4, input_b: 2, result: 2},
	{operator: "*", input_a: 4, input_b: 2, result: 8},
}

func TestCalculate(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.operator, func(t *testing.T) {
			got := Calculate(tc.operator, tc.input_a, tc.input_b)
			assertCorrectResult(t, got, tc.result)
		})
	}
}
