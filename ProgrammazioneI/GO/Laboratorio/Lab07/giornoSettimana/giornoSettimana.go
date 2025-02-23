//Autore: Francesco Corrado

package main

import(
	"fmt"
	"os"
)

func main(){
	var input string
	var g, pos, day int
	giorniDellaSettimana := [7]string{"lun", "mart", "merc", "giov", "ven", "sab", "dom"}
	input = os.Args[1]
	for i:=0; i < len(giorniDellaSettimana); i++{
		if input == giorniDellaSettimana[i]{
			pos = i + 1
			break
		}
		if i == len(giorniDellaSettimana)-1{
			fmt.Println("Il giorno inserito non è un valore valido, arresto del programma")
			os.Exit(0)
		}
	}
	for{
		fmt.Scan(&g)
		fmt.Println("Numero: ", g)
		if g == -1{
			break
		}
		if g < 0 || g > 365{
			fmt.Println("Il numero inserito non è valido (deve essere compreso in (0,365)")
		}else{
			x := (pos+g) % 7
			if pos + x > 7{
				day = pos + x - 7
			}else{
				day = pos + x
			}
			fmt.Println("Il giorno dell'anno numero", g, "corrisponde al giorno", giorniDellaSettimana[day-1])
		}
	}
}