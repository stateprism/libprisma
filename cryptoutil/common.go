package cryptoutil

type Blocker struct {
	blockSize int
	last      int
	buff      []byte
}

func NewBlocker(blockSize int, data []byte) *Blocker {
	return &Blocker{
		blockSize: blockSize,
		last:      0,
		buff:      data,
	}
}

func (b *Blocker) Next() (int, []byte) {
	if b.last*b.blockSize >= len(b.buff) {
		return 0, nil
	}
	first := b.last * b.blockSize
	last := (b.last + 1) * b.blockSize
	if last > len(b.buff) {
		last = len(b.buff)
	}
	b.last++
	return last - first, b.buff[first:last]
}

func SecureCompare(tag []byte, finish []byte) bool {
	if len(tag) != len(finish) {
		return false
	}
	var res byte
	for i := 0; i < len(tag); i++ {
		res ^= tag[i] ^ finish[i]
	}
	return res == 0
}
