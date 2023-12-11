package internal

import "fmt"

const (
	fieldsMismatchError = "fields do not match"
	zeroDivisionError   = "can not divide by zero"
)

type FieldPolynomial struct {
	f Field
	p Polynomial
}

func NewFieldPolynomial(base int, c ...int) (FieldPolynomial, error) {
	f, err := NewField(base)
	if err != nil {
		return FieldPolynomial{}, fmt.Errorf("failed to create field: %w", err)
	}

	fc := f.ApplySeq(c...)
	p := NewPolynomial(fc...)
	return FieldPolynomial{
		f: f, p: p,
	}, nil
}

func (fp FieldPolynomial) NewPolynomial(c ...int) FieldPolynomial {
	return FieldPolynomial{
		f: fp.f,
		p: NewPolynomial(c...),
	}
}

func (fp FieldPolynomial) NewEmptyPolynomial(pow int) FieldPolynomial {
	return FieldPolynomial{
		f: fp.f,
		p: make(Polynomial, pow),
	}
}

func (fp FieldPolynomial) NewMonomial(pow, c int) FieldPolynomial {
	return FieldPolynomial{
		f: fp.f,
		p: NewPolynomial(pow, c),
	}
}

func (fp FieldPolynomial) Mlt(oth FieldPolynomial) FieldPolynomial {
	if fp.f != oth.f {
		panic(fieldsMismatchError)
	}

	size := len(fp.p) + len(oth.p) - 1
	p := make(Polynomial, size)

	for powOwn, cOwn := range fp.p {
		for powOth, cOth := range oth.p {
			p[powOwn+powOth] = fp.f.Mlt(cOwn, cOth)
		}
	}

	return FieldPolynomial{
		f: fp.f,
		p: p,
	}
}

func (fp FieldPolynomial) DivMod(oth FieldPolynomial) (quo FieldPolynomial, rem FieldPolynomial) {
	if fp.f != oth.f {
		panic(fieldsMismatchError)
	}

	div := oth
	if div.p.IsZero() {
		panic(zeroDivisionError)
	}

	rem = fp
	quo = fp.NewEmptyPolynomial(rem.p.Degree())
	divDegree := div.p.Degree()
	divLeadCf := div.p[divDegree]
	for {
		remDegree := rem.p.Degree()
		if remDegree < divDegree {
			break
		}

		remLeadCf := rem.p[remDegree]
		leadCf := fp.f.Div(remLeadCf, divLeadCf)
		leadDegree := remDegree - divDegree
		quo.p[leadDegree] += leadCf
		rem = rem.Sub(fp.NewMonomial(leadDegree, leadCf).Mlt(div))
	}

	quo.p = quo.p.Trim()
	rem.p = rem.p.Trim()
	return quo, rem
}

func (fp FieldPolynomial) Div(oth FieldPolynomial) FieldPolynomial {
	div, _ := fp.DivMod(oth)
	return div
}

func (fp FieldPolynomial) Mod(oth FieldPolynomial) FieldPolynomial {
	_, mod := fp.DivMod(oth)
	return mod
}

func (fp FieldPolynomial) Add(oth FieldPolynomial) FieldPolynomial {
	if fp.f != oth.f {
		panic(fieldsMismatchError)
	}

	minSize, maxSize := fp.p.MinMaxSize(oth.p)

	p := make(Polynomial, 0, maxSize)
	for i := 0; i < minSize; i++ {
		c := fp.f.Add(fp.p[i], oth.p[i])
		p = append(p, c)
	}

	p = append(p, fp.p.MaxOf(oth.p)[minSize:]...)
	return FieldPolynomial{
		f: fp.f,
		p: p,
	}
}

func (fp FieldPolynomial) Sub(oth FieldPolynomial) FieldPolynomial {
	return fp.Add(oth.Inv())
}

func (fp FieldPolynomial) Inv() FieldPolynomial {
	p := make(Polynomial, 0, len(fp.p))
	for _, c := range fp.p {
		p = append(p, fp.f.Base()-c)
	}
	return FieldPolynomial{
		f: fp.f,
		p: p,
	}
}

func (fp FieldPolynomial) Compute(val int) int {
	return fp.f.Apply(fp.p.Compute(val))
}