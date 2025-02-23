package main

import(
	"fmt"
	"os"
	"strconv"
)

func main(){
	var s []string
	var err error
	var x, y int
	s = os.Args[1:]
	if len(s) > 0{
		for i:=0; i<len(s); i++{
			if i == 0{
				x, err = strconv.Atoi(s[i])
				if err != nil{
					fmt.Println("elemento in posizione 1 non valido")
					break
				}
			}else{
				y = x
				x, err = strconv.Atoi(s[i])
				if err != nil || x % 2 == 0 && y % 2 == 0 || x % 2 != 0 && y % 2 != 0 {
					fmt.Println("elemento in posizione", i+1, "non valido")
					break
				}
			}
			if i == len(s)-1{
				fmt.Println("sequenza valida")
			}
		}
	}else{
		fmt.Println("nessun valore in ingresso")
	}
}