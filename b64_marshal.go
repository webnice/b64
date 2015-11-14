package b64

func init() {
	var i int
	mapBase = make(map[int]byte)
	for i = 0; i < len(encodeBase); i++ {
		mapBase[i] = encodeBase[i]
	}
}

func newUint64(num uint64) (ret *b64Marshal) {
	ret = new(b64Marshal)
	ret.SourceUint64 = num
	return
}

func newInt64(num int64) (ret *b64Marshal) {
	ret = new(b64Marshal)
	ret.SourceUint64 = uint64(num)
	return
}

func (self *b64Marshal) Uint64() uint64 {
	return self.SourceUint64
}

func (self *b64Marshal) Int64() int64 {
	return int64(self.SourceUint64)
}

func (self *b64Marshal) Error() error {
	return nil
}

func (self *b64Marshal) convertFromUint64() *b64Marshal {
	var tmp uint64 = self.SourceUint64
	for i := 0; i <= 10; i++ {
		if tmp != 0 {
			self.Bytes[i] = mapBase[int(mask6bit&tmp)]
		} else {
			self.Bytes[i] = mapBase[0]
		}
		tmp = tmp >> 6
	}
	return self
}

func (self *b64Marshal) String() (ret string) {
	var i, n int
	var skeep bool = true
	var tmp [11]byte
	self.convertFromUint64()
	for i = 10; i >= 0; i-- {
		if self.Bytes[i] != mapBase[0] && skeep == true {
			skeep = false
		}
		if skeep && i != 0 {
			continue
		} else {
			tmp[n] = self.Bytes[i]
			n += 1
		}
	}
	ret = string(tmp[:n])
	return
}

func (self *b64Marshal) Bits() string {
	return toBits(self.SourceUint64)
}
