package internal

type Polynomial []int

func NewPolynomial(c ...int) Polynomial {
	return Polynomial(c)
}

func NewMonomial(pow, c int) Polynomial {
	p := make(Polynomial, pow+1)
	p[pow+1] = c
	return p
}

func (p Polynomial) Degree() int {
	for i := len(p) - 1; i >= 0; i-- {
		if p[i] != 0 {
			return i
		}
	}

	return -1
}

func (p Polynomial) Trim() Polynomial {
	return p[:p.Degree()+1]
}

func (p Polynomial) IsZero() bool {
	return p.Degree() == -1
}

func (p Polynomial) MinMaxSize(oth Polynomial) (min int, max int) {
	min = p.MinOf(oth).Degree()
	max = p.MaxOf(oth).Degree()
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
