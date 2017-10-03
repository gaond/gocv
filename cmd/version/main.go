// What it does:
//
// 	This program outputs the current OpenCV library version to the console.
//
// How to run:
//
// 		go run ./examples/showinfo.go
//
// +build example

package main

import (
	"fmt"

	gocv ".."
)

func main() {
	fmt.Printf("go-gocv version: %s\n", gocv.Version())
	fmt.Printf("opencv lib version: %s\n", gocv.OpenCVVersion())
}