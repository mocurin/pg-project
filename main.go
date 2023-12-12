package main

import (
	"fmt"

	"github.com/mocurin/pg-project/internal"
)

func main() {
	fp := internal.NewCorrectFieldPolynomial(17, []int{-2, -1, 4, -7, 3, -7, 1})

	fmt.Println(internal.FactorizeSequential(fp))
}
