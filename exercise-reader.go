package main

import (
	"io"
	"golang.org/x/tour/reader"
)
type rot13Reader struct {
	r io.Reader
}
type MyReader struct{}
// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (r MyReader) Read(b []byte) (n int,err error){
	for i:= range b {
		b[i] = 'A'
	}
	return len(b),nil
}
func main() {
	reader.Validate(MyReader{})
}