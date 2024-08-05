package main

import (
	"bufio"
	"fmt"
	"os"
)

func ReadInput() {
	var issue string
	var input string
	var data []int
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		input = scanner.Text()
		num, err := atoi(input)
		if err != "" {
			issue = err
			fmt.Println(issue)
			// return // []int{}, issue
		}
		data = append(data, num)
		fmt.Println(data)

	}
}
