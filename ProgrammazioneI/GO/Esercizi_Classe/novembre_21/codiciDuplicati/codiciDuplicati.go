//Autore: Francesco Corrado

func containsDuplicates(p []Persone) bool{
	for i:=0; i<len(p); i++{
		for j:=i+1; j<len(p); j++{
			if p[i].cf == p[j].cf{
				return true
			}
		}
	}
	return false
}