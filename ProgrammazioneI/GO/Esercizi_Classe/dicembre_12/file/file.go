//Autore: Francesco Corrado

package main

import(
	"fmt"
	"os"
)

func main(){
	var buffer []byte
	buffer = make([]byte, 1)
	for{
		n, err:= f.Read(buffer)
		if n == 0{
			break
		}
		if err !=nil{
			fmt.Println("AIUTOOOO")
			break
		}
		fmt.Printf("%c", buffer[0])
	}
	defer f.Close() //Quando finisce l'esecuzione della funzione chiudo il file "defer"
	var b byte[]
	b = []byte("Francesco Corrado")
	n, err:=f.Write(b)
	if err!=nil{
		
	}
}
