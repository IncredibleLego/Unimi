//Autore: Francesco Corrado

package main

import "fmt"

func main() {
	var x int
	x = 1
	for i := 0; i < 100; i++ {
		fmt.Println(x)
		x *= 3
	}
}
