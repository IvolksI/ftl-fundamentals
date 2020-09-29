package calculator_test

import (
	"calculator"
	"testing"
)

type testCase struct {
	a, b        float64
	want        float64
	name        string
	errExpected bool
}

func TestAdd(t *testing.T) {
	t.Parallel()
	testCases := []testCase{
		{a: 2, b: 2, want: 4, name: "Two numbers which sum to a positive", errExpected: false},
		{a: 1, b: 1, want: 2, name: "Two numbers which sum to a positive", errExpected: false},
		{a: 5, b: 0, want: 5, name: "Adding 0 to a number returns the number", errExpected: false},
	}

	for _, tc := range testCases {
		got := calculator.Add(tc.a, tc.b)
		if tc.want != got {
			t.Error(tc.name)
			t.Errorf("Add(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}
}

func TestSubtract(t *testing.T) {
	t.Parallel()
	testCases := []testCase{
		{a: 2, b: 2, want: 0, name: "Two numbers which multiply to a positive", errExpected: false},
		{a: 1, b: -1, want: 2, name: "One negative and a positive number at same value which results to a positive", errExpected: false},
		{a: -5, b: 2, want: -7, name: "Negative and a positive number which results to a negative number", errExpected: false},
	}
	for _, tc := range testCases {
		got := calculator.Subtract(tc.a, tc.b)
		if tc.want != got {
			t.Error(tc.name)
			t.Errorf("Subtract(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()
	testCases := []testCase{
		{a: 2, b: 2, want: 4, name: "Two numbers which multiply to a positive", errExpected: false},
		{a: 1, b: -1, want: -1, name: "Multiply with a negative number results in negative", errExpected: false},
		{a: -5, b: 2, want: -10, name: "Multiply with a negative number results in negative", errExpected: false},
		{a: -5, b: 0, want: 0, name: "Multiply with 0 results in 0", errExpected: false},
	}
	for _, tc := range testCases {
		got := calculator.Multiply(tc.a, tc.b)
		if tc.want != got {
			t.Error(tc.name)
			t.Errorf("Multiply(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}
}

func TestDivide(t *testing.T) {
	t.Parallel()
	testCases := []testCase{
		{a: 2, b: 2, want: 1, name: "Two numbers which result to a positive", errExpected: false},
		{a: 1, b: 0, want: 0, name: "Dividing by 0 returns error", errExpected: true},
	}

	for _, tc := range testCases {
		got, err := calculator.Divide(tc.a, tc.b)
		errReceived := err != nil
		if tc.errExpected != errReceived {
			t.Fatalf("Divide(%f, %f): unexpected error status: %v", tc.a, tc.b, errReceived)
		}
		if !tc.errExpected && tc.want != got {
			t.Error(tc.name)
			t.Errorf("Divide(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}
}

func TestSqrt(t *testing.T) {
	t.Parallel()
	var want float64 = 4
	got, err := calculator.Sqrt(16)
	if err != nil {
		t.Fatalf("Unexpected input, number cannot be negative")
	}
	if err == nil && want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}
