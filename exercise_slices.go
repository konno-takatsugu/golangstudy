package main

import "github.com/golang.org/x/tour/pic"

func Pic(dx, dy int) [][]byte {

	pic := make([][]byte, dy)

	for i := range pic {

		pic[i] = make([]byte, dx)

		for j := range pic {

			pic[i][j] = byte ((i + j)/2)
		}
	}
	return pic
}

func main() {
	pic.Show(Pic)
}
