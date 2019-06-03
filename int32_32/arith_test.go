package int32_32

import (
	"testing"
)

func TestMul(t *testing.T) {
	tests := []struct {
		op1      Type
		op2      Type
		expected Type
	}{
		// Sign tests
		{+1 << 32, +1 << 32, +1 << 32},
		{-1 << 32, +1 << 32, -1 << 32},
		{+1 << 32, -1 << 32, -1 << 32},
		{-1 << 32, -1 << 32, +1 << 32},
		// Other
		{1<<32 + 1<<31, 1<<32 + 1<<31, 2<<32 + 1<<30}, // 1.5 * 1.5 = 2.25
		{1 << 33, 1 << 31, 1 << 32},                   // 2 * 0.5 = 1.0
		{1 << 34, 1 << 30, 1 << 32},                   // 4 * 0.25 = 1.0
		{1 << 35, 1 << 29, 1 << 32},                   // 8 * 0.125 = 1.0
	}
	for _, test := range tests {
		result := test.op1.Mul(test.op2)
		if result != test.expected {
			t.Errorf("Expected %v.Mul(%v) == %v, got: %v", test.op1, test.op2, test.expected, result)
		}
	}
}
