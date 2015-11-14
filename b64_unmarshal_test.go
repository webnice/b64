package b64

import (
	"testing"
)

func TestInit(t *testing.T) {
	if len(encodeBase) != 64 {
		t.Errorf("Error len(encodeBase)")
		return
	}
	if mapKey == nil {
		t.Errorf("Error in init()")
		return
	}
	if len(mapKey) != len(encodeBase) {
		t.Errorf("Error in init()")
		return
	}
}

func TestNewString(t *testing.T) {
	a := newString(``)
	if a.Error() != nil {
		t.Errorf("Error in Error()")
		return
	}
	if a.Uint64() != 0 {
		t.Errorf("Error in newString()")
		return
	}
	if a.String() != `` {
		t.Errorf("Error in newString()")
		return
	}

	a = newString(`MAAAAAAAAAA`)
	if a.Error() != nil {
		t.Errorf("Error in Error()")
		return
	}
	if a.Uint64() != 13835058055282163712 || a.Int64() != int64(-4611686018427387904) {
		t.Errorf("Error in newString()")
		return
	}
	if a.String() != `MAAAAAAAAAA` {
		t.Errorf("Error in newString()")
		return
	}
}

func TestB64UnarshalConvertFromString(t *testing.T) {
	a := newString(``)
	a.SourceString = `(`
	a.convertFromString()
	if a.Error() == nil {
		t.Errorf("Error in convertFromString()")
		return
	}

	a.SourceString = `P__________`
	a.convertFromString()
	if a.Uint64() != 0xFFFFFFFFFFFFFFFF {
		t.Errorf("Error in convertFromString()")
		return
	}
	if a.Int64() != -1 {
		t.Errorf("Error in convertFromString()")
		return
	}
	if a.Bits() != `1111111111111111111111111111111111111111111111111111111111111111` {
		t.Errorf("Error in convertFromString()")
		return
	}

	// Overflow
	a.SourceString = `___________`
	a.convertFromString()
	if a.Uint64() != 0xFFFFFFFFFFFFFFFF {
		t.Errorf("Error in convertFromString()")
		return
	}

}

func BenchmarkB64UnarshalConvertFromStringOneBits(b *testing.B) {
	obj := newString(``)
	obj.SourceString = `B`
	for i := 0; i < b.N; i++ {
		obj.convertFromString()
	}
}

func BenchmarkB64UnarshalConvertFromStringAllBits(b *testing.B) {
	obj := newString(``)
	obj.SourceString = `P__________`
	for i := 0; i < b.N; i++ {
		obj.convertFromString()
	}
}
