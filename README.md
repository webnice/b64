# b64 [![Build Status](https://drone.io/github.com/webdeskltd/b64/status.png)](https://drone.io/github.com/webdeskltd/b64/latest) [![Travis status](https://travis-ci.org/webdeskltd/b64.svg?branch=master "travis status")](https://travis-ci.org/webdeskltd/b64/#) [![GoDoc](https://godoc.org/github.com/webdeskltd/b64?status.png)](http://godoc.org/github.com/webdeskltd/b64)

Go library convert Int64 or Uint64 variable to string(11 symbols), and convert 11 symbols string back to int64 or uint64

### What is it?
Library converts a single intenger to a string of 1 to 11 characters. Conversion is to reduce number characters when writing a long string of numbers.

Example: `1048575 (7 symbols) -> "D___" (4 symbols)`

Yes, it is a regular base64 for use in a URL, but it works faster than the standard library "encoding/base64"
And it made a little more comfortable when using a conversion of only one number