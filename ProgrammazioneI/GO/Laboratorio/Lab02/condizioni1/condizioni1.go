//Autore: Francesco Corrado

package main

import(
	"fmt"
	"math"
)

func main() {
	var n int
	var x float64
	const DISTANZA = 1e-6
	fmt.Println("Programma che testa una serie di condizioni su degli input inseriti dall'utente e restituisce valore true or false")
	fmt.Print("Inserire un valore n (int): ")
	fmt.Scan(&n)
	fmt.Print("Uguale a 10: ")
	fmt.Println(n == 10)
	fmt.Print("Diverso da 10: ")
	fmt.Println(!(n == 10))
	fmt.Print("Diverso da 10 e 20: ")
	fmt.Println(!(n == 10) && !(n == 20))
	fmt.Print("Diverso da 10 o da 20: ")
	fmt.Println(!(n == 10) || !(n == 20))
	fmt.Print("Maggiore o uguale a 10: ")
	fmt.Println(n >= 10)
	fmt.Print("Compreso tra 10 e 20, nell'intervallo [10,20]: ")
	fmt.Println(n>=10 && n<=20)
	fmt.Print("Compreso tra 10 e 20, nell'intervallo (10,20]: ")
	fmt.Println(n>10 && n<=20)
	fmt.Print("Compreso tra 10 e 20, nell'intervallo [10,20): ")
	fmt.Println(n>=10 && n<20)
	fmt.Print("Esterno all'intervallo [10,20]: ")
	fmt.Println(n<10 || n>20)
	fmt.Print("Tra 10 e 20 (nell'intervallo [10,20]) o tra 30 e 40 ([30,40]): ")
	fmt.Println((n>=10 && n<=20)||(n>=30 && n<=40))
	fmt.Print("Multiplo di 4 ma non di 100: ")
	fmt.Println(n%4==0 && !(n%100==0))
	fmt.Print("Dispari e compreso tra 0 e 100 [0,100]: ")
	fmt.Println(!(n%2==0) && n>=0 && n<=100)
	fmt.Print("Inserire un valore x (float64): ")
	fmt.Scan(&x)
	fmt.Print("Vicino a 10.0 (non più lontano di 10^-6): ")
	fmt.Println(math.Abs(x-10) <= 1e-6) //Inserito dopo aver visto soluzione
//	fmt.Println(x<DISTANZA)
}

//math.Abs(x-10)