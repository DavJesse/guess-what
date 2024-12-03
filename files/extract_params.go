package files

func ExtractParams(data []int) ([]int, []int) {
	var input, output []int

	// Populate input and output arrays with indices and values from data slice
	for i, num := range data {
		input = append(input, i)
		output = append(output, num)
	}
	return input, output
}
