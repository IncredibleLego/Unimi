//Autore: Francesco Corrado

package main

import(
	"fmt"
)

const DIM = 5

func reverse (array [DIM]int) [DIM]int{
	for i:=0; i < DIM/2; i++{
		array[i], array[DIM-i-1] = array[DIM-i-1], array[i]
	}
	return array
}

func switchFirstLast (array [DIM]int) [DIM]int{
	array[0], array[DIM-1] = array[DIM-1], array[0]
	return array
}

func main(){
	var arr[DIM]int
	for i:=0; i<DIM; i++{
		arr[i] = i
	}
	fmt.Println("array:", arr)
	arr = reverse(arr)
	fmt.Println("reverse:", arr)
	arr = switchFirstLast(arr)
	fmt.Println("switchFirstLast:", arr)
}