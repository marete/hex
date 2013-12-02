package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"os"
)

var dec bool // Decode instead of encode

func init() {
	flag.BoolVar(&dec, "d", false, "decode instead of encoding")
}

func encode() {
	var in [4096]byte
	var out [4096 * 2]byte // Hex encoding produces twice the bytes
	for {
		n, err := os.Stdin.Read(in[:])

		// We'll check the error later. Deal with what we have now
		m := hex.EncodedLen(n)
		hex.Encode(out[:m], in[:n])
		_, wErr := os.Stdout.Write(out[:m])
		if wErr != nil {
			fmt.Fprintf(os.Stderr, "ERROR: encode(): %v\n", wErr)
			os.Exit(1)
		}

		if err != nil {
			if err == io.EOF {
				return
			}

			fmt.Fprintf(os.Stderr, "ERROR: encode(): %v\n", err)
			os.Exit(1)
		}

	}
}

func decode() {
	var in [4096]byte
	var out [4096 / 2]byte // The output of the hex decoding yields is 1/2 the length of the input

	for {
		n, err := os.Stdin.Read(in[:])
		l := n
		// n should be a multiple of 2. What to do now if it is not?
		if n%2 != 0 {
			// We try and read an extra byte
			_, err := os.Stdin.Read(in[n : n+1])
			if err != nil {
				fmt.Fprintln(os.Stderr, "ERROR: decode(): Input must contain an even number of bytes")
				os.Exit(1)
			}
			l += 1
		}

		// Process what we've got so far
		m, dErr := hex.Decode(out[:], in[:l])
		// Write what we've got first, check any decode error later
		_, wErr := os.Stdout.Write(out[:m])
		if wErr != nil {
			fmt.Fprintf(os.Stderr, "ERROR: decode(): %v\n", wErr)
			os.Exit(1)
		}
		if dErr != nil {
			fmt.Fprintf(os.Stderr, "ERROR: decode(): %v\n", dErr)
			os.Exit(1)
		}

		// Finally handle any error from the Read() operation
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Fprintf(os.Stderr, "ERROR: decode(): %v\n",
					err)
			}
		}

	}
}

func main() {
	flag.Parse()
	if dec {
		decode()
	} else {
		encode()
	}
}
