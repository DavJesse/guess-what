package main

import (
	"bufio"
	"fmt"
	"log"
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
	var data []int
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

		data = append(data, num)
	}

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

	y := slope*float64(len(input)) + float64(c.Value)

	upperLimit := int(y) + 83
	lowerLimit := int(y) - 83

	fmt.Printf("%d - %d\n", upperLimit, lowerLimit)
}
