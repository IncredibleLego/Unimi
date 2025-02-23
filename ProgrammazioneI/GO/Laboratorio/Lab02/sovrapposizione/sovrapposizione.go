//Autore: Francesco Corrado

package main

import(
	"fmt"
)

func main(){
	var gg1, gg2, start1, start2, end1, end2, c int
	for {
		fmt.Print("Inserire i valori di appuntamento 1 (gg, start, end): ")
		fmt.Scan(&gg1, &start1, &end1)
		if (gg1 >= 1 && gg1<=31) && (start1 >= 0 && start1<=24) && (end1 >= 0 && end1<=24 && end1 > start1){
			break
		}else{
			fmt.Println("Uno o piu valori inseriti non sono validi, riprovare")
		}
	}
	for {
		fmt.Print("Inserire i valori di appuntamento 2 (gg, start, end): ")
		fmt.Scan(&gg2, &start2, &end2)
		if (gg2 >= 1 && gg2<=31) && (start2 >= 0 && start2<=24) && (end2 >= 0 && end2<=24 && end2 > start2){
			break
		}else{
			fmt.Println("Uno o piu valori inseriti non sono validi, riprovare")
		}
	}
	if gg2 == gg1{
		if (start2>=start1 && start2 <=end1) || (end2>=start1 && end2<=end1){
			c++
		}
	}
	if c == 0{
		fmt.Println("non si sovrappongono")
	}else{
		fmt.Println("si sovrappongono")
	}
}