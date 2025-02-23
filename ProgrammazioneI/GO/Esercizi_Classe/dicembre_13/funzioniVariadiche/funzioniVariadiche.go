//Autore: Francesco Corrado

func sum1 (x ...int) int{ //Dalla funzione x  considerata una slice
	s:=0
	for _, v:=range x{
		s+=v
	}
	return s
}

func sum2 (x []int) int{ //Le due funzioni sono divere
	s:=0
	for _, v:=range x{
		s+=v
	}
	return s
}

func main(){
	x:=sum1(3, 1, 1, 3, 5, 5,6)
//	x:=sum1([]int{3, 4, 4, 3}) NON SI PUò FARE
//	y:=sum2(5, 6, 8) NON SI PUò FARE
	y:=sum2([]int{3, 2, 8, 3, 1})

	var s[]int
	s = []int{4, 4, 6, 2, 5, 2}
	s = (s...) //Trasforma s in un insieme di interi che si può passare a sum1

	var r, t []string
	r = append(r, "ciao")
	r = append(r, "x", "y", "z")
	t = append(t, "pippo", "pluto")
	t = append(t, "topolini")
	r = append(r, t...) //Si può appendere una slice a un altra slice

	y = append(x[:i], x[i+1:]...) //Si può appendere una slice a un altra rimuovendo un elemento
}