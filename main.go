package main

import (
	"fmt"
	"sort"

	"github.com/mocurin/pg-project/internal"
)

func main() {
	// fp := internal.NewCorrectFieldPolynomial(17, []int{-2, -1, 4, -7, 3, -7, 1})

	// fmt.Println(internal.FactorizeSequential(fp), internal.FactorizeBrute(fp))

loop:
	for {
		fp := internal.Field(7).RandomPolynomial(6)

		fmt.Println(fp)
		r1 := internal.FactorizeSequential(fp)
		sort.Slice(r1, func(i, j int) bool {
			return r1[i] > r1[j]
		})
		r2 := internal.FactorizeBrute(fp)
		sort.Slice(r2, func(i, j int) bool {
			return r2[i] > r2[j]
		})
		fmt.Println(r1, r2)
		if len(r1) != len(r2) {
			break
		}
		for i := 0; i < len(r1); i++ {
			if r1[i] != r2[i] {
				break loop
			}
		}
	}
}
