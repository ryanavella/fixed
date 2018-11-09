package internal

import (
	"testing"
)

func TestProdAndCarry128(t *testing.T) {
	tests := []struct {
		op1  uint64
		op2  uint64
		exp1 uint64
		exp2 uint64
	}{
		{3, 5, 15, 0},
		{5, 3, 15, 0},
		{1<<32 + 1<<31, 1<<32 + 1<<31, 1 << 62, 2},
		{1 << 63, 2, 0, 1},
		{2, 1 << 63, 0, 1},
		{0xFFFFFFFFFFFFFFFF, 0xFFFFFFFFFFFFFFFF, 1, 0xFFFFFFFFFFFFFFFE},
	}
	for _, test := range tests {
		result1, result2 := ProdAndCarry128(test.op1, test.op2)
		if result1 != test.exp1 || result2 != test.exp2 {
			t.Errorf("Expected ProdAndCarry128(%v, %v) == %v, %v; got: %v, %v", test.op1, test.op2, test.exp1, test.exp2, result1, result2)
		}
	}
}
