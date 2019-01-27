// Package internal provides functionality that is common between fixed-point integer types.
package internal

// ProdAndCarry128 performs multiplication on two uint64's and returns the product and carried bits
func ProdAndCarry128(x, y uint64) (prod, crry uint64) {
	var i uint
	for i = 0; i < int64Size; i++ {
		if y&(1<<i) != 0 {
			crry, prod = lshiftAndAdd128(crry, prod, x, i)
		}
	}
	return prod, crry
}

// lshiftAndAdd128 returns xhi*2**64 + xlo + (ylo<<i) as the upper and lower 64 bits
func lshiftAndAdd128(xhi, xlo, ylo uint64, n uint) (zhi, zlo uint64) {
	zhi = xhi + (ylo >> (int64Size - n))
	zlo = xlo + (ylo << n)
	if zlo < xlo {
		zhi++
	}
	return zhi, zlo
}
