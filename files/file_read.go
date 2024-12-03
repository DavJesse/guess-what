package files

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func ReadFile(path string) []int {
	var result []int

	// Open file for reading
	// Hanndle errors during file opening here
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	// Read file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		dataPoint := scanner.Text()

		// Convert data read into integer
		// Handle any errors encountered
		num, err := strconv.Atoi(dataPoint)
		if err != nil {
			log.Fatalf("File may contain non-numeric data: %v", err)
		}

		result = append(result, num)
	}

	return result
}
