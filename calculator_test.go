package calculator_test

import (
	"calculator"
	"math"
	"testing"
)

type Testcase struct {
	a, b float64
	want float64
}

func closeEnough(a, b, tolerance float64) bool {
	return math.Abs(a-b) <= tolerance
}

func TestAdd(t *testing.T) {
	t.Parallel()

	testCases := []Testcase{
		{a: 2, b: 2, want: 4},
		{a: 5, b: 8, want: 13},
		{a: 15, b: 4, want: 19},
	}

	for _, test := range testCases {
		got := calculator.Add(test.a, test.b)
		if test.want != got {
			t.Errorf("add(%f, %f): want %f, got %f", test.a, test.b, test.want, got)
		}
	}

}

func TestSubtract(t *testing.T) {
	t.Parallel()
	testCases := []Testcase{
		{a: 2, b: 2, want: 0},
		{a: 5, b: 8, want: -3},
		{a: 15, b: 4, want: 11},
	}

	for _, test := range testCases {
		got := calculator.Subtract(test.a, test.b)
		if test.want != got {
			t.Errorf("Subtract(%f, %f): want %f, got %f", test.a, test.b, test.want, got)
		}
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()
	testCases := []Testcase{
		{a: 2, b: 2, want: 4},
		{a: 5, b: 8, want: 40},
		{a: 15, b: 4, want: 60},
	}

	for _, test := range testCases {
		got := calculator.Multiply(test.a, test.b)
		if test.want != got {
			t.Errorf("Multiply(%f, %f): want %f, got %f", test.a, test.b, test.want, got)
		}
	}
}
func TestDivide(t *testing.T) {
	t.Parallel()
	testCases := []Testcase{
		{a: 2, b: 2, want: 1},
		{a: 20, b: 5, want: 4},
		{a: 15, b: 3, want: 5},
		{a: 1, b: 3, want: 0.333333},
	}

	for _, test := range testCases {
		got, err := calculator.Divide(test.a, test.b)
		if err != nil {
			t.Fatalf("want no error for valid input, got %v", err)
		}
		if !closeEnough(test.want, got, 0.001) {
			t.Errorf("Divide(%f, %f): want %f, got %f", test.a, test.b, test.want, got)
		}
	}
}

func TestDivideInvalid(t *testing.T) {
	t.Parallel()
	_, err := calculator.Divide(1, 0)
	if err == nil {
		t.Error("want error for invalid input, got nil")
	}
}

func TestSqrt(t *testing.T) {
	t.Parallel()
	testCases := []Testcase{
		{a: 4, want: 2},
		{a: 25, want: 5},
		{a: 81, want: 9},
		{a: 2, want: 1.4142},
	}

	for _, test := range testCases {
		got, err := calculator.Sqrt(test.a)
		if err != nil {
			t.Fatalf("want no error for valid input, got %v", err)
		}
		if !closeEnough(test.want, got, 0.1) {
			t.Errorf("Sqrt(%f): want %f, got %f", test.a, test.want, got)
		}

	}
}

func TestSqrtInvalid(t *testing.T) {
	t.Parallel()
	_, err := calculator.Sqrt(-5)
	if err == nil {
		t.Error("Sqrt(-5): want error for invalid input, got nil")
	}
}
