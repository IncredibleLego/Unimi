//Autore: Francesco Corrado

package main

import "fmt"

func main() {
	const K = 10
	var prev, curr int
	for i := 1; i < K; i++ {
		fmt.Scan(&curr)
		if curr <= prev {
			ascending = true
		}
		prev = curr
	}
	if !ascending {
		fmt.Print("non ")
	}
	fmt.Println("crescente")
    fmt.Println(sum)
}

sum = curr
ascending = false
ascending = true
sum += curr
fmt.Scan(&curr)



/* COMANDI DA ISTRUZIONI
package main

import "fmt"

func main() {
	const K = 10
	var prev, curr int

       ....

	if !ascending {
		fmt.Print("non ")
	}
	fmt.Println("crescente")
    fmt.Println(sum)
}

ascending = false
sum += curr
prev = curr
fmt.Scan(&curr)
ascending = true
if curr <= prev {
sum = curr
for i := 1; i < K; i++ {
fmt.Scan(&curr)
}
}
*/