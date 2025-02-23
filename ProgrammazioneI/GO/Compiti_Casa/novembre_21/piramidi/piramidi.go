//Autore: Francesco Corrado

package main

import(
    "fmt"
)

func spazi(s int, c int) bool{
	for x:=0; x<s; x++{
		if c == 0{
			fmt.Print(" ")
		}else{
			fmt.Print("*")
		}
	}
	return true
}

func main(){
    var n, spazio1, spazio2 int
	fmt.Print("Inserisci la dimensione della piramide: ")
	fmt.Scan(&n)
	spazio2 = n*2 -5 //Si può usare spazio2 in riga 32, e mettere tutti output in una riga
	fmt.Println("")
	for r:=0; r<n; r++{
		if n == 1{
			fmt.Println("*")
			break
		}else if r == 0{
			fmt.Print("*")
			spazi((n-2)*2+1, 0)
			fmt.Print("*")
		}else if r == n-1 {
			spazi(n*2-1, 1)
		}else{
			fmt.Print("*")
			spazi(spazio1, 0)
			fmt.Print("*")
			spazi(spazio2, 0)
			fmt.Print("*")
			spazi(spazio1, 0)
			fmt.Print("*")
			spazio1++
			spazio2-=2
		}
		fmt.Println()
	}
}