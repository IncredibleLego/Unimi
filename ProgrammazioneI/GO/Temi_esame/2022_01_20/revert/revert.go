package main

import(
	"fmt"
	"os"
	"bufio"
)

func main(){
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan(){
		parola := scanner.Text()
		if parola == "stop"{
			break
		}
		if len(parola) % 2 == 0{
			var str string
			for i:= len(parola)-1 ; i >= 0; i--{
				str += string(parola[i])
			}
			fmt.Println(str)
		}
	}
}