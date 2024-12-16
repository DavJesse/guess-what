package tests

import (
	"os"
	"reflect"
	"testing"

	"linear-stats/files"
)

func TestReadFile_EmptyFile(t *testing.T) {
	// Create a temporary empty file
	tempFile, err := os.CreateTemp("", "empty.txt")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	defer os.Remove(tempFile.Name())
	tempFile.Close()

	// Call ReadFile with the empty file
	result := files.ReadFile(tempFile.Name())

	// Check if the result is an empty slice
	if len(result) != 0 {
		t.Errorf("Expected empty slice, but got slice with length %d", len(result))
	}
}

func TestReadFile_SingleDigitIntegers(t *testing.T) {
	// Create a temporary file with single-digit integers
	tempFile, err := os.CreateTemp("", "single_digits.txt")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	defer os.Remove(tempFile.Name())

	// Write single-digit integers to the file
	content := []byte("1\n2\n3\n4\n5\n")
	if _, err := tempFile.Write(content); err != nil {
		t.Fatalf("Failed to write to temporary file: %v", err)
	}
	tempFile.Close()

	// Call ReadFile with the temporary file
	result := files.ReadFile(tempFile.Name())

	// Check if the result matches the expected output
	expected := []int{1, 2, 3, 4, 5}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestReadFile_MultiDigitIntegers(t *testing.T) {
	// Create a temporary file with multi-digit integers
	tempFile, err := os.CreateTemp("", "multi-digit.txt")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	defer os.Remove(tempFile.Name())

	// Write multi-digit integers to the file
	content := []byte("123\n4567\n89012\n345678\n")
	if _, err := tempFile.Write(content); err != nil {
		t.Fatalf("Failed to write to temporary file: %v", err)
	}
	tempFile.Close()

	// Call ReadFile with the multi-digit integer file
	result := files.ReadFile(tempFile.Name())

	// Check if the result matches the expected values
	expected := []int{123, 4567, 89012, 345678}
	if len(result) != len(expected) {
		t.Errorf("Expected %d integers, but got %d", len(expected), len(result))
	}
	for i, v := range expected {
		if result[i] != v {
			t.Errorf("Expected %d at index %d, but got %d", v, i, result[i])
		}
	}
}

func TestReadFile_PositiveNegativeIntegers(t *testing.T) {
	// Create a temporary file with positive and negative integers
	tempFile, err := os.CreateTemp("", "integers.txt")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	defer os.Remove(tempFile.Name())

	// Write positive and negative integers to the file
	content := []byte("10\n-5\n3\n-8\n0\n")
	if _, err := tempFile.Write(content); err != nil {
		t.Fatalf("Failed to write to temporary file: %v", err)
	}
	tempFile.Close()

	// Call ReadFile with the temporary file
	result := files.ReadFile(tempFile.Name())

	// Check if the result matches the expected values
	expected := []int{10, -5, 3, -8, 0}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}
