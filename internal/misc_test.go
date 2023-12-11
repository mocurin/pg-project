package internal

import (
	"fmt"
	"math"
	"testing"
)

func TestIntegerPow(t *testing.T) {
	for _, c := range []struct {
		Base int
		Pow  int
	}{
		{
			Base: 1,
			Pow:  1984,
		},
		{
			Base: 0,
			Pow:  1337,
		},
		{
			Base: 265,
			Pow:  0,
		},
		{
			Base: 9000,
			Pow:  2,
		},
	} {
		c := c
		t.Run(fmt.Sprintf("%d to the power of %d", c.Base, c.Pow), func(t *testing.T) {
			exp := math.Pow(float64(c.Base), float64(c.Pow))
			if got := Pow(c.Base, c.Pow); float64(got) != exp {
				t.Fatalf("expected to match math.Pow result, got %d", got)
			}
		})
	}
}

func TestIntegerAbs(t *testing.T) {
	for _, val := range []int{0, -1, 1} {
		val := val
		t.Run(fmt.Sprintf("%d absolute", val), func(t *testing.T) {
			exp := math.Abs(float64(val))
			if got := Abs(val); float64(got) != exp {
				t.Fatalf("expected to match math.Abs result, got %d", got)
			}
		})
	}
}
