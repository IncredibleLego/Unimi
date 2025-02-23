//Autore: Francesco Corrado

package main

import(
	"fmt"
	"math"
)

func main(){
	var n, x, N, i int
	fmt.Println("Programma che converte in decimale un numero binario inserito")
	fmt.Print("Inserisci un valore binario: ")
	fmt.Scan(&n)
	fmt.Print("Il numero ", n, "convertito in decimale è uguale a ")
	for n != 0{
		x = n % 10
		n/=10
		N+= x * int(math.Pow(2, float64(i)))
		i++
	}
	fmt.Println(N)
}



/* SOLUZIONE ORIGINALE, soluzione soprastante trovata su internet e modificata

func main(){
	var n, c, somma, a, y int
	var b float64
	fmt.Println("Programma che converte in decimale un numero binario inserito")
	fmt.Print("Inserisci un numero da convertire: ")
	fmt.Scan(&n)
	c = contaCifre(n)
	b = 0
	for i:=math.Pow(10,float64(c)-1); c>0; i/=10{
		if b == 0{
			y = 0
		}else{
			y = int(math.Pow(10, b))
		}
		fmt.Println("=", int(i) -y, "Int: ", int(i), "Y= ", y)
		fmt.Println("Condizione if: ", n/int(i)- y)
		if n/int(i)- y == 1{
			a=1
		}else{
			a=0
		}
		fmt.Println("Cifra numero: ", c, " Operazione = ", a, "* 2*", c, "=", a * int((math.Pow(2, float64(c)-1))))
		somma += a * int((math.Pow(2, float64(c)-1)))
		c--
		b++
	}
	fmt.Println("Il numero", n, "convertito in binario è uguale a", somma)
}

func contaCifre (n int) int{
	var contatore int
	for n/10 != 0{
		contatore++
		n = n/10
	}
	return contatore+1
} */