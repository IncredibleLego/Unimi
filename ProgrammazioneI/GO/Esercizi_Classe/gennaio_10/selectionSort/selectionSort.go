//Autore: Francesco Corrado

func selectionSort (a []int){
	n := len(a)
	for i:=0; i < n; i++{
		minindex = i
		for j = i+1; j < n; j++{
			if a[i] < a[minindex]{
				minindex = j
			}
		}
		a[i], a[minindex] = a[minindex], a[i]
	}
}