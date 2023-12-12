package internal

import (
	"fmt"
	"math/rand"
)

const zeroElementMultiplicativeInverseError = "can not find multiplicative inverse for a zero element"

type Field int

func NewField(base int) (Field, error) {
	if base <= 2 {
		return Field(0), fmt.Errorf("can not make field of base <= 2, got: %d", base)
	}

	if !IsPrime(base) {
		return Field(0), fmt.Errorf("field base should at least be prime, got %d", base)
	}

	return Field(base), nil
}

func (f Field) Base() int {
	return int(f)
}

func (f Field) AddInv(val int) int {
	return f.Apply(f.Base() - val)
}

func (f Field) MltInv(val int) int {
	if val == 0 {
		panic(zeroElementMultiplicativeInverseError)
	}

	// Since Field is formed from a prime number (base^1), we can find multiplicative inverse using val^(base^1-2)
	// according to wikipedia:
	// https://en.wikipedia.org/wiki/Finite_field_arithmetic#Multiplicative_inverse
	return f.Apply(Pow(val, f.Base()-2))
}

func (f Field) Div(lhs, rhs int) int {
	return f.Apply(lhs * f.MltInv(rhs))
}

func (f Field) Mlt(lhs, rhs int) int {
	return f.Apply(lhs * rhs)
}

func (f Field) Add(lhs, rhs int) int {
	return f.Apply(lhs + rhs)
}

func (f Field) Sub(lhs, rhs int) int {
	return f.Apply(lhs - rhs)
}

func (f Field) Apply(victim int) int {
	res := victim % f.Base()
	if res < 0 {
		return f.Base() + res
	}
	return res
}

func (f Field) ApplySeq(victims ...int) []int {
	res := make([]int, 0, len(victims))
	for _, val := range victims {
		res = append(res, f.Apply(val))
	}
	return res
}

func (f Field) NewPolynomial(c ...int) FieldPolynomial {
	return FieldPolynomial{
		F: f,
		P: NewPolynomial(c...),
	}
}

func (f Field) NewEmptyPolynomial(pow int) FieldPolynomial {
	return FieldPolynomial{
		F: f,
		P: make(Polynomial, pow+1),
	}
}

func (f Field) NewMonomial(pow, c int) FieldPolynomial {
	fp := f.NewEmptyPolynomial(pow)
	fp.P[pow] = c
	return fp
}

func (f Field) NewFullPolynomial() FieldPolynomial {
	fp := f.NewMonomial(f.Base(), 1)
	fp.P[1] = f.Apply(-1)
	return fp
}

func (f Field) RandomPolynomial(pow int) FieldPolynomial {
	fp := f.NewEmptyPolynomial(pow)
	for i := range fp.P {
		fp.P[i] = f.Apply(rand.Int())
	}

	return fp
}
