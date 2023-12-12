package internal

import "fmt"

const (
	fieldsMismatchError  = "fields do not match"
	zeroDivisionError    = "can not divide by zero"
	normalizeFailedError = "normalization ended up with remainder"
)

type FieldPolynomial struct {
	F Field
	P Polynomial
}

func NewFieldPolynomial(base int, c []int) (FieldPolynomial, error) {
	f, err := NewField(base)
	if err != nil {
		return FieldPolynomial{}, fmt.Errorf("failed to create field: %w", err)
	}

	fc := f.ApplySeq(c...)
	p := NewPolynomial(fc...)
	return FieldPolynomial{
		F: f, P: p,
	}, nil
}

func NewCorrectFieldPolynomial(base int, c []int) FieldPolynomial {
	f, err := NewFieldPolynomial(base, c)
	if err != nil {
		panic(err)
	}
	return f
}

func (fp FieldPolynomial) Mlt(oth FieldPolynomial) FieldPolynomial {
	if fp.F != oth.F {
		panic(fieldsMismatchError)
	}

	fpSize := fp.P.Degree()
	if fpSize == zeroPolynomialDegree {
		fpSize += 1
	}

	othSize := oth.P.Degree()
	if othSize == zeroPolynomialDegree {
		othSize += 1
	}

	size := fpSize + othSize + 1
	p := make(Polynomial, size)

	for powOwn, cOwn := range fp.P {
		for powOth, cOth := range oth.P {
			powRes := powOwn + powOth
			p[powRes] = fp.F.Add(p[powRes], fp.F.Mlt(cOwn, cOth))
		}
	}

	return FieldPolynomial{
		F: fp.F,
		P: p.Trim(),
	}
}

func (fp FieldPolynomial) DivMod(oth FieldPolynomial) (quo FieldPolynomial, rem FieldPolynomial) {
	if fp.F != oth.F {
		panic(fieldsMismatchError)
	}

	div := oth
	if div.P.IsZero() {
		panic(zeroDivisionError)
	}

	rem = fp
	quo = fp.F.NewEmptyPolynomial(rem.P.Degree())
	divDegree, divLeadCf := div.P.Leading()
	for {
		remDegree, remLeadCf := rem.P.Leading()
		if remDegree < divDegree {
			break
		}

		leadCf := fp.F.Div(remLeadCf, divLeadCf)
		leadDegree := remDegree - divDegree
		quo.P[leadDegree] = fp.F.Add(quo.P[leadDegree], leadCf)
		rem = rem.Sub(fp.F.NewMonomial(leadDegree, leadCf).Mlt(div))
	}

	quo.P = quo.P.Trim()
	rem.P = rem.P.Trim()
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
	if fp.F != oth.F {
		panic(fieldsMismatchError)
	}

	minSize, maxSize := fp.P.MinMaxSize(oth.P)

	p := make(Polynomial, 0, maxSize)
	for i := 0; i < minSize; i++ {
		c := fp.F.Add(fp.P[i], oth.P[i])
		p = append(p, c)
	}

	p = append(p, fp.P.MaxOf(oth.P)[minSize:]...)
	return FieldPolynomial{
		F: fp.F,
		P: p.Trim(),
	}
}

func (fp FieldPolynomial) Sub(oth FieldPolynomial) FieldPolynomial {
	return fp.Add(oth.InvAdd())
}

func (fp FieldPolynomial) InvAdd() FieldPolynomial {
	p := make(Polynomial, 0, len(fp.P))
	for _, c := range fp.P {
		p = append(p, fp.F.AddInv(c))
	}

	return FieldPolynomial{
		F: fp.F,
		P: p,
	}
}

func (fp FieldPolynomial) Compute(val int) int {
	return fp.F.Apply(fp.P.Compute(val))
}

func (fp FieldPolynomial) Normalize() FieldPolynomial {
	oth := fp.F.NewPolynomial(fp.P.LeadingCoeff())
	return fp.Div(oth)
}

func (fp FieldPolynomial) Substitute(oth FieldPolynomial) FieldPolynomial {
	if fp.F.Base() != oth.F.Base() {
		panic(fieldsMismatchError)
	}

	if fp.P.IsZero() {
		return fp
	}

	res := fp.F.NewPolynomial(fp.P[0])
	deg := fp.P.Degree()
	acc := oth
	for pow := 1; pow <= deg; pow++ {
		coef := fp.P[pow]
		c := fp.F.NewPolynomial(coef)
		res = res.Add(acc.Mlt(c))
		acc = acc.Mlt(oth)
	}

	res.P = res.P.Trim()
	return res
}

func (fp FieldPolynomial) GCD(oth FieldPolynomial) FieldPolynomial {
	divident := fp
	divisor := oth
	if divisor.P.Degree() > divident.P.Degree() {
		divident, divisor = divisor, divident
	}

	for {
		r := divident.Mod(divisor)
		if r.P.IsZero() {
			return divisor
		}

		divident = divisor
		divisor = r
	}
}
