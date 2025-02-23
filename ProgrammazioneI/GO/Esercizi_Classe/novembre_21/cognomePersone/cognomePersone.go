//Autore: Francesco Corrado

type Persona struct{
	nome, cognome string
	nascita data cf string
}

func conta(p []Persona, c string) int {
	var conta int
	for _; persona := range p{
		if persona.cognome == c{
			conta++
		}
	}
	return conta
}

func selez (p []Persona, c string){
	var risultato[]Persona
	for _; persona := range p{
		if persona.cognome == c{
			risultato = append(risultato, persona)
		}
	}
	return risultato
}