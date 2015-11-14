package b64

func init() {
	var i int
	mapKey = make(map[byte]int)
	for i = 0; i < len(encodeBase); i++ {
		mapKey[encodeBase[i]] = i
	}
}

func newString(inp string) (b64m *b64Unarshal) {
	b64m = new(b64Unarshal)
	b64m.SourceString = inp
	b64m.convertFromString()
	return
}

func (b64m *b64Unarshal) convertFromString() *b64Unarshal {
	var i, n int
	var ok bool
	b64m.SourceUint64 = 0
	for i = 10; i >= 0; i-- {
		b64m.Bytes[i] = mapBase[0]
		if len(b64m.SourceString)-i-1 >= 0 {
			b64m.Bytes[i] = b64m.SourceString[len(b64m.SourceString)-i-1]
			if n, ok = mapKey[b64m.Bytes[i]]; !ok {
				b64m.LastError = ErrorUnknownSymbolInString
				b64m.SourceUint64 = 0
				break
			} else {
				b64m.SourceUint64 += uint64(n) << (6 * uint(i))
			}
		}
	}
	return b64m
}

func (b64m *b64Unarshal) String() string {
	return b64m.SourceString
}

func (b64m *b64Unarshal) Uint64() uint64 {
	return b64m.SourceUint64
}

func (b64m *b64Unarshal) Int64() int64 {
	return int64(b64m.SourceUint64)
}

func (b64m *b64Unarshal) Bits() string {
	return toBits(b64m.SourceUint64)
}

func (b64m *b64Unarshal) Error() error {
	return b64m.LastError
}
