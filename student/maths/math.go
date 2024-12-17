package maths

import (
	"log"
	"math"
	"sort"
)

func Mean(nums []int) float64 {
	var result float64

	// Find cummulative sum of all numbers
	for _, num := range nums {
		result += float64(num)
	}

	return result / float64(len(nums))
}

func PearsonCoefficient(input, output []int) float64 {
	var numerator, totalSquaredInputDeviation, totalSquaredOutputDeviation float64
	var i int

	// input and output arrays must have the same size
	if len(input) != len(output) {
		log.Fatal("Input and output arrays must have the same length.")
	}

	meanInput := Mean(input)
	meanOutput := Mean(output)

	// Calculate cumulative sum of multiple deviations
	for i < len(input) {
		inputDeviation := float64(input[i]) - meanInput
		outputDeviation := float64(output[i]) - meanOutput

		// accumulate sums of covariance component(numerator), squared deviations
		numerator += inputDeviation * outputDeviation
		totalSquaredInputDeviation += inputDeviation * inputDeviation
		totalSquaredOutputDeviation += outputDeviation * outputDeviation

		i++
	}

	// Establish denominator of equation
	denominator := math.Sqrt(totalSquaredInputDeviation * totalSquaredOutputDeviation)
	if denominator == 0 {
		log.Fatal("Cannot calculate Pearson correlation coefficient, the standard deviations are zero.")
	}

	return numerator / denominator
}

func CalculateSlope(input, output []int) float64 {
	var numerator, denominator, totalInput, totalOutput, totalInputOutput, totalSquareInput float64
	var i int

	for i < len(input) {
		// Establish parameters
		totalInputOutput += float64(input[i]) * float64(output[i])
		totalInput += float64(input[i])
		totalOutput += float64(output[i])
		totalSquareInput += float64(input[i]) * float64(input[i])

		i++
	}

	// Establish denominator of equation
	numerator = (float64(len(input)) * totalInputOutput) - (totalInput * totalOutput)
	denominator = (float64(len(input)) * totalSquareInput) - (totalInput * totalInput)

	return numerator / denominator
}

func CalculateYIntercept(input, output []int) float64 {
	// c = yˉ ​− m⋅xˉ
	// Where:
	// 'yˉ' is mean of the output
	// 'm' is slope
	// 'xˉ' is mean of the input
	slope := CalculateSlope(input, output)
	meanInput := Mean(input)
	meanOutput := Mean(output)

	return meanOutput - (slope * meanInput)
}

// Calculates the median of a given data set
func Median(numSlc []int) float64 {
	var median float64
	var medSlc []int

	// Establish a reference index for the median value
	half := len(numSlc) / 2
	sort.Ints(numSlc)

	if len(numSlc) > 2 {
		// If even, the median value will be at index 'half'
		if len(numSlc)%2 != 0 {
			median = float64(numSlc[half])

		} else {
			// if even the median will be the average of the middle values at indices 'half-1' and 'half'
			medSlc = append(medSlc, numSlc[half-1], numSlc[half])
			median = Mean(medSlc)
		}

	} else {
		// If one or two items in the slice, the median is the average
		median = Mean(numSlc)
	}

	return median
}

// Identifies and expells outliers from the given data
func RemoveOutlier(data []int) []int {
	var result []int
	if len(data) < 3 {
		return data
	}
	copy := append([]int(nil), data...)

	midPoint := len(copy) / 2
	sort.Ints(copy)

	firstQuartile := Median(copy[:midPoint])
	thirdQuartile := Median(copy[midPoint:])
	IQR := thirdQuartile - firstQuartile

	// Establish uper and lower bounds for outliers
	lowLimit := int(firstQuartile - IQR*2.5)
	upLimit := int(thirdQuartile + IQR*2.5)

	// Only include non-outliers in the result
	for _, num := range data {
		if num >= lowLimit && num <= upLimit {
			result = append(result, num)
		}
	}

	return result
}
