package tests

import (
	"math"
	"reflect"
	"testing"

	"student/files"
	"student/maths"
)

func TestMean(t *testing.T) {
	subject := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expected := 5.5
	got := maths.Mean(subject)

	if got != expected {
		t.Errorf("Expected: %f, Got: %f", expected, got)
		t.Errorf("TestMean_EmptyDataSet failed!")
	}
}

func TestPearsonCoefficient(t *testing.T) {
	input := []int{2, 4, 6, 8}
	output := []int{3, 5, 7, 9}
	got := maths.PearsonCoefficient(input, output)
	expected := float64(1)

	tolerance := 1e-9
	difference := math.Abs(got - expected)

	if difference > tolerance {
		t.Errorf("Expected: %f, Got: %f", expected, got)
		t.Error("TestVariance Failed")
	}
}

func TestExtractParams(t *testing.T) {
	subject := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expected1, expected2 := []int{0, 2, 3, 4, 5, 6, 7, 8, 9}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	got1, got2 := files.ExtractParams(subject)

	if reflect.DeepEqual(got1, expected1) && reflect.DeepEqual(got2, expected2) {
		t.Errorf("Expected: %v, %v, Got: %v, %v", expected1, expected2, got1, got2)
		t.Error("TestExtractParams Failed")
	}
}

func TestCalculateSlope(t *testing.T) {
	input := []int{1, 2, 3, 4}
	output := []int{2, 3, 5, 7}
	got := maths.CalculateSlope(input, output)
	expected := 1.7

	tolerance := 1e-9
	difference := math.Abs(got - expected)

	if difference > tolerance {
		t.Errorf("Expected: %f, Got: %f", expected, got)
		t.Error("TestVariance Failed")
	}
}

func TestCalculateYIntercept(t *testing.T) {
	input := []int{1, 2, 3, 4}
	output := []int{2, 3, 5, 7}
	got := maths.CalculateYIntercept(input, output)
	expected := float64(0)

	tolerance := 1e-9
	difference := math.Abs(got - expected)

	if difference > tolerance {
		t.Errorf("Expected: %f, Got: %f", expected, got)
		t.Error("TestVariance Failed")
	}
}

func TestPrematureGuess(t *testing.T) {
	input := []int{136, 190, 173}
	gotLower, gotUpper := files.PrematureGuess(input)
	expectedLower, expectedUpper := 83, 249

	if gotLower != expectedLower || gotUpper != expectedUpper {
		t.Errorf("Expected: (%d, %d), Got: (%d, %d)", expectedLower, expectedUpper, gotLower, gotUpper)
		t.Error("TestPrematureGuess Failed")
	}
}

func TestMedian(t *testing.T) {
	intSlc := []int{1, 2, 3}
	got := maths.Median(intSlc)
	expected := float64(2)

	// Compare 'got' and 'expected'
	if got != expected {
		t.Errorf("Got: %f", got)
		t.Errorf("Expected: %f", expected)
		t.Error("TestMedian Failed")
	}
}
