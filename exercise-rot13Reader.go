package main

import (
	"io"
	"os"
	"strings"
)
type rot13Reader struct {
	r io.Reader
}
func rot13(x byte) byte {
    switch {
    case x >= 'A' && x <= 'M' || x >= 'a' && x <= 'm':
        x = x + 13
    case x >= 'N' && x <= 'Z' || x >= 'n' && x <= 'z':
        x = x - 13
    }
    return x
}
func (r13 rot13Reader) Read(b []byte) (int, error){
	n,err := r13.r.Read(b)
	for i:= range b {
		b[i] = rot13(b[i])
	}
	return n,err
}
func main() {
	s := strings.NewReader("Jul qvq gur puvpxra pebff gur ebnq?")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}