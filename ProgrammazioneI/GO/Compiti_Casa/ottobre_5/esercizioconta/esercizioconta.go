package main

import "fmt"

func main() {
	var x, y, z, w int
	x = 17
	y = 5
	z = 60
	z = z + (x % 10)
	fmt.Println("z = ", z)
	y = y*x - (z + 1)
	fmt.Println("Y = ", y)
	x = x % (z + y)
	fmt.Println("X = ", x)
	w = x + y + z
	fmt.Println("W ovvero x + y + z vale ", w)
}
