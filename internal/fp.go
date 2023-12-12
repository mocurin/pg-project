package internal

import "fmt"

const (
	fieldsMismatchError  = "fields do not match"
	zeroDivisionError    = "can not divide by zero"
	normalizeFailedError = "normalization ended up with remainder"
)

type FieldPolynomial struct {
	f Field
	p Polynomial
}

func NewFieldPolynomial(base int, c []int) (FieldPolynomial, error) {
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

func NewCorrectFieldPolynomial(base int, c []int) FieldPolynomial {
	f, err := NewFieldPolynomial(base, c)
	if err != nil {
		panic(err)
	}
	return f
}

func (fp FieldPolynomial) Mlt(oth FieldPolynomial) FieldPolynomial {
	if fp.f != oth.f {
		panic(fieldsMismatchError)
	}

	size := fp.p.Degree() + oth.p.Degree() + 1
	p := make(Polynomial, size)

	for powOwn, cOwn := range fp.p {
		for powOth, cOth := range oth.p {
			powRes := powOwn + powOth
			p[powRes] = fp.f.Add(p[powRes], fp.f.Mlt(cOwn, cOth))
		}
	}

	return FieldPolynomial{
		f: fp.f,
		p: p.Trim(),
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
	quo = fp.f.NewEmptyPolynomial(rem.p.Degree())
	divDegree, divLeadCf := div.p.Leading()
	for {
		remDegree, remLeadCf := rem.p.Leading()
		if remDegree < divDegree {
			break
		}

		leadCf := fp.f.Div(remLeadCf, divLeadCf)
		leadDegree := remDegree - divDegree
		quo.p[leadDegree] = fp.f.Add(quo.p[leadDegree], leadCf)
		rem = rem.Sub(fp.f.NewMonomial(leadDegree, leadCf).Mlt(div))
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
		p: p.Trim(),
	}
}

func (fp FieldPolynomial) Sub(oth FieldPolynomial) FieldPolynomial {
	return fp.Add(oth.Inv())
}

func (fp FieldPolynomial) Inv() FieldPolynomial {
	p := make(Polynomial, 0, len(fp.p))
	for _, c := range fp.p {
		p = append(p, fp.f.AddInv(c))
	}
	return FieldPolynomial{
		f: fp.f,
		p: p,
	}
}

func (fp FieldPolynomial) Compute(val int) int {
	return fp.f.Apply(fp.p.Compute(val))
}

// func (fp FieldPolynomial) Substitute(oth FieldPolynomial) FieldPolynomial {
// 	if fp.f.Base() != oth.f.Base() {
// 		panic(fieldsMismatchError)
// 	}

// 	res := fp.f.NewEmptyPolynomial(fp.p.Degree() + oth.p.Degree() + 1)

// }

func (fp FieldPolynomial) GCD(oth FieldPolynomial) FieldPolynomial {
	divident := fp
	divisor := oth
	dividentDegree, dividentLeadCf := divident.p.Leading()
	divident, rem := divident.DivMod(fp.f.NewPolynomial(dividentLeadCf))
	if !rem.p.IsZero() {
		panic(normalizeFailedError)
	}

	divisorDegree, divisorLeadCf := divisor.p.Leading()
	divisor, rem = divisor.DivMod(fp.f.NewPolynomial(divisorLeadCf))
	if !rem.p.IsZero() {
		panic(normalizeFailedError)
	}

	if divisorDegree > dividentDegree {
		divident, divisor = divisor, divident
	}

	for {
		q, r := divident.DivMod(divisor)
		fmt.Printf("for %s mod %s rem is %s and quo is %s\n", divident.p, divisor.p, r.p, q.p)
		if r.p.IsZero() {
			return divisor
		}

		divident = divisor
		divisor = r
	}
}
