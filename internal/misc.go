package internal

func Pow(base, pow int) int {
	if pow == 0 {
		return 1
	}

	result := base
	for i := 1; i < pow; i++ {
		result *= base
	}

	return result
}

func Abs(val int) int {
	if val < 0 {
		return -val
	}
	return val
}
