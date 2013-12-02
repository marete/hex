hex
====

`hex` is a small utility to decode from and encode to hexadecimal encoding as the encoding is defined and supported by the Go [encoding/hex](http://golang.org/pkg/encoding/hex/) package.

Usage
---------
`hex` reads input from standard input and writes output to standard output.

To decode: `cat foo.hex | hex -d`

To encode: `cat foo.bin | hex`
