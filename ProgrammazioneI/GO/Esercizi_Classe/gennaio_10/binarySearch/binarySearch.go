//Autore: Francesco Corrado

func binarySearch (a []int, x int) int{
	low := 0
	high := len(a)-1
	for low <= high{
		median:=(low+high)/2
		if a[median] == x{
			return median
		}
		if a[median] < x{
			low = median + 1
		}else{
			high = median-1
		}
	}
	return -1
}