package internal

import (
	"fmt"
	"testing"
)

func TestPolynomialString(t *testing.T) {
	for _, c := range []struct {
		Expect     string
		Polynomial Polynomial
	}{
		{
			Expect:     "0",
			Polynomial: NewPolynomial(),
		},
		{
			Expect:     "X",
			Polynomial: NewPolynomial(0, 1),
		},
		{
			Expect:     "-X",
			Polynomial: NewPolynomial(0, -1),
		},
		{
			Expect:     "3X^2",
			Polynomial: NewPolynomial(0, 0, 3),
		},
		{
			Expect:     "-3X^2",
			Polynomial: NewPolynomial(0, 0, -3),
		},
		{
			Expect:     "1",
			Polynomial: NewPolynomial(1, 0, 0, 0),
		},
		{
			Expect:     "-1",
			Polynomial: NewPolynomial(-1, 0, 0, 0),
		},
		{
			Expect:     "-3+2X+5X^2+6X^3-7X^4",
			Polynomial: NewPolynomial(-3, 2, 5, 6, -7),
		},
	} {
		c := c
		t.Run(fmt.Sprintf("%v polynom against %s", []int(c.Polynomial), c.Expect), func(t *testing.T) {
			if got := c.Polynomial.String(); got != c.Expect {
				t.Fatalf("expected polynom string repr to match, got %s", got)
			}
		})
	}
}

func TestPolynomialEq(t *testing.T) {
	for _, c := range []struct {
		LHS   Polynomial
		RHS   Polynomial
		Equal bool
	}{
		{
			LHS:   NewPolynomial(0, 0, 1),
			RHS:   NewPolynomial(0, 0, 1),
			Equal: true,
		},
		{
			LHS:   NewPolynomial(),
			RHS:   NewPolynomial(),
			Equal: true,
		},
		{
			LHS:   NewPolynomial(0, 0, 1, 0),
			RHS:   NewPolynomial(0, 0, 1),
			Equal: true,
		},
		{
			LHS:   NewPolynomial(0, 1, 0),
			RHS:   NewPolynomial(0, 0, 1),
			Equal: false,
		},
	} {
		t.Run(fmt.Sprintf("%s eq to %s: %t", c.LHS, c.RHS, c.Equal), func(t *testing.T) {
			if got := c.LHS.Eq(c.RHS); got != c.Equal {
				t.Fatalf("expected to match")
			}
		})
	}
}

func TestPolynomialDegree(t *testing.T) {
	for _, c := range []struct {
		Name       string
		Expect     int
		Polynomial Polynomial
	}{
		{
			Name:       "simple",
			Expect:     2,
			Polynomial: NewPolynomial(1, 1, 1),
		},
		{
			Name:       "trailing zero",
			Expect:     0,
			Polynomial: NewPolynomial(1, 0, 0, 0, 0, 0),
		},
		{
			Name:       "zero",
			Expect:     -1,
			Polynomial: NewPolynomial(0, 0, 0, 0, 0, 0, 0),
		},
	} {
		c := c
		t.Run(fmt.Sprintf("%s polynomial %s of degree %d", c.Name, c.Polynomial, c.Expect), func(t *testing.T) {
			if got := c.Polynomial.Degree(); got != c.Expect {
				t.Fatalf("degrees do not match, got %d", got)
			}
		})
	}
}

func TestPolynomialTrim(t *testing.T) {
	for _, c := range []struct {
		Expect     int
		Polynomial Polynomial
	}{
		{
			Expect:     4,
			Polynomial: NewPolynomial(0, 0, 0, 1, 0, 0, 0),
		},
		{
			Expect:     3,
			Polynomial: NewPolynomial(0, 0, 1),
		},
		{
			Expect:     0,
			Polynomial: NewPolynomial(0, 0, 0, 0),
		},
	} {
		c := c
		t.Run(fmt.Sprintf("%s polynomial of expected len %d", c.Polynomial, c.Expect), func(t *testing.T) {
			oldDeg := c.Polynomial.Degree()
			p := c.Polynomial.Trim()
			if newDeg := p.Degree(); newDeg != oldDeg {
				t.Errorf("expected degree to stay the same after trim, got %d", newDeg)
			}

			if got := len(p); got != c.Expect {
				t.Fatalf("actual legnths do not match, got %d", got)
			}
		})
	}
}
