package internal

import "sync"

func GetSubPolynomial(f Field, lambda int) FieldPolynomial {
	return f.NewPolynomial(lambda, 1)
}

func GetPolynomialPair(f Field, pow int) (FieldPolynomial, FieldPolynomial) {
	pq := f.NewMonomial(pow, 1)
	c := f.NewPolynomial(1)
	return pq.Add(c), pq.Sub(c)
}

func factorizeSequentialBody(fp FieldPolynomial, lambda int) (res []int) {
	if fp.P.Degree() == 1 {
		res = []int{fp.F.Apply(-fp.P[0])}
		return
	}

	if fp.P.Degree() < 1 {
		res = []int{}
		return
	}

	q := (fp.F.Base() - 1) / 2
	res = []int{}
	for ; lambda < fp.F.Base(); lambda++ {
		fp = fp.Substitute(GetSubPolynomial(fp.F, -lambda))
		if fp.Mod(fp.F.NewMonomial(1, 1)).P.IsZero() {
			res = append(res, fp.F.Apply(-lambda))
			fp = fp.Div(fp.F.NewMonomial(1, 1)).Substitute(GetSubPolynomial(fp.F, lambda))
			continue
		}
		break
	}

	q1, q2 := GetPolynomialPair(fp.F, q)

	g1 := q1.Normalize().GCD(fp.Normalize()).Normalize()
	g1 = g1.Substitute(GetSubPolynomial(fp.F, lambda))
	res = append(res, factorizeParallelBody(g1, lambda+1)...)

	g2 := q2.Normalize().GCD(fp.Normalize()).Normalize()
	g2 = g2.Substitute(GetSubPolynomial(fp.F, lambda))
	res = append(res, factorizeSequentialBody(g2, lambda+1)...)

	return res
}

func FactorizeSequential(fp FieldPolynomial) []int {
	if fp.P.IsZero() {
		return []int{}
	}

	gx := fp.Normalize().GCD(fp.F.NewFullPolynomial()).Normalize()
	return factorizeSequentialBody(gx, 0)
}

func factorizeParallelBody(fp FieldPolynomial, lambda int) (res []int) {
	if fp.P.Degree() == 1 {
		res = []int{fp.F.Apply(-fp.P[0])}
		return
	}

	if fp.P.Degree() < 1 {
		res = []int{}
		return
	}

	q := (fp.F.Base() - 1) / 2
	res = []int{}
	for ; lambda < fp.F.Base(); lambda++ {
		fp = fp.Substitute(GetSubPolynomial(fp.F, -lambda))
		if fp.Mod(fp.F.NewMonomial(1, 1)).P.IsZero() {
			res = append(res, fp.F.Apply(-lambda))
			fp = fp.Div(fp.F.NewMonomial(1, 1)).Substitute(GetSubPolynomial(fp.F, lambda))
			continue
		}
		break
	}

	q1, q2 := GetPolynomialPair(fp.F, q)

	reqRes := make([][]int, 2)
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		g1 := q1.Normalize().GCD(fp.Normalize()).Normalize()
		g1 = g1.Substitute(GetSubPolynomial(fp.F, lambda))
		reqRes[0] = factorizeParallelBody(g1, lambda+1)
	}()

	g2 := q2.Normalize().GCD(fp.Normalize()).Normalize()
	g2 = g2.Substitute(GetSubPolynomial(fp.F, lambda))
	reqRes[1] = factorizeParallelBody(g2, lambda+1)

	wg.Wait()

	res = append(res, reqRes[0]...)
	res = append(res, reqRes[1]...)

	return res
}

func FactorizeParralel(fp FieldPolynomial) []int {
	if fp.P.IsZero() {
		return []int{}
	}

	gx := fp.Normalize().GCD(fp.F.NewFullPolynomial()).Normalize()
	return factorizeParallelBody(gx, 0)
}

func FactorizeBrute(fp FieldPolynomial) []int {
	if fp.P.IsZero() {
		return []int{}
	}

	res := []int{}
	for i := 0; i < fp.F.Base(); i++ {
		if fp.Compute(i) == 0 {
			res = append(res, i)
		}
	}

	return res
}
