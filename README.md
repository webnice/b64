# b64
Go library convert Int64 or Uint64 variable to string(11 symbols), and convert 11 symbols string back to int64 or uint64

[![GoDoc](https://godoc.org/github.com/webnice/b64?status.png)](http://godoc.org/github.com/webnice/b64)

### What is it?
Library converts a single intenger to a string of 1 to 11 characters.
Conversion is to reduce number characters when writing a long string of numbers.

Example: `1048575` (7 symbols) -> `"D___"` (4 symbols)

Yes, it is a regular base64 for use in a URL, but it works faster than the standard library "encoding/base64"
And it made a little more comfortable when using a conversion of only one number

### How to use?

######Convert from number to string
	package main

	import "gopkg.in/webnice/b64.v1"

	func main() {
	    o := b64.NewUint64(10485251)
	    print(o.String(), "\n")
    	print(o.Uint64(), "\n")
	}

Result:

	n_4D
	10485251
--

######Convert from string to number

	package main

	import "gopkg.in/webnice/b64.v1"

	func main() {
		obj := b64.NewString("P____A____1")
		if obj.Error() != nil {
			panic(obj.Error())
		}
		print(obj.Uint64(), "\n")
		print(obj.Int64(), "\n")
	}

Result:

	18446744006063816693
	-67645734923

--

#### Dependencies

	NONE

#### Install
```bash
go get gopkg.in/webnice/b64.v1
```
