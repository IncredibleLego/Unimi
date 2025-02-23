package main

import(
	"fmt"
)

func Abbondante(n int) bool{
	var divisori []int
	var somma int
	check := true
	for i:=1; i < n; i++{
		if n % i == 0{
			for k := range divisori{
				if divisori[k] == i{
					check = false
				}
			}
			if check{
				divisori = append(divisori, i)
			}
		}
	}
	for i:=0; i < len(divisori); i++{
		somma += divisori[i]
	}
	if n < somma{
		return true
	}else{
		return false
	}
}

func main(){
	var n int
	fmt.Scan(&n)
	x := 2
	for i:=0; i < n; i++{
		for{
			if Abbondante(x){
				fmt.Println(x)
				x++
				break
			}
			x++
		}
	}
}