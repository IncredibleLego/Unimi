package main

import "fmt"

func main() {
	/* 	 
	Programma che chiede all'utente di inserire un voto tra 0 e 100 inclusi (in caso contrario
	stampa un messaggio di errore e termina il programmma) e successivamente stampa il voto in
	lettere corrsipondente al voto inserito secondo lo schema: >=90: A, >=80: B, >=70: C, >=60: D,
	altrimenti F
	*/

	var voto int
	fmt.Print("voto? ")
	fmt.Scan(&voto)

	if voto < 0 || voto > 100 {
		fmt.Println("voto non valido")
		return //interrompe l'esecuzione del programma
	}

	if voto >= 90 {
		fmt.Println("A")
	} else if voto >= 80 {
		fmt.Println("B")
	} else if voto >= 70 {
		fmt.Println("C")
	} else if voto >= 60 {
		fmt.Println("D")
	} else {
		fmt.Println("F")
	}
}
/* DOMANDE */
/*
- D1. Per quali valori è vera la condizione del primo if? (quello subito dopo la Scan)
- R1. 0 - 100

- D2. Per quali valori viene eseguita la seguente istruzione?
    fmt.Println("B")
- R2. >=80 e <90

- D3. Giustificate la vostra risposta alla domanda D2
- R3. Tutti i valori >=90 una volta raggiunto il primo if saranno stampati li, senza scendere al
	  secondo if contente la stampa della B

- D4. Il programma è impostato correttamente o sarebbe stato (più) corretto scrivere invece così?
 } else if voto >= 80 && voto < 90 {
- R4. è impostato correttamente

- D5. Perché?
- R5. Perchè l'ordine in cui sono messi gli IF fa in modo che non siano necessari controlli sui
	  valori superiori
*/
