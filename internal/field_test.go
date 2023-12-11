package internal

import (
	"fmt"
	"testing"
)

func TestFieldAddInv(t *testing.T) {
	base := 7
	f, err := NewField(base)
	if err != nil {
		t.Fatalf("expected to successfully create field of base %d, got error: %s", base, err)
	}

	for num, inv := range map[int]int{
		0: 0,
		1: 6,
		2: 5,
		3: 4,
		4: 3,
		5: 2,
		6: 1,
	} {
		t.Run(fmt.Sprintf("%d with expected additive inverse %d", num, inv), func(t *testing.T) {
			if got := f.AddInv(num); got != inv {
				t.Fatalf("do not match, got %d", got)
			}
		})
	}
}

func TestFieldMltInv(t *testing.T) {
	base := 7
	f, err := NewField(base)
	if err != nil {
		t.Fatalf("expected to successfully create field of base %d, got error: %s", base, err)
	}

	for num, inv := range map[int]int{
		1: 1,
		2: 4,
		3: 5,
		4: 2,
		5: 3,
		6: 6,
	} {
		t.Run(fmt.Sprintf("%d with expected miltiplicative inverse %d", num, inv), func(t *testing.T) {
			if got := f.MltInv(num); got != inv {
				t.Fatalf("do not match, got %d", got)
			}
		})
	}
}
