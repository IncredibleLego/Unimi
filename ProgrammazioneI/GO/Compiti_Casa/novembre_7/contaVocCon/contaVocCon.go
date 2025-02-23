//Autore: Francesco Corrado

func contaVocCon (s string) (v int, c int){
	/* Funzione che data in input una stringa permette di calcolare il numero
	di vocali e consonanti presenti in essa*/
	for i:=0; i < len(s); i++{
		cr:=rune(s[i])
		if cr == 'a' || cr == 'e' || cr == 'i' || cr == 'o' || cr == 'u'{
			v++
		}else{
			c++
		}
	}
	return v, c
}