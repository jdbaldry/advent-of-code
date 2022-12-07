// Various implementations for counting set bits in a uint64, also known as sideways addition.
// Taken from https://graphics.stanford.edu/~seander/bithacks.html#CountBitsSetNaive.
package main

//nolint:gochecknoglobals
var countBits = countBitsBrianKernighan

//nolint:varnamelen
func countBitsNaive(v uint64) uint64 {
	var c uint64

	for ; v != 0; v >>= 1 {
		c += v & 1
	}

	return c
}

//nolint:varnamelen
func countBitsBrianKernighan(v uint64) uint64 {
	var c uint64

	for ; v != 0; c++ {
		v &= v - 1
	}

	return c
}

//nolint:varnamelen
func countBitsParallel(v uint64) uint64 {
	var c uint64

	s := [...]uint64{1, 2, 4, 8, 16, 32}
	b := [...]uint64{
		0x5555555555555555,
		0x3333333333333333,
		0x0F0F0F0F0F0F0F0F,
		0x00FF00FF00FF00FF,
		0x0000FFFF0000FFFF,
		0x00000000FFFFFFFF,
	}

	c = v - ((v >> 1) & b[0])
	c = ((c >> s[1]) & b[1]) + (c & b[1])
	c = ((c >> s[2]) + c) & b[2]
	c = ((c >> s[3]) + c) & b[3]
	c = ((c >> s[4]) + c) & b[4]
	c = ((c >> s[5]) + c) & b[5]

	return c
}
