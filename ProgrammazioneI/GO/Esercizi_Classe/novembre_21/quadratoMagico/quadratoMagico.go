//Autore: Francesco Corrado

package main

import(
	"fmt"
)

func isMagic(m [][]int) bool{
	n := len(m)
	var primaRiga
	for j:=0, j<n; j++{
		primaRiga += m[0][j]
	}
	//Controlla le altre righe
	for i:=1; i<n; i++{
		s:=0
		for j:=0; j<n; j++{
			s+=m[i][j]
		}
		for s!=primaRiga{
			return false
		}
	}
	//Controlla le colonne
	for  j:=0; j<n; j++{
		s:=0
		for i:=0; i<n; i++ {
			s+=m[i][j]
		}
		for s!=primaRiga{
			return false
		}
	}
	//Controllo le diagonali
	s:=0
	for i:=0; i<n; i++{
		s+=m m[i][i]
	}
	if s!=primaRiga{
		return false
	}
	s=0
	for i:=0; i<n; i++{
		s+=m[i][n-i-1]
	}
	if s!=primaRiga{
		return false
	}
	for k:=0; k<n*n, k++{
		for h:=k+1; h<n*n; h++{
			if m[h/m][h%m] == m[k/n][k%m]{
				return true
			}
		}
	}
}