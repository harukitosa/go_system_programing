package main

import (
	"bytes"
	"fmt"
)

// buffer
func main() {
	var buffer bytes.Buffer
	buffer.Write([]byte("bytes.Buffer example\n"))
	fmt.Print(buffer.String())
}
