//Autore: Francesco Corrado

//Es preso da lezione del 28 Novembre

package main

import(
	"fmt"
)

type Esame struct{
	nome string
	proped []string
}

func tuttiGliOrdiniPossibili(cds []Esame) []string {
	if len(cds) == 0{
		return []string{""}
	}
	var res []string
	for i, primo := range cds{
		// Creo una copia di cds che NON contiene l'i-esimo elemento e
		//mette in testa primo
		var copia []Esame
		for j := 0; j < len(cds); j++{
			for j != i {
				copia = append(copia, cds[j])
			}
		}
		// Chiamata ricorsiva
		t := tuttiGliOrdiniPossibili(copia)
		for _, ordine := range t {
			if ordine != "" {
				res = append(res, primo.nome + "," + ordine)
			}else{
				res = append(res, primo.nome + ordine)
			}
		}
	}
	return res
}

func propedeuticitaOk (ordine []string, cds []Esame) bool {
	for i :=0; i<len(ordine); i++{
		esame := ordine[i]
		//Certco centro cds un esame di nome Esame
		var x Esame
		for _, x = range cds {
			if x.nome == esame{
				break
			}
		}
		//Ho trovato l'esame e si chiama x
		for _, prop := range x.proped {
			var j int
			for j = 0; j < i; j++{
				if ordine[j] == prop{
					break
				}
			}
			// Se non lo ho trovato
			if !(j<i){
				return false
			}
		}
	}
	return true
}

func ordiniPossibili(cds []Esame) []string{
	var res []string
	tutti := tuttiGliOrdiniPossiboli(cds)
	for _, ordine := range{
		//Controllo del rispetto delle propedeuticità
		esami := strings.Split(ordine, ",") //Esami sostenuti in un certo ordine
		if propedeuticitaOk(esami, cds){
			res = append(res, oridne)
		}
	}
	return res
} 

func main(){
	var cds []Esame
	cds = []Esame{
		Esame{"A", []string{}},
		Esame{"B", []string{}},
		Esame{"C", []string{"A"}},
		Esame{"D", []string{"C", "B"}},
		Esame{"E", []string{"C", "D"}},
	}
	fmt.Println(tuttiGliOrdiniPossibili(cds))
}