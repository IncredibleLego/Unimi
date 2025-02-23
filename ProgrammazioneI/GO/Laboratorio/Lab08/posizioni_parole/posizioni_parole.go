//Autore: Francesco Corrado

package mattina

import(
	"fmt"
)

func main(){
//	mappa := make(map[string][]int)
	var x string
	fmt.Println("scrivi parole (ctrl D per terminare)")
	for i:=0; i < 20; i++{
		_, err := fmt.Scan(&x)
		if err != nil{
			fmt.Println("Input finito!")
			break
		}
		fmt.Println(x)
	}
}