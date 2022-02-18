package lfsr

type LFSR struct {
	n    uint32
	seed uint32
	taps []uint32
}

func NewLFSR(n uint32, seed uint32, taps []uint32) *LFSR {
	if n > 32 {
		panic("max bits in this LFSR can be 32")
	}
	return &LFSR{
		n:    n,
		seed: seed,
		taps: taps,
	}
}

func (l *LFSR) NextBit() uint32 {
	seed := l.seed
	outputBit := l.seed & 1
	for _, tap := range l.taps {
		seed = seed ^ (l.seed >> tap)
	}
	msb := seed & 1
	l.seed = (l.seed >> 1) | (msb << (l.n - 1))

	return outputBit
}

func (l *LFSR) NextNumber(k int) uint32 {
	if k > 32 {
		panic("integer width can be at max 32")
	}
	var n uint32 = 0
	for i := 0; i < k; i++ {
		n = n << 1
		n |= l.NextBit()
	}
	return n
}
