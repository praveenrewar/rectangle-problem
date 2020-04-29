package main

import "os"

//InputFile is used for the input json
var InputFile = os.Getenv("SOURCE_FILE")

func main() {
	InitFileSystem(&osFS{})
	InitStdOut(&osFMT{out: os.Stdout})
	getIntersections()
}
