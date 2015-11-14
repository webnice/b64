package b64

func init() {
	var i int
	mapBase = make(map[int]byte)
	for i = 0; i < len(encodeBase); i++ {
		mapBase[i] = encodeBase[i]
	}
}

func newUint64(num uint64) (b64m *b64Marshal) {
	b64m = new(b64Marshal)
	b64m.SourceUint64 = num
	return
}

func newInt64(num int64) (b64m *b64Marshal) {
	b64m = new(b64Marshal)
	b64m.SourceUint64 = uint64(num)
	return
}

func (b64m *b64Marshal) Uint64() uint64 {
	return b64m.SourceUint64
}

func (b64m *b64Marshal) Int64() int64 {
	return int64(b64m.SourceUint64)
}

func (b64m *b64Marshal) Error() error {
	return nil
}

func (b64m *b64Marshal) convertFromUint64() *b64Marshal {
	var tmp = b64m.SourceUint64
	for i := 0; i <= 10; i++ {
		if tmp != 0 {
			b64m.Bytes[i] = mapBase[int(mask6bit&tmp)]
		} else {
			b64m.Bytes[i] = mapBase[0]
		}
		tmp = tmp >> 6
	}
	return b64m
}

func (b64m *b64Marshal) String() (ret string) {
	var i, n int
	var skeep = true
	var tmp [11]byte
	b64m.convertFromUint64()
	for i = 10; i >= 0; i-- {
		if b64m.Bytes[i] != mapBase[0] && skeep == true {
			skeep = false
		}
		if skeep && i != 0 {
			continue
		} else {
			tmp[n] = b64m.Bytes[i]
			n++
		}
	}
	ret = string(tmp[:n])
	return
}

func (b64m *b64Marshal) Bits() string {
	return toBits(b64m.SourceUint64)
}
