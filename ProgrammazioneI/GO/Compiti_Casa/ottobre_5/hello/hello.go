package main
import "fmt"
func main(){
	var n, i int
	fmt.Print("Quante volte, padre?")
	fmt.Scan(&n)
	for i=0; i<n; i++ {
		fmt.Println("Ciao")
	}
}