package main

func atoi(s string) int {
	var result int
	var intSlc []int

	// Convert to integer character by character
	for _, v := range s {
		intSlc = append(intSlc, int(v-'0'))
	}

	for i := 0; i < len(intSlc); i++ {
		result = (result * 10) + intSlc[i]
	}

	return result
}

func isNumeric(s string) bool {
	status := true

	// End function when non-numeric character is found
	for _, v := range s {
		if !(v >= '0' && v <= '9') {
			status = false
			break
		}
	}
	return status
}
