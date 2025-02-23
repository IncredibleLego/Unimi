//Autore: Francesco Corrado

package main

import "fmt"

func main() {
	var giorno, mese, anno, mancanti int
	fmt.Println("Programma che data in imput la data di nascita di una persona stabilisce se oggi (10/10/2023) essa sia maggiorenne")
	fmt.Print("Inserisci la data di nascita della persona (giorno mese anno) separati da uno spazio: ")
	fmt.Scan(&giorno, &mese, &anno)
	maggiorenne := 30 * 12 * 18
	if giorno < 10 {
		mancanti += 10 - giorno
		if mese < 10 {
			mancanti += (10 - mese) * 30
		} else {
			mancanti += (22 - mese) * 30
		}
		mancanti += (2023 - anno) * 360
	} else {
		mancanti += 40 - giorno
		if mese+1 < 10 {
			mancanti += (10 - mese + 1) * 30
		} else {
			mancanti += (22 - mese + 1) * 30
		}
		mancanti += (2023 - anno + 1) * 360
	}
	if mancanti > maggiorenne {
		fmt.Println("Congratulazioni! La persona nata il", giorno, mese, anno, "è maggiorenne al 10 10 2023")

	} else {
		fmt.Println("Mi dispiace, ma putroppo la persona nata il", giorno, mese, anno, "non è maggiorenne al 10 10 2023")
	}
}
