//Autore: Francesco Corrado

package main

import(
	"fmt"
)

func main() {
	var n, somma int
	const d=10
	maxSum := 0
	for {
			fmt.Scan(&n)
			if n == 999{
				break
			}
			x:=n
			for{
				somma += x % d
				x = x / d
				if x < d{
					somma += x
					break
				}
			}
			if somma > maxSum{
				maxSum = somma
			}
			somma=0
	}
	if maxSum > 0{
		fmt.Println(maxSum)
	}
}