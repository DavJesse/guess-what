package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"student/files"
	"student/maths"
)

func main() {
	var c int
	args := os.Args[1:]

	// Handle insufficient or incorrect number of arguments
	if len(args) != 1 {
		log.Fatal("Please provide a single file path to a valid text file as an argument.")
	}

	// Define file path from arguments
	// Check file validity from its extenstion
	filePath := args[0]

	if !(strings.HasSuffix(filePath, ".txt")) {
		log.Fatalf("Invalid file format: %v is not a text file", filePath)
	}

	data := files.ReadFile(filePath)

	// Filter out empty files parsed as data
	if len(data) == 0 {
		log.Fatal("File parsed as data is empty.")
	}

	// Establish input and output arrays as parameters
	input, output := files.ExtractParams(data)

	// Calculate slope, y-intercept, and Pearson correlation coefficient to express linear regression and Pearson correlation coefficient
	slope := maths.CalculateSlope(input, output)
	// yIntercept := maths.CalculateYIntercept(input, output)
	// PearsonCoefficient := maths.PearsonCoefficient(input, output)

	y := slope*float64(len(input)) + float64(c)

	upperLimit := int(y) + 83
	lowerLimit := int(y) - 83

	fmt.Printf("%d - %d\n", upperLimit, lowerLimit)
}
