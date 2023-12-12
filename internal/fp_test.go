package internal

import (
	"fmt"
	"testing"
)

func TestFPAdd(t *testing.T) {
	for _, c := range []struct {
		LHS    FieldPolynomial
		RHS    FieldPolynomial
		Expect FieldPolynomial
	}{
		{
			LHS:    NewCorrectFieldPolynomial(3, []int{1, 1, 1}),
			RHS:    NewCorrectFieldPolynomial(3, []int{1, 1, 1}),
			Expect: NewCorrectFieldPolynomial(3, []int{2, 2, 2}),
		},
		{
			LHS:    NewCorrectFieldPolynomial(3, []int{1, 2, 1}),
			RHS:    NewCorrectFieldPolynomial(3, []int{2, 1, 2}),
			Expect: NewCorrectFieldPolynomial(3, []int{}),
		},
		{
			LHS:    NewCorrectFieldPolynomial(3, []int{1, 2, 1, 1, 2}),
			RHS:    NewCorrectFieldPolynomial(3, []int{2, 1}),
			Expect: NewCorrectFieldPolynomial(3, []int{0, 0, 1, 1, 2}),
		},
		{
			LHS:    NewCorrectFieldPolynomial(3, []int{}),
			RHS:    NewCorrectFieldPolynomial(3, []int{}),
			Expect: NewCorrectFieldPolynomial(3, []int{}),
		},
		{
			LHS:    NewCorrectFieldPolynomial(3, []int{1, 2, 1, 0}),
			RHS:    NewCorrectFieldPolynomial(3, []int{}),
			Expect: NewCorrectFieldPolynomial(3, []int{1, 2, 1}),
		},
	} {
		t.Run(fmt.Sprintf("addition of %s to %s equal to %s", c.LHS.p, c.RHS.p, c.Expect.p), func(t *testing.T) {
			if got := c.LHS.Add(c.RHS); !got.p.Eq(c.Expect.p) {
				t.Fatalf("expected result to match, got %s", got.p)
			}
		})
	}
}

func TestFPSub(t *testing.T) {
	for _, c := range []struct {
		LHS    FieldPolynomial
		RHS    FieldPolynomial
		Expect FieldPolynomial
	}{
		{
			LHS:    NewCorrectFieldPolynomial(3, []int{1, 1, 1}),
			RHS:    NewCorrectFieldPolynomial(3, []int{1, 1, 1}),
			Expect: NewCorrectFieldPolynomial(3, []int{}),
		},
		{
			LHS:    NewCorrectFieldPolynomial(3, []int{1, 2, 1}),
			RHS:    NewCorrectFieldPolynomial(3, []int{2, 1, 2}),
			Expect: NewCorrectFieldPolynomial(3, []int{2, 1, 2}),
		},
		{
			LHS:    NewCorrectFieldPolynomial(3, []int{1, 2, 1, 1, 2}),
			RHS:    NewCorrectFieldPolynomial(3, []int{1, 2}),
			Expect: NewCorrectFieldPolynomial(3, []int{0, 0, 1, 1, 2}),
		},
		{
			LHS:    NewCorrectFieldPolynomial(3, []int{}),
			RHS:    NewCorrectFieldPolynomial(3, []int{}),
			Expect: NewCorrectFieldPolynomial(3, []int{}),
		},
		{
			LHS:    NewCorrectFieldPolynomial(3, []int{1, 2, 1, 0}),
			RHS:    NewCorrectFieldPolynomial(3, []int{}),
			Expect: NewCorrectFieldPolynomial(3, []int{1, 2, 1}),
		},
	} {
		t.Run(fmt.Sprintf("subtraction of %s from %s equal to %s", c.RHS.p, c.LHS.p, c.Expect.p), func(t *testing.T) {
			if got := c.LHS.Sub(c.RHS); !got.p.Eq(c.Expect.p) {
				t.Fatalf("expected result to match, got %s", got.p)
			}
		})
	}
}

func TestFPInv(t *testing.T) {
	for _, c := range []struct {
		Polynomial FieldPolynomial
		Expect     FieldPolynomial
	}{
		{
			Polynomial: NewCorrectFieldPolynomial(3, []int{1, 1, 1}),
			Expect:     NewCorrectFieldPolynomial(3, []int{2, 2, 2}),
		},
		{
			Polynomial: NewCorrectFieldPolynomial(3, []int{1, 2, 1}),
			Expect:     NewCorrectFieldPolynomial(3, []int{2, 1, 2}),
		},
		{
			Polynomial: NewCorrectFieldPolynomial(3, []int{0, 0, 0}),
			Expect:     NewCorrectFieldPolynomial(3, []int{0, 0, 0}),
		},
		{
			Polynomial: NewCorrectFieldPolynomial(3, []int{}),
			Expect:     NewCorrectFieldPolynomial(3, []int{}),
		},
	} {
		t.Run(fmt.Sprintf("inversion of %s equal to %s", c.Polynomial.p, c.Expect.p), func(t *testing.T) {
			if got := c.Polynomial.Inv(); !got.p.Eq(c.Expect.p) {
				t.Fatalf("expected result to match, got %s", got.p)
			}
		})
	}
}

func TestFPMlt(t *testing.T) {
	for _, c := range []struct {
		LHS    FieldPolynomial
		RHS    FieldPolynomial
		Expect FieldPolynomial
	}{
		{
			LHS:    NewCorrectFieldPolynomial(3, []int{1, 1, 1}),
			RHS:    NewCorrectFieldPolynomial(3, []int{1, 1, 1}),
			Expect: NewCorrectFieldPolynomial(3, []int{1, 2, 0, 2, 1}),
		},
		{
			LHS:    NewCorrectFieldPolynomial(3, []int{1, 1, 1}),
			RHS:    NewCorrectFieldPolynomial(3, []int{}),
			Expect: NewCorrectFieldPolynomial(3, []int{}),
		},
		{
			LHS:    NewCorrectFieldPolynomial(3, []int{1, 1, 1}),
			RHS:    NewCorrectFieldPolynomial(3, []int{1}),
			Expect: NewCorrectFieldPolynomial(3, []int{1, 1, 1}),
		},
		{
			LHS:    NewCorrectFieldPolynomial(3, []int{1, 1, 1}),
			RHS:    NewCorrectFieldPolynomial(3, []int{2}),
			Expect: NewCorrectFieldPolynomial(3, []int{2, 2, 2}),
		},
	} {
		t.Run(fmt.Sprintf("multiplication of %s to %s equal to %s", c.LHS.p, c.RHS.p, c.Expect.p), func(t *testing.T) {
			if got := c.LHS.Mlt(c.RHS); !got.p.Eq(c.Expect.p) {
				t.Fatalf("expected result to match, got %s", got.p)
			}
		})
	}
}

func TestFPDivMod(t *testing.T) {
	for _, c := range []struct {
		LHS       FieldPolynomial
		RHS       FieldPolynomial
		Quoitet   FieldPolynomial
		Remainder FieldPolynomial
	}{
		{
			LHS:       NewCorrectFieldPolynomial(3, []int{1, 1, 1}),
			RHS:       NewCorrectFieldPolynomial(3, []int{1, 1, 1}),
			Quoitet:   NewCorrectFieldPolynomial(3, []int{1}),
			Remainder: NewCorrectFieldPolynomial(3, []int{}),
		},
		{
			LHS:       NewCorrectFieldPolynomial(3, []int{1, 1, 1}),
			RHS:       NewCorrectFieldPolynomial(3, []int{1}),
			Quoitet:   NewCorrectFieldPolynomial(3, []int{1, 1, 1}),
			Remainder: NewCorrectFieldPolynomial(3, []int{}),
		},
		{
			LHS:       NewCorrectFieldPolynomial(3, []int{2, 2, 2}),
			RHS:       NewCorrectFieldPolynomial(3, []int{2}),
			Quoitet:   NewCorrectFieldPolynomial(3, []int{1, 1, 1}),
			Remainder: NewCorrectFieldPolynomial(3, []int{}),
		},
		{
			LHS:       NewCorrectFieldPolynomial(3, []int{2}),
			RHS:       NewCorrectFieldPolynomial(3, []int{2, 2, 2}),
			Quoitet:   NewCorrectFieldPolynomial(3, []int{}),
			Remainder: NewCorrectFieldPolynomial(3, []int{2}),
		},
		{
			// An non-trivial (so imho) example from here
			// https://engineering.purdue.edu/kak/compsec/NewLectures/Lecture6.pdf
			LHS:       NewCorrectFieldPolynomial(7, []int{6, 4, 5}),
			RHS:       NewCorrectFieldPolynomial(7, []int{1, 2}),
			Quoitet:   NewCorrectFieldPolynomial(7, []int{6, 6}),
			Remainder: NewCorrectFieldPolynomial(7, []int{}),
		},
		{
			// An non-trivial (so imho) example from here
			// https://math.stackexchange.com/questions/350809/polynomial-division-in-a-field
			// UPD: Funny enough, "accepted" answer is wrong
			LHS:       NewCorrectFieldPolynomial(7, []int{5, 0, 2, 1}),
			RHS:       NewCorrectFieldPolynomial(7, []int{3, 0, 1}),
			Quoitet:   NewCorrectFieldPolynomial(7, []int{2, 1}),
			Remainder: NewCorrectFieldPolynomial(7, []int{6, 4}),
		},
	} {
		t.Run(fmt.Sprintf("division of %s by %s equal to %s quoitet and %s remainder", c.LHS.p, c.RHS.p, c.Quoitet.p, c.Remainder.p), func(t *testing.T) {
			q, r := c.LHS.DivMod(c.RHS)
			if !q.p.Eq(c.Quoitet.p) {
				t.Errorf("expected quoitet to match, got %s", q.p)
			}

			if !r.p.Eq(c.Remainder.p) {
				t.Errorf("expected remainder to match, got %s", r.p)
			}
		})
	}
}

func TestFPGCD(t *testing.T) {
	for _, c := range []struct {
		LHS    FieldPolynomial
		RHS    FieldPolynomial
		Expect FieldPolynomial
	}{
		{
			LHS:    NewCorrectFieldPolynomial(3, []int{1, 1, 1}),
			RHS:    NewCorrectFieldPolynomial(3, []int{1, 1, 1}),
			Expect: NewCorrectFieldPolynomial(3, []int{1, 1, 1}),
		},
		{
			LHS:    NewCorrectFieldPolynomial(3, []int{1, 1, 1}),
			RHS:    NewCorrectFieldPolynomial(3, []int{1, 1, 1, 1}),
			Expect: NewCorrectFieldPolynomial(3, []int{1}),
		},
		{
			LHS:    NewCorrectFieldPolynomial(3, []int{2, 2, 2}),
			RHS:    NewCorrectFieldPolynomial(3, []int{2}),
			Expect: NewCorrectFieldPolynomial(3, []int{1}),
		},
		{
			LHS:    NewCorrectFieldPolynomial(3, []int{1, 1, 1}),
			RHS:    NewCorrectFieldPolynomial(3, []int{2, 2, 2}),
			Expect: NewCorrectFieldPolynomial(3, []int{1, 1, 1}),
		},
		{
			LHS:    NewCorrectFieldPolynomial(3, []int{2, 2, 2}),
			RHS:    NewCorrectFieldPolynomial(3, []int{1, 1, 1}),
			Expect: NewCorrectFieldPolynomial(3, []int{1, 1, 1}),
		},
		// {
		// 	LHS:    NewCorrectFieldPolynomial(17, []int{-2, -1, 4, -7, 3, -7, 1}),
		// 	RHS:    NewCorrectFieldPolynomial(17, []int{0, -1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}),
		// 	Expect: NewCorrectFieldPolynomial(17, []int{-2, 7, -5, 6, 1}),
		// },
	} {
		t.Run(fmt.Sprintf("GCD of %s and %s equal to %s", c.LHS.p, c.RHS.p, c.Expect.p), func(t *testing.T) {
			if got := c.LHS.GCD(c.RHS); !got.p.Eq(c.Expect.p) {
				t.Fatalf("expected result to match, got %s", got.p)
			}
		})
	}
}
