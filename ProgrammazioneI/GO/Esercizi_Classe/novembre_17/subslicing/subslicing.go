//Autore: Francesco Corrado

package main

import(
	"fmt"
)

func avg(x []int) float64{
	s:=0
	for _, t:=range x{
		s+=t
	}
	return float64(s)/float64(len(x))
}

func main() {
	var a []int
	a=[]int{71, 123, 43, 12, 131}
	fmt.Println(avg(a))
	var b[3]int
	b=[...]int{10, 20, 30}
	avg(b[:])
}