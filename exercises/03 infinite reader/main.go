package main

import "golang.org/x/tour/reader"

type MyReader struct{}

func (r MyReader) Read(buffer []byte) (n int, err error) {
	for i := range buffer {
		buffer[i] = 'A'
	}
	return len(buffer), nil
}

func main() {
	reader.Validate(MyReader{})
}
