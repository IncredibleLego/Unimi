//Autore: Francesco Corrado

func primoMarsenne(n int) bool{
	/* Funzione dato un numero n int stabilisce se
	si tratta di un numero primo di Marsenne (2^x-1) */
	for i:=0; math.Pow(2, float64(i)) < float64(n); i++{
		if math.Pow(2, float64(i)) == float64(n){
			return true
		}
	}
	return false
}