//Autore: Francesco Corrado

package main

import(
	"fmt"
	"strings"
)

func main(){
	var s string
	fmt.Print("Inserire un URL: ")
	fmt.Scan(&s)
	x:= setHost(s)
	fmt.Println(x)
}

func setHost (url string) string{
	bb:= strings.Index(url, "//") //Trovo la posizione di inizio di //
	b:= strings.Index(url[bb+2:], "/")+bb+2 //Trovo la posizione della seconda / ma in relazione alla seconda stringa
	return url[bb+2:b] //b e escluso
}