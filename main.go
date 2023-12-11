package main

import (
	"fmt"

	"github.com/mocurin/pg-project/internal"
)

func main() {
	lhs := internal.NewCorrectFieldPolynomial(3, []int{1, 1, 1})
	rhs := internal.NewCorrectFieldPolynomial(3, []int{2, 2, 2})
	fmt.Println(lhs.GCD(rhs))
}
