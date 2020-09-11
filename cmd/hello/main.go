// +build example
package main

import (
	"github.com/gaond/gocv"
)

func main() {
	webcam, _ := gocv.OpenVideoCapture(0)
	window := gocv.NewWindow("Hello")
	img := gocv.NewMat()

	var myData uint8 = uint8(24)

	myDataPtr := &myData

	img.SetData(myDataPtr)

	for {
		webcam.Read(&img)
		window.IMShow(img)
		window.WaitKey(1)
	}
}
