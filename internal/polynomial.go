package internal

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	zeroPolynomialDegree = -1
)

type Polynomial []int

func NewPolynomial(c ...int) Polynomial {
	return Polynomial(c)
}

func NewMonomial(pow, c int) Polynomial {
	p := make(Polynomial, pow+1)
	p[pow] = c
	return p
}

func (p Polynomial) String() string {
	var builder strings.Builder
	first := true
	for pow, c := range p {
		if c == 0 {
			continue
		}

		if pow == 0 {
			builder.WriteString(fmt.Sprint(c))
		} else {
			if c > 0 && !first {
				builder.WriteString("+")
			}

			if Abs(c) != 1 || pow == 0 {
				builder.WriteString(strconv.Itoa(c))
			} else if c < 0 {
				builder.WriteString("-")
			}

			if pow == 1 {
				builder.WriteString("X")
			} else {
				builder.WriteString("X^")
				builder.WriteString(strconv.Itoa(pow))
			}
		}

		first = false
	}

	s := builder.String()
	if len(s) == 0 {
		return "0"
	}
	return s
}

func (p Polynomial) Eq(oth Polynomial) bool {
	if p.Degree() != oth.Degree() {
		return false
	}

	p = p.Trim()
	oth = oth.Trim()
	for i := 0; i < len(p); i++ {
		if p[i] != oth[i] {
			return false
		}
	}

	return true
}

func (p Polynomial) Degree() int {
	for i := len(p) - 1; i >= 0; i-- {
		if p[i] != 0 {
			return i
		}
	}

	return zeroPolynomialDegree
}

func (p Polynomial) Trim() Polynomial {
	return p[:p.Degree()+1]
}

func (p Polynomial) IsZero() bool {
	return p.Degree() == zeroPolynomialDegree
}

func (p Polynomial) MinMaxSize(oth Polynomial) (min int, max int) {
	min = p.MinOf(oth).Degree() + 1
	max = p.MaxOf(oth).Degree() + 1
	return
}

func (p Polynomial) MaxOf(oth Polynomial) Polynomial {
	if p.Degree() > oth.Degree() {
		return p
	}
	return oth
}

func (p Polynomial) MinOf(oth Polynomial) Polynomial {
	if p.Degree() < oth.Degree() {
		return p
	}
	return oth
}

func (p Polynomial) Compute(val int) int {
	acc, factor := 0, 1
	for _, c := range p {
		acc += c * factor
		factor *= val
	}
	return acc
}

func (p Polynomial) Leading() (degree int, coeff int) {
	degree = p.Degree()
	coeff = 0
	if degree != zeroPolynomialDegree {
		coeff = p[degree]
	}

	return degree, coeff
}

func (p Polynomial) LeadingCoeff() int {
	_, coeff := p.Leading()
	return coeff
}
