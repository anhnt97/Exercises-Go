package main

import (
	"io"
	"os"
	"strings"
)
type rot13Reader struct {
	r io.Reader
}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (r MyReader) Read(b []byte) (n int,err error){
	for i:= range b {
		b[i] = 'A'
	}
	return len(b),nil
}
func main() {

}