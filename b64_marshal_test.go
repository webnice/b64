package b64

import (
	"math/rand"
	"testing"
	"time"
)

func TestNewUint64(t *testing.T) {
	if len(encodeBase) != 64 {
		t.Errorf("Error len(encodeBase)")
		return
	}
	if mapBase == nil {
		t.Errorf("Error in init()")
		return
	}
	if len(mapBase) != len(encodeBase) {
		t.Errorf("Error in init()")
		return
	}

	a := newUint64(1)
	if a.Error() != nil {
		t.Errorf("Error in newUint64()")
		return
	}
	if a.Uint64() != uint64(1) {
		t.Errorf("Error in newUint64()")
		return
	}
	if a.Int64() != int64(1) {
		t.Errorf("Error in newUint64()")
		return
	}
	if a.String() != `B` {
		t.Errorf("Error in newUint64()")
		return
	}
	if a.Bits() != `0000000000000000000000000000000000000000000000000000000000000001` {
		t.Errorf("Error in newUint64()")
		return
	}
}

func TestNewInt64(t *testing.T) {
	a := newInt64(-1)
	if a.Error() != nil {
		t.Errorf("Error in newInt64()")
		return
	}
	if a.Bits() != `1111111111111111111111111111111111111111111111111111111111111111` {
		t.Errorf("Error in newInt64()")
		return
	}
}

func TestB64MarshalUint64(t *testing.T) {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	r := uint64(rnd.Int63n(time.Now().UnixNano()))
	a := newUint64(r)
	if a.Uint64() != r {
		t.Errorf("Error in Uint64()")
	}
}

func TestB64MarshalInt64(t *testing.T) {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	r := rnd.Int63n(time.Now().UnixNano())
	a := newInt64(r)
	if a.Int64() != r {
		t.Errorf("Error in Int64()")
	}
	a = newUint64(0xFFFFFFFFFFFFFFFF)
	if a.Int64() != -1 {
		t.Errorf("Error in Int64()")
	}
}

func TestB64MarshalError(t *testing.T) {
	a := newInt64(-1000000)
	if a.Error() != nil {
		t.Errorf("Error in Error()")
	}
}

func TestB64MarshalconvertFromUint64(t *testing.T) {
	td := []struct {
		A uint64
		B string
	}{{0xFFFFFFFFFFFFFFFF, `__________P`}, {0x0, `AAAAAAAAAAA`}, {0xFFEE, `u_PAAAAAAAA`}}
	a := newUint64(0)
	for i := range td {
		a.SourceUint64 = td[i].A
		a.convertFromUint64()
		if string(a.Bytes[:]) != td[i].B {
			t.Errorf("Error in convertFromUint64(), return %v expected %v", string(a.Bytes[:]), td[i].B)
		}
	}
}

func BenchmarkB64MarshalconvertFromUint64AllBits(b *testing.B) {
	var vt uint64 = 0xFFFFFFFFFFFFFFFF
	obj := newUint64(vt)
	for i := 0; i < b.N; i++ {
		obj.convertFromUint64()
	}
}

func BenchmarkB64MarshalconvertFromUint64OneBits(b *testing.B) {
	var vt uint64 = 0x1
	obj := newUint64(vt)
	for i := 0; i < b.N; i++ {
		obj.convertFromUint64()
	}
}

func TestB64MarshalString(t *testing.T) {
	td := []struct {
		A uint64
		B string
	}{
		{0x0, `A`},
		{0xF, `P`},
		{0xFE, `D-`},
		{0xFF, `D_`},
		{0xFFF, `__`},
		{0xFFEE, `P_u`},
		{0xFFFF, `P__`},
		{0xFFFFF, `D___`},
		{0xFFFFFFFFFFFFFFFF, `P__________`},
		{0x3FFFFFFFFFFFFFFF, `D__________`},
		{0xf0FFFFFFFFFFFFFF, `PD_________`},
		{0xFFCCF, `D_zP`},
		{0xFFFF00FF0, `__8A_w`},
		{9223372036854775809, `IAAAAAAAAAB`},
	}
	for i := range td {
		obj := newUint64(td[i].A)
		if obj.String() != td[i].B {
			t.Errorf("Error in String(), return %v expected %v", obj.String(), td[i].B)
		}
	}

}

func BenchmarkB64MarshalStringAllBits(b *testing.B) {
	var vt uint64 = 0xFFFFFFFFFFFFFFFF
	obj := newUint64(vt)
	for i := 0; i < b.N; i++ {
		_ = obj.String()
	}
}

func BenchmarkB64MarshalStringOneBits(b *testing.B) {
	var vt uint64 = 0x1
	obj := newUint64(vt)
	for i := 0; i < b.N; i++ {
		_ = obj.String()
	}
}

func TestB64MarshalBits(t *testing.T) {
	td := []struct {
		B string
		A uint64
	}{
		{`1000000000000000000000000000000000000000000000000000000000000001`, 9223372036854775809},
		{`0100000000000000000000000000000000000000000000000000000000000010`, 4611686018427387906},
		{`0010000000000000000000000000000000000000000000000000000000000100`, 2305843009213693956},
		{`0001000000000000000000000000000000000000000000000000000000001000`, 1152921504606846984},
		{`0000100000000000000000000000000000000000000000000000000000010000`, 576460752303423504},
		{`0000010000000000000000000000000000000000000000000000000000100000`, 288230376151711776},
		{`0000001000000000000000000000000000000000000000000000000001000000`, 144115188075855936},
		{`0000000100000000000000000000000000000000000000000000000010000000`, 72057594037928064},
		{`0000000010000000000000000000000000000000000000000000000100000000`, 36028797018964224},
		{`0000000001000000000000000000000000000000000000000000001000000000`, 18014398509482496},
		{`0000000000100000000000000000000000000000000000000000010000000000`, 9007199254742016},
		{`0000000000010000000000000000000000000000000000000000100000000000`, 4503599627372544},
		{`0000000000001000000000000000000000000000000000000001000000000000`, 2251799813689344},
		{`0000000000000100000000000000000000000000000000000010000000000000`, 1125899906850816},
		{`0000000000000010000000000000000000000000000000000100000000000000`, 562949953437696},
		{`0000000000000001000000000000000000000000000000001000000000000000`, 281474976743424},
		{`0000000000000000100000000000000000000000000000010000000000000000`, 140737488420864},
		{`0000000000000000010000000000000000000000000000100000000000000000`, 70368744308736},
		{`0000000000000000001000000000000000000000000001000000000000000000`, 35184372350976},
		{`0000000000000000000100000000000000000000000010000000000000000000`, 17592186568704},
		{`0000000000000000000010000000000000000000000100000000000000000000`, 8796094070784},
		{`0000000000000000000001000000000000000000001000000000000000000000`, 4398048608256},
		{`0000000000000000000000100000000000000000010000000000000000000000`, 2199027449856},
		{`0000000000000000000000010000000000000000100000000000000000000000`, 1099520016384},
		{`0000000000000000000000001000000000000001000000000000000000000000`, 549772591104},
		{`0000000000000000000000000100000000000010000000000000000000000000`, 274911461376},
		{`0000000000000000000000000010000000000100000000000000000000000000`, 137506062336},
		{`0000000000000000000000000001000000001000000000000000000000000000`, 68853694464},
		{`0000000000000000000000000000100000010000000000000000000000000000`, 34628173824},
		{`0000000000000000000000000000010000100000000000000000000000000000`, 17716740096},
		{`0000000000000000000000000000001001000000000000000000000000000000`, 9663676416},
		{`0000000000000000000000000000000110000000000000000000000000000000`, 6442450944},
		{`1111111111111111111111111111111111111111111111111111111111111111`, 0xFFFFFFFFFFFFFFFF},
	}
	for i := range td {
		obj := newUint64(td[i].A)
		if obj.Bits() != td[i].B {
			t.Errorf("Error in Bits(), return %v expected %v", obj.Bits(), td[i].B)
		}
	}
}
