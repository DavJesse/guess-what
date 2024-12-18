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
	expectedLower, expectedUpper := 73, 261

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

func TestRemoveOutlier_SmallData(t *testing.T) {
	subject := []int{1, 2, 3, 4, 5}
	got := maths.RemoveOutlier(subject)
	expected := []int{1, 2, 3, 4, 5}

	var g, e int

	for g < len(got) && e < len(expected) {
		if got[g] != expected[e] {
			t.Errorf("Got: %d", got)
			t.Errorf("Expected: %d", expected)
			t.Error("TestRemoveOutlier Failed!")
		}
		g++
		e++
	}
}

func TestRemoveOutlier_BigDataNoOutlier(t *testing.T) {
	subject := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	got := maths.RemoveOutlier(subject)
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}

	var g, e int

	if len(got) != len(expected) {
		t.Error("TestRemoveOutlier Failed!")
		t.Errorf("Expected slice legth of: %d, got: %d\n", len(expected), len(got))
		return
	}
	for g < len(got) && e < len(expected) {
		if got[g] != expected[e] {
			t.Errorf("Got: %d", got)
			t.Errorf("Expected: %d", expected)
			t.Error("TestRemoveOutlier Failed!")
		}
		g++
		e++
	}
}

func TestIsOutlier(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	nums := []int{7, 170}
	var got []bool

	for _, num := range nums {
		if maths.IsOutlier(num, data) {
			got = append(got, true)
		} else {
			got = append(got, false)
		}
	}

	expected := []bool{false, true}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Expected: %t, Got: %t", expected, got)
		t.Error("TestIsOutlier Failed!")
	}
}
