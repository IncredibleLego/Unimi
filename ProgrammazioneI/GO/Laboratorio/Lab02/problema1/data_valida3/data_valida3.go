//Autore: Francesco Corrado

package main

import "fmt"

func main() {
	var g, m, a, x int
	fmt.Println("Programma che data una data verifica che la data sia valida")
	for {
		fmt.Print("Inserire giorno mese e anno della data (anno => 0): ")
		fmt.Scan(&g, &m, &a)
		if a<0{
			fmt.Println("L'anno inserito è minore di 0. Riprovare")
		}else{
			break
		}
	}
	if g>0 && g<32 && (m==1 || m==3 || m==5 || m==7 || m==8 || m ==10 || m==12){
		x++
	}else if g>0 && g<31 && (m==4 || m==6 || m==9 || m==11){
		x++	
	}else if m==2 {
		if verificaBisestile(a)==1 {
			if g>0 && g<30{
				x++
			}
		}else{
			if g>0 && g<29{
				x++
			}
		}
    }
	if x==1{
		fmt.Println("data valida")
	}else {
		fmt.Println("data non valida")
	}
}


func verificaBisestile (x int) int{
	if ((x%4==0 && x%100!=0) || x%400==0) && x>1581{
		return 1 //L'anno è bisestile
	}else{
		return 0 //L'anno non è bisestile
	}
}
