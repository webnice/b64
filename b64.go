package b64

import (
	"errors"
)

const (
	encodeBase string = `ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_`
	mask6bit   uint64 = 0x3F
)

var (
	mapBase                    map[int]byte
	mapKey                     map[byte]int
	ErrorUnknownSymbolInString error = errors.New(`Unknown symbol in string`)
)

type (
	b64Marshal struct {
		SourceUint64 uint64
		Bytes        [11]byte
	}
	b64Unarshal struct {
		b64Marshal
		SourceString string
		LastError    error
	}
	B64 interface {
		String() string // Return 64 bit integer in 64-based scale of notation
		Uint64() uint64 // Return 64 bit unsigned integer
		Int64() int64   // Return 64 bit signed integer
		Bits() string   // Return 64 bit integer by bits - string of 0 or 1, string length 64 symbols
		Error() error   // Return last error after conversion
	}
)

func NewUint64(inp uint64) B64 {
	var resp = []B64{
		newUint64(inp),
	}
	return resp[0]
}

func NewInt64(inp int64) B64 {
	var resp = []B64{
		newInt64(inp),
	}
	return resp[0]
}

func NewString(inp string) B64 {
	var resp = []B64{
		newString(inp),
	}
	return resp[0]
}
