package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	slicesPic := make([][]uint8, dy)
	for i := 0; i < dy; i++ {
		slicesPic[i] = make([]uint8, dx)
		for j := 0; j < dx; j++ {
			slicesPic[i][j] = uint8(i^j+(i+j)/2)
			// slicesPic[i][j] = uint8(i*j)
			// slicesPic[i][j] = uint8(i^j)
		}
	}
	return slicesPic
}

func main() {
	pic.Show(Pic)
}