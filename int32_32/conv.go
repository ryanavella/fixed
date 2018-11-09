package int32_32

// Floor returns the greatest integer less than or equal to an int32_32.Type
func (x Type) Floor() int {
	return int(x >> int32Size)
}

// Round returns the nearest integer value to x. Ties are rounded up.
func (x Type) Round() int {
	return int(((x >> (int32Size - 1)) + 1) >> 1)
}

// Ceil returns the least integer greater than or equal to an int32_32.Type
func (x Type) Ceil() int {
	if x|0x0000FFFF == 0x00000000 {
		return int(x >> int32Size)
	}
	return int(x>>int32Size) + 1
}

// Float64 returns a float64 representation of an int32_32.Type
func (x Type) Float64() float64 {
	return float64(x) / (1 << int32Size)
}
