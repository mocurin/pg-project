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
