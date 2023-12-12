package internal

func Pow(base, pow int) int {
	result := 1
	for pow > 0 {
		if pow&1 == 1 {
			result = result * base
		}
		base = base * base
		pow = pow >> 1
	}

	return result
}

func PowMod(base, pow, mod int) int {
	result := 1
	for pow > 0 {
		if pow&1 == 1 {
			result = result * base % mod
		}
		base = base * base % mod
		pow = pow >> 1
	}

	return result
}

func Abs(val int) int {
	if val < 0 {
		return -val
	}
	return val
}
