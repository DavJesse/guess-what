package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"student/files"
	"student/maths"
)

type FirstValue struct {
	Value   int
	IsEmpty bool
}

func main() {
	var upperLimit, lowerLimit int
	var rawData, cleanData []int
	var sampleData []int                  // Used to purge outliers
	c := FirstValue{IsEmpty: true}        // To capture value of x when y is zero
	scanner := bufio.NewScanner(os.Stdin) // Initialize scanner

	// Capture user input until EOF
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid number.")
		}
		if c.IsEmpty {
			c.Value = num
			c.IsEmpty = false
		}

		rawData = append(rawData, num)

		windowSize := 10
		end := len(rawData)
		start := end - windowSize

		// Handle window limits if data is less than window size
		if start < 0 {
			start = 0
		}

		// Capture sample data from the last 10 inputs, remove outliers
		sampleData = rawData[start:end]
		sampleData = maths.RemoveOutlier(sampleData)

		if len(rawData) > len(sampleData) && !maths.IsOutlier(num, sampleData) {
			cleanData = append(cleanData, num)
		}

		// Establish input and output arrays as parameters
		input, output := files.ExtractParams(cleanData)

		if len(input) < 100 {
			lowerLimit, upperLimit = files.PrematureGuess(sampleData)
		} else {
			// Calculate slope, y-intercept, and Pearson correlation coefficient to express linear regression and Pearson correlation coefficient
			slope := maths.CalculateSlope(input, output)
			y := slope*float64(len(input)) + float64(c.Value)

			upperLimit = int(y) + 75
			lowerLimit = int(y) - 100
		}

		fmt.Printf("%d %d\n", lowerLimit, upperLimit)
	}
}
