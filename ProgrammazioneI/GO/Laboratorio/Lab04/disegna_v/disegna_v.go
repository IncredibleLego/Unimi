//Autore: Francesco Corrado

package main

import(
	"fmt"
)

func main(){
	var n, s1, s2 int
//	fmt.Print("Inserire un valore n: ")
	fmt.Scan(&n)
	if n == 1{
		s2=0
	}else if n == 2{
		s2=1
	}else{
		s2=1+2*(n-2)
	}
	for r:=1; r<=n; r++{
		for i:=1; i<=s1; i++{
			fmt.Print(" ")
		}
		s1++
		fmt.Print("*")
		for i:=1; i<=s2; i++{
			fmt.Print(" ")
		}
		s2-=2
		if r<n{
			fmt.Println("*")
		}
	}
	fmt.Println()
}