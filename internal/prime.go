package internal

import "math"

var primes = map[int]bool{}

func IsPrime(victim int) bool {
	if prime, exists := primes[victim]; exists {
		return prime
	}

	boundary := math.Sqrt(float64(victim))
	for i := 2; float64(i) <= boundary; i++ {
		if !IsPrime(i) {
			continue
		}

		if victim%i == 0 {
			primes[victim] = false
			return false
		}
	}

	primes[victim] = true
	return true
}
