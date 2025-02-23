//Autore: Francesco Corrado

package main

import(
	"fmt"
	"strings"
	"strconv"
)

func giorniInMese (mese int) int{
	var giorni int
	switch mese{
		case 2:
			giorni=28
		case 4, 6, 9, 11:
			giorni=30
		default:
			giorni=31
	}
	return giorni
}

func main(){
	var s string
	var n int
	for{
		fmt.Scan(&s)
		x:=strings.Index(s, "-")+1
		n, _ = strconv.Atoi(s[x:x+2])
		if n >= 1 && n <= 12{
			break
		}else{
			fmt.Println("Il mese inserito non è valido. Riprovare")
		}
	}
	fmt.Println("il mese", n, "ha", giorniInMese(n), "giorni")
}