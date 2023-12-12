package main

import (
	"fmt"

	"github.com/mocurin/pg-project/internal"
)

func main() {
	f := internal.Field(17)
	for i := 1; i < f.Base(); i++ {
		fmt.Println(i, f.MltInv(i))
	}
}
