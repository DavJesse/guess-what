package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"linear-stats/maths"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var index, input []int
	var i int

	// Grab input and indices from standard input
	// Convert input to integer
	// Append input to data
	// Handle errors
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Println("This program only accepts numeric input\nPlease enter numeric values only.")
			continue
		}

		input = append(input, num)
		index = append(index, i)
		i++
		fmt.Println(index, "\n", input)

		// Calculate slope, y-intercept, and Pearson correlation coefficient to express linear regression and Pearson correlation coefficient
		slope := maths.CalculateSlope(index, input)
		yIntercept := maths.CalculateYIntercept(index, input)
		PearsonCoefficient := maths.PearsonCoefficient(index, input)

		fmt.Println(slope, "\n", yIntercept, "\n", PearsonCoefficient)
	}
}
