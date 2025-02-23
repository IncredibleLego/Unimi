//Autore: Francesco Corrado

package main

import(
	"fmt"
)

func bubbleSort(a []int) {
	fmt.Println(a)
	l := len(a)
	for j:=0; j<len(a); j++{
		for i:=0; i < l-1; i++{
			if a[i] > a[i+1]{
				a[i], a[i+1] = a[i+1], a[i]
			}
		}
		l--
		fmt.Println(a)
	}
}

func selectionSort(a []int){
	fmt.Println(a)
	var minindex int
	n := len(a)
	for i:=0; i < n; i++{
		minindex = i
		for j:=i+1; j < n; j++{
			if a[j] < a[minindex]{
				minindex = j
			}
		}
		a[i], a[minindex] = a[minindex], a[i]
		fmt.Println(a)
	}
}

func insertionSort(a []int) {
	fmt.Println(a)
	l := len(a)
	for i:=1; i < l; i++{
		j:=1
		for j > 0{
			if a[j-1] > a[j]{
				a[j-1], a[j] = a[j], a[j-1]
			}
			j--
		}
		fmt.Println(a)
	}
}

func main(){
	var s, str[]int
	s = []int{9, 4, 3, 2, 1, 5, 8, -4}
	str = append(str, s...)
	fmt.Println("s, str:", s, str)
	fmt.Println("BubbleSort:")
	bubbleSort(str)
	str = s
	fmt.Println("s, str:", s, str)
	fmt.Println("SelectionSort:")
	selectionSort(str)
	str = s
	fmt.Println("s, str:", s, str)
	fmt.Println("InsertionSort:")
	insertionSort(str)
}

