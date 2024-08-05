package main

func atoi(s string) (int, string) {
	var result int
	var intSlc []int
	var err string

	for _, v := range s {
		intSlc = append(intSlc, int(v-'0'))
	}

	for i := 0; i < len(intSlc); i++ {
		result = (result * 10) + intSlc[i]
	}

	return result, err
}

func isNumeric(s string) bool {
	status := true

	for _, v := range s {
		if !(v >= '0' && v <= '9') {
			status = false
			break
		}
	}
	return status
}
