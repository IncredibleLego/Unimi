package main

import(
	"fmt"
	"os"
	"strconv"
)

func main(){
	var max int
	var ripetuti []int
	var controllo bool
	spazi := make(map[int]int)
	numeri := os.Args[1:]
	for i:=0; i < len(numeri); i++{
		trova, _ := strconv.Atoi(numeri[i])
		if trova > max{
			max = trova
		}
	}
	for i:=0; i < len(numeri); i++{
		trova, _ := strconv.Atoi(numeri[i])
		bianchi := max - trova
		spazi[trova] = bianchi
	}
	for i:=0; i < max; i++{
		for j:=0; j < len(numeri); j++{
			trova, _ := strconv.Atoi(numeri[j])
			if spazi[trova] > 0{
				controllo = true
				fmt.Print(" ")
				for x := range ripetuti{
					if ripetuti[x] == trova{
						controllo = false
					}
				}
				if controllo == true{
					spazi[trova]--
				}
			}else{
					fmt.Print("*")
			}
			ripetuti = append(ripetuti, trova)
		}
		fmt.Println()
//		fmt.Println("mappa: ", spazi)
		ripetuti = nil
	}
}