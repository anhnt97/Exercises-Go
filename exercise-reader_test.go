package main

import (
	"testing"
)

func TestReader(t *testing.T) {
	byteTest := make([]byte, 1024, 2048)
	i, o := 0, 0
	for ; i < 1<<20 && o < 1<<20; i++ { // test 1mb
		n, err := MyReader.Read(MyReader{}, byteTest)
		for i, v := range byteTest[:n] {
			if v != 'A' {
				t.Errorf("Byte %x at offset %v, want 'A'\n", v, o+i)
				return
			}
		}
		o += n
		if err != nil {
			t.Errorf("Read error: %v\n", err)
			return
		}
	}
	if o == 0 {
		t.Errorf("Read zero bytes after %d Read calls\n", i)
		return
	}
}
