package main

import "fmt"

func main() {
	/*
	 Il programma deve, dato un voto int come input chiesto all'utente, controllare
	 che il voto inserito sia valido (ovvero sia compreso tra 0 e 30) e in caso contrario
	 stampare un messaggio di errore
	*/

	var voto int

	fmt.Print("voto: ")
	fmt.Scan(&voto)

	if voto < 0 || voto > 30 {
		fmt.Println("voto non valido")
	}
}
