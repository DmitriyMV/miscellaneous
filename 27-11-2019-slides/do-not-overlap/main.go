// +build ignore,OMIT
package main

import "io"

type ReadWriteCloser interface {
	io.Reader
	io.Writer
	io.Closer
}

func main() {}
