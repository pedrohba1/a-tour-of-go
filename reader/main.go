package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.

func (r MyReader) Read(bytes []byte) (int, error) {
	for i := range bytes {
		bytes[i] = 'A' // ASCII value for 'A' is 65
	}
	return len(bytes), nil // return the number of bytes written and no error
}

func main() {
	reader.Validate(MyReader{})
}
