//Autore: Francesco Corrado
package main
import "fmt"
func main() {
	var inputGiorni int
	fmt.Println("Programma che converte un numero intero inserito di giorni in anni + settimane + giorni")
	giorniAnno:=365
	giorniSettimana:=7
	fmt.Print("Inserire il numero di giorni da convertire (int): ")
	fmt.Scan(&inputGiorni)
	anni:=inputGiorni/giorniAnno
	settimane:=(inputGiorni%giorniAnno)/giorniSettimana
	giorni:=(inputGiorni%giorniAnno)%giorniSettimana
	fmt.Println(inputGiorni, " giorni = ", anni, " anni + ", settimane, " settimane + ", giorni, " giorni")
}