package int32_32

import (
	"fmt"
	"math/rand"
)

// Type is a signed 32.32 fixed-point number.
//
// All integers in the interval [-2147483648, +2147483647] can be represented.
// The fractional part has 32 bits of precision, with an approximate step size of 2.3283e-10.
//
// Example:
//   1.5 is represented as int32_32.Type(1<<32 + 1<<31)
type Type int64

// New returns a default-value int32_32.Type.
func New() Type {
	return Type(0)
}

// FromInt32 returns a int32_32.Type representation of an int32.
//
// This will never overflow, as int32 is a strict subset of int32_32.Type.
func FromInt32(x int32) Type {
	return (Type(x) << int32Size)
}

// FromFloat64 returns a int32_32.Type representation of a float64, rounding to the nearest fixed point
//
// This has the potential to overflow (or even panic) depending on the mantissa and exponent.
func FromFloat64(x float64) Type {
	return Type(x * (1 << int32Size))
}

// Rand returns a new pseudo-random int32_32.Type.
func Rand() Type {
	return Type(rand.Uint64())
}

// String returns a hexadecimal representation of a 32.32 fixed-point number.
//
// For example, the number one-and-a-half is represented as "0x1.80000000".
func (x Type) String() string {
	if x < 0 {
		x = -x
		return fmt.Sprintf("-%#x.%08x", int32(x>>int32Size), uint32(x))
	}
	return fmt.Sprintf("%#x.%08x", int32(x>>int32Size), uint32(x))
}
