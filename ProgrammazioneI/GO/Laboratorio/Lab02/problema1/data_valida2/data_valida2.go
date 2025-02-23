//Autore: Francesco Corrado

package main

import(
	"fmt"
)

func main(){
	var g, m, x int
	fmt.Print("Inserire giorno e mese dell'anno: ")
	fmt.Scan(&g, &m)
	if g>0 && g<32 && (m==1 || m==3 || m==5 || m==7 || m==8 || m ==10 || m==12){
		x++
	}else if g>0 && g<31 && (m==4 || m==6 || m==9 || m==11){
		x++	
	}else if g>0 && g<29 && m==2{
		x++
	}
	if x==1{
		fmt.Println("data valida")
	}else {
		fmt.Println("data non valida")
	}
}