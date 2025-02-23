//Autore: Francesco Corrado

package main

import(
	"fmt"
	"os"
	"strconv"
)

type Rettangolo struct{
	base, altezza int
}

func (re Rettangolo) String() string{
	var output string
	if re.base == 0 || re.altezza == 0{
		output = "rettangolo degenere"
	}else{
		for i := 0; i < re.altezza; i++{
			for j := 0; j < re.base; j++{
				output += "."
			}
			output += "\n"
		}
	}
	return output
}

func main(){
	var r Rettangolo
	r.base, _ = strconv.Atoi(os.Args[1])
	r.altezza, _ = strconv.Atoi(os.Args[2])
	fmt.Println(r.String())
}