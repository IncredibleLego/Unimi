package main
import "fmt"
func main() {
    /*
    Programma che chiede in input un valore n int e procede a stampare tutti i valori interi da n a 0
	su una riga separati da spazi
    */
	var n int
	fmt.Print("un numero: ")
	fmt.Scan(&n)
	for i := n; i >= 0; i-- {
		fmt.Print(i, " ")
	}
	fmt.Println()
}
