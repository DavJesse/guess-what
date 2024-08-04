package main

func atoi(s string) (int, string) {
	var result int
	var intSlc []int
	var err string

	for _, v := range s {
		if v >= '0' && v <= '9' {
			intSlc = append(intSlc, int(v-'0'))
		} else {
			err = "Input includes invalid characters"
			break
		}
	}

	for i := 0; i < len(intSlc); i++ {
		result = (result * 10) + intSlc[i]
	}

	return result, err
}
