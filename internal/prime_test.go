package internal

import (
	"fmt"
	"testing"
)

func TestIsPrime(t *testing.T) {
	for num, isPrime := range map[int]bool{
		1:  true,
		2:  true,
		3:  true,
		4:  false,
		5:  true,
		6:  false,
		7:  true,
		8:  false,
		9:  false,
		10: false,
		11: true,
	} {
		num, isPrime := num, isPrime
		t.Run(fmt.Sprintf("is %d prime", num), func(t *testing.T) {
			if IsPrime(num) != isPrime {
				t.Error("is not prime")
			}
		})
	}
}
