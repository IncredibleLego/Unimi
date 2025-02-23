package main

import "fmt"

func main() {
	/*
	 Programma che chuede all'utente di scegliere tramite due numeri tra un triangolo o un rettangolo
	 e successivamente chiede se si vuole calcolare area o perimetro. Infine stampa la formula
	 per l'operazione richiesta
	*/
	const (
		RETTANGOLO = 1
		TRIANGOLO  = 2
		PERIMETRO  = 1
		AREA       = 2
	)
	var figura, operazione int

	fmt.Println("di che figura si tratta?")
	fmt.Print("rettangolo (1) o triangolo (2)? ")
	fmt.Scan(&figura)
	fmt.Println("cosa vuoi calcolare?")
	fmt.Print("perimetro (1) o area (2)? ")
	fmt.Scan(&operazione)

	if figura == RETTANGOLO {
		if operazione == PERIMETRO {
			fmt.Println("formula = (base + altezza) * 2")
		} else { //operazione == AREA
			fmt.Println("formula = base * altezza")
		}
	} else { //figura == TRIANGOLO
		if operazione == PERIMETRO {
			fmt.Println("formula = lato1 + lato2 + lato3 ")
		} else { //operazione == AREA
			fmt.Println("formula = (base * altezza)/2")
		}
	}
}
/* DOMANDE */
/*
- D1. Quanti diversi messaggi da stampare ci sono?
- R1. 8

- D2. Sarebbe possibile ridurre il numero di if/else? Perché?
- R2. No, sono gli IF minimi possibili
*/
