package b64

func init() {
	var i int
	mapKey = make(map[byte]int)
	for i = 0; i < len(encodeBase); i++ {
		mapKey[encodeBase[i]] = i
	}
}

func newString(inp string) (ret *b64Unarshal) {
	ret = new(b64Unarshal)
	ret.SourceString = inp
	ret.convertFromString()
	return
}

func (self *b64Unarshal) convertFromString() *b64Unarshal {
	var i, n int
	var ok bool
	self.SourceUint64 = 0
	for i = 10; i >= 0; i-- {
		self.Bytes[i] = mapBase[0]
		if len(self.SourceString)-i-1 >= 0 {
			self.Bytes[i] = self.SourceString[len(self.SourceString)-i-1]
			if n, ok = mapKey[self.Bytes[i]]; !ok {
				self.LastError = ErrorUnknownSymbolInString
				self.SourceUint64 = 0
				break
			} else {
				self.SourceUint64 += uint64(n) << (6 * uint(i))
			}
		}
	}
	return self
}

func (self *b64Unarshal) String() string {
	return self.SourceString
}

func (self *b64Unarshal) Uint64() uint64 {
	return self.SourceUint64
}

func (self *b64Unarshal) Int64() int64 {
	return int64(self.SourceUint64)
}

func (self *b64Unarshal) Bits() string {
	return toBits(self.SourceUint64)
}

func (self *b64Unarshal) Error() error {
	return self.LastError
}
