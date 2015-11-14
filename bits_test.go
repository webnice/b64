package b64

import (
	"testing"
)

func TestToBits(t *testing.T) {
	a := toBits(0xFFFFFFFFFFFFFFFF)
	if a != `1111111111111111111111111111111111111111111111111111111111111111` {
		t.Errorf("Error in toBits()")
	}
}

func TestUnsetBit(t *testing.T) {
	a := unsetBit(0xFFFFFFFFFFFFFFFF, 63)
	if a != 9223372036854775807 {
		t.Errorf("Error in unsetBit()")
	}
}

func TestSetBit(t *testing.T) {
	a := setBit(0, 62)
	a = setBit(a, 63)
	if a != 13835058055282163712 {
		t.Errorf("Error in setBit()")
	}
}

func TestIsBit(t *testing.T) {
	a := uint64(0)
	if isBit(a, 63) != false {
		t.Errorf("Error in setBit()")
	}
	a = setBit(a, 63)
	if isBit(a, 63) != true {
		t.Errorf("Error in setBit()")
	}
}
