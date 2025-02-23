//Autore: Francesco Corrado

package main

import(
	"fmt"
)

func main(){
	var n, c, x int
	x=1
	fmt.Println("Programma che stabilisce in che posizione si trova la prima cifra pari di un numero")
	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&n)
	for i:=1; i > 0; i*=10{
		c = n / i
		if c == 0{
			x = -1
		}
		if pari(c){
			break
		}
		x++
	}
	fmt.Println("La prima ci cifra pari si trova nella posizione", x)
}

func pari(n int) bool{
	if n % 2 == 0{
		return true
	}else{
		return false
	}
}