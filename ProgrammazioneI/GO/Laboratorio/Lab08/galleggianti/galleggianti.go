//Autore: Francesco Corrado

package main

import(
	"fmt"
	"bufio"
	"os"
	"strconv"
)

func ordinaMatrice(ma [][]string, ri int, co int) [][]string{
	var asterischi[]string
	var caratteri[]string
	for y:=0; y<co; y++{
		for x:=0; x<ri; x++{
			if ma[x][y] == "*"{
				asterischi = append(asterischi, ma[x][y])
			}else{
				caratteri = append(caratteri, ma[x][y])
			}
		}
		for k:=0; k<len(asterischi); k++{
			ma[k][y] = asterischi[k]
		}
		l := 0
		for k:=len(asterischi); k<ri; k++{
			ma[k][y] = caratteri[l]
			l++
		}
		asterischi = nil
		caratteri = nil
	}
	return ma
}

func stampaMatrice(matrix [][]string, righe int) bool{
	for k:=0; k<righe; k++{
		fmt.Println(matrix[k])
	}
	return true
}

func main(){
	var r, c, x, y int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	r, _ = strconv.Atoi(scanner.Text())
	scanner.Scan()
	c, _ = strconv.Atoi(scanner.Text())
	var m [][]string
	m = make([][]string, r)
	for z:=0; z<r; z++{
		m[z] = make([]string, c)
	}
	for scanner.Scan(){
		carattere := scanner.Text() //Bisognerebbe utilizzare strings.Split(s, " ")
		m[x][y] = carattere
		y++
		if y == c{
			y = 0
			x++
		}
	}
	fmt.Println("m iniziale:")
	stampaMatrice(m, r)
	m = ordinaMatrice(m, r, c)
	fmt.Println("m ordinata:")
	stampaMatrice(m, r)
}