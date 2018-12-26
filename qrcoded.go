package main

import (
	// "fmt"
	"image/png"
	"image"
	"io"
	// "os"
	// "log"
)
type Version int8
// func GenerateQRCode(w io.Writer, code string) error{
// 	img := image.NewNRGBA(image.Rect(0,0,21,21))
// 	fmt.Printf("Image : %+v",w)
// 	return png.Encode(w,img)
// }
func DeriveSizeFromVersion(version Version) int {
	return 4*int(version) + 17
}
func (version Version)PatternSize() int  {
	return 4*int(version) + 17
}
func GenerateQRCode(w io.Writer, code string, version Version) error{
	size := version.PatternSize()
	img := image.NewNRGBA(image.Rect(0,0,size,size))
	return png.Encode(w,img)
}
func main() {
	// file,err := os.Create("qrcode.png")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer file.Close()
	// err = GenerateQRCode(file, "555-2368")
	// if err != nil {
	// 	log.Fatal(err)
	// }
}
