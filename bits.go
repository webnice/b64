package b64

func toBits(n uint64) string {
	var bytes [64]byte
	for i := 63; i >= 0; i-- {
		if n&(1<<uint(i)) > 0 {
			bytes[63-i] = '1'
		} else {
			bytes[63-i] = '0'
		}
	}
	return string(bytes[:])
}

//func unsetBit(n uint64, pos uint) uint64 {
//	mask := ^(uint64(1) << pos)
//	n &= mask
//	return n
//}

//func setBit(n uint64, pos uint) uint64 {
//	n |= (1 << pos)
//	return n
//}

//func isBit(n uint64, pos uint) bool {
//	val := n & (1 << pos)
//	return (val > 0)
//}
