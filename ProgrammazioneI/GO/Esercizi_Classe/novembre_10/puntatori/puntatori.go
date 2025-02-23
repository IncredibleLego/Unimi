//Autore: Francesco Corrado

package main

import (
	"fmt"
)

/*
func main(){
    var x int
	var p *int
	var q ** int
	x = 7
	p = &x
	q = &p
	*p = 50
	**q = *p - (x+1) // x = x - (x+1)
	(*p)++ //x++
	(**q)-- //x--
	fmt.Println(x)
} */

func main() {
	var x int
	var p *int
	var q **int
	x = 7
	p = &x
	q = &p
	p = new(int)
	*p = 50
	q = new(*int)
	*q = &x
	**q = 5
	*q = new(int)
	**q = 12
	fmt.Println(x, *p, **q)
}
