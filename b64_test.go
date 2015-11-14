package b64

import (
	"testing"
)

func TestNewMarshalUint64(t *testing.T) {
	var o B64
	o = NewMarshalUint64(0xFFFFFFFFFFFFFFFF)
	if o == nil {
		t.Errorf("Error in NewMarshalUint64()")
	}
	if o.Error() != nil {
		t.Errorf("Error in NewMarshalUint64()")
	}
	if o.Int64() != -1 {
		t.Errorf("Error in NewMarshalUint64()")
	}
	if o.Uint64() != 0xFFFFFFFFFFFFFFFF {
		t.Errorf("Error in NewMarshalUint64()")
	}
	if o.Bits() != `1111111111111111111111111111111111111111111111111111111111111111` {
		t.Errorf("Error in NewMarshalUint64()")
	}
}

func TestNewMarshalInt64(t *testing.T) {
	var o B64
	o = NewMarshalInt64(-1)
	if o == nil {
		t.Errorf("Error in NewMarshalInt64()")
	}
	if o.Uint64() != 0xFFFFFFFFFFFFFFFF {
		t.Errorf("Error in NewMarshalInt64()")
	}
}

func TestNewUnmarshalString(t *testing.T) {
	var o B64
	o = NewUnmarshalString(`1`)
	if o == nil {
		t.Errorf("Error in NewUnmarshalString()")
	}
	if o.Uint64() != 53 {
		t.Errorf("Error in NewMarshalInt64()")
	}
}
