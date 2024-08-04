package main

import (
	"bufio"
	"os"
)

func scanInput() []int {
	var input string
	var data []int
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		input = scanner.Text()
	}
}
