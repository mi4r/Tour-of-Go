package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	res := make([][]uint8, dy)
	for i := range res {
		res[i] = make([]uint8, dx)
	}
	for y := range res {
		for x := range res[y] {
			res[y][x] = uint8(y * x)
		}
	}
	return res
}

func main() {
	pic.Show(Pic)
}
