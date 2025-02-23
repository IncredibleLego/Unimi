package main
import "fmt"
func main() {
    /*
    Programma che chiede un valore int in input e stampa tutti i numeri pari da 0 al numero stesso
    */
	var n int
	fmt.Print("un numero: ")
	fmt.Scan(&n)
	for i := 0; i <= n; i = i + 2 {
		fmt.Print(i, " ")
	}
	fmt.Println()
}
