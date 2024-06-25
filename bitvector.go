package dlx

type bitvector struct {
	v   []uint64
	idx uint64
}

func newBitvector(n int) *bitvector {
	if n <= 0 {
		n = 1
	}
	return &bitvector{v: make([]uint64, n>>6+1)}
}

func (b *bitvector) add(n int) {
	b.idx = 1 << (uint64(n) & 63)
	b.v[uint64(n)>>6] |= b.idx
}

func (b *bitvector) isPresent(n int) bool {
	b.idx = 1 << (uint64(n) & 63)
	if b.v[uint64(n)>>6]&b.idx == b.idx {
		return true
	}
	return false
}
