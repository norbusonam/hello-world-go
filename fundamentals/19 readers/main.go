package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {

	var reader io.Reader = strings.NewReader("Hello, world!")

	buffer := make([]byte, 8)
	for n, err := reader.Read(buffer); err != io.EOF; n, err = reader.Read(buffer) {
		fmt.Printf("%d bytes read: \"%s\"\n", n, buffer[:n])
	}

}
