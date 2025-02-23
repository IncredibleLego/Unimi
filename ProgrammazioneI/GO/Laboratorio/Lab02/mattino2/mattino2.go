//Autore: Francesco Corrado

package main

import "fmt"

func main(){
	var ora, min int
	fmt.Scan(&ora, &min)
	if (ora == 5 && min >=30) || (ora == 12 && min<30) || (ora > 5 && ora <12) {
		fmt.Println("mattino")
	}
}