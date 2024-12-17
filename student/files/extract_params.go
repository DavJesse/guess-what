package files

import (
	"math"
	"student/maths"
)

func ExtractParams(data []int) ([]int, []int) {
	var input, output []int

	// Populate input and output arrays with indices and values from data slice
	for i, num := range data {
		input = append(input, i)
		output = append(output, num)
	}
	return input, output
}

func PrematureGuess(input []int) (int, int) {
	// Eliminate outliers from the window and calculate range
	processed := maths.RemoveOutlier(input)

	// Use average to establish upper and lower limits for a guess
	average := maths.Mean(processed)

	upperLimit := int(math.Round(average)) + 95
	lowerLimit := int(math.Round(average)) - 93

	return int(lowerLimit), int(upperLimit)
}
