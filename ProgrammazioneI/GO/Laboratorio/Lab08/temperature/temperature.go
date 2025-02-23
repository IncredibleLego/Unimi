//Autore: Francesco Corrado

package main

import(
	"fmt"
	"sort"
)

func media(temp []int) int{
	var media int
	l := len(temp)
	for x:=0; x<l; x++{
		media = media + temp[x]
	}
	media = media/l
	return media
}

func mediana(temp []int) int{
	var mediana, m int
	l := len(temp)
	m = l/2
	if l % 2 == 0{
		mediana = (temp[m] + temp[m-1])/2
	}else{
		mediana = temp[m]
	}
	return mediana
}

func main(){
	var temperature []int
	var output []int
	var n, c int
	for{
		fmt.Scan(&n)
		if n == 999{
			break
		}
		temperature = append(temperature, n)
	}
	sort.Ints(temperature)
	fmt.Println("Temp ordinato:", temperature)
	fmt.Println("Media:", media(temperature))
	fmt.Println("Mediana:", mediana(temperature))
	for x:=0; x < len(temperature); x++{
		if temperature[x] < media(temperature){
			c++
		}
	}
	fmt.Println("Numero di temperature sotto la media:", c)
	if len(temperature) >= 3{
		c = 3
	}else{
		c = len(temperature)
	}
	fmt.Print("Temperature più basse: ")
	for x:=0; x < c; x++{
		fmt.Print(temperature[x], " ")
	}
	if len(temperature) >= 3{
		c = len(temperature) -4
	}
	fmt.Print("\nTemperature più alte: ")
	for x:=len(temperature)-1; x > c; x--{
		output = append(output, temperature[x])
	}
	sort.Ints(output)
	if len(temperature) >= 3{
		c = 3
	}
	for x:=0; x < c; x++{
		fmt.Print(output[x], " ")
	}
	fmt.Println()
}