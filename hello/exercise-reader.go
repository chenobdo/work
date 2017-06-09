package main

import "code.google.com/p/go-tour/reader"

type MyReader struct{}

// TODO: Add a Read(byte) (int, error) method to MyReader.

func (rd MyReader) Read(b []byte) (n int, err error) {
	i := 0
	for i < len(b) {
		b[i] = 'A'
		i++
	}
	return i, nil
}

func main() {
	reader.Validate(MyReader{})
}
