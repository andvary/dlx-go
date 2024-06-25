package dlx

type bitvector struct {
	v []uint64
}

func newBitvector(n int) *bitvector {
	if n <= 0 {
		n = 1
	}
	return &bitvector{v: make([]uint64, n>>6+1)}
}

func (b *bitvector) add(n int) {
	var pos, idx uint64
	pos = uint64(n) >> 6
	idx = 1 << (uint64(n) & 63)
	b.v[pos] |= idx
}

func (b *bitvector) isPresent(n int) bool {
	var pos, idx uint64
	pos = uint64(n) >> 6
	idx = 1 << (uint64(n) & 63)
	if b.v[pos]&idx == idx {
		return true
	}
	return false
}
