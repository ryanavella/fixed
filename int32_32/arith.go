package int32_32

import (
	"github.com/ryanavella/fixed/internal"
)

// Max returns the maximum of two int32_32.Type's
func Max(x, y Type) Type {
	if x > y {
		return x
	}
	return y
}

// MaxN returns the maximum of N int32_32.Type's
func MaxN(nums ...Type) Type {
	res := Type(minInt64)
	for _, n := range nums {
		if n > res {
			res = n
		}
	}
	return res
}

// Min returns the minimum of two int32_32.Type's
func Min(x, y Type) Type {
	if x < y {
		return x
	}
	return y
}

// MinN returns the minimum of N int32_32.Type's
func MinN(nums ...Type) Type {
	res := Type(maxInt64)
	for _, n := range nums {
		if n < res {
			res = n
		}
	}
	return res
}

// Mul returns the product of two int32_32.Type's
func (x Type) Mul(y Type) Type {
	sign := Type(+1)
	if x < 0 {
		x = -x
		sign = -sign
	}
	if y < 0 {
		y = -y
		sign = -sign
	}
	prod, carry := internal.ProdAndCarry128(uint64(x), uint64(y))
	return sign * Type(((prod+1<<(int32Size-1))>>int32Size)|(carry<<int32Size))
}
