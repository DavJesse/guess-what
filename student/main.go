package main

import (
	"os"
	"bufio"
)

func main() {
	var input string
	var data []int
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		input = scanner.Text()
	}
}