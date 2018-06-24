package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		fmt.Println("sum:", sum)
		return sum
	}
}

func main() {
	pos := adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			"i:", i,
			pos(i),
		)
	}
	neg := adder()
	for j := 0; j < 10; j++ {
		fmt.Println(
			"j:", j,
			neg(-2*j),
		)
	}
}
