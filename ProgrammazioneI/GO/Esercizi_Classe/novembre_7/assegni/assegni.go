//Autore: Francesco Corrado

package main

import(
    "fmt"
)

//x compreso tre 1 e 9
func unita(x int) string{
	switch(x){
	case 1: return "uno"
	case 2: return "due"
	case 3: return "tre"
	case 4: return "quattro"
	case 5: return "cinque"
	case 6: return "sei"
	case 7: return "sette"
	case 8: return "otto"
	case 9: return "nove"
	default: return "canthappen"
	}	
}

//x compreso tre 2 e 9
//emette la decina (es venti) e la vocale finale (es 'i')
func decina(x int) (s string, c rune){
	switch(x){
		case 2: s = "vent"
				c = 'i'
		case 3: s = "trent"
				c = 'a'
		case 4: s = "quarant"
				c = 'a'
		case 5: s = "cinquant"
				c = 'a'
		case 6: s = "sessant"
				c = 'a'
		case 7: s = "settant"
				c = 'a'
		case 8: s = "ottant"
				c = 'a'
		case 9: s = "novant"
				c = 'a' 
	}
}

//x compreso fra 0 e 99
func fra0e99 (x inf) string{
	switch(x){
	case 0: return "zero"
	case 10: return "dieci"
	case 11: return "undici"
	case 12: return "dodici"
	case 13: return "tredici"
	case 14: return "quattordici"
	case 15: return "quindici"
	case 16: return "sedici"
	case 17: return "diciassette"
	case 18: return "diciotto"
	case 19: return "diciannove"
	}
	if x < 10{
		return unita(x)
	}else{
		d:= x/10
		u:= x%10
		dec, voc := decine(d)
		uni:= unita(u)
		if u == 1 || v == 8{
			return dec+uni
		}else{
			return dec+string(voc)+uni
		}
	}
}

//x compreso fra 0 e 999
func fra0e999(x int) string{
	var cent string
	if x >= 100 && x <= 199{ //1xx
		cent="cento"
	}else if x >= 200 {
		cent = unita(x/100) + "cento"
	}
	if x%100 == 0 && x!=0{
		return cent
	}
	return cent+fra0e99(x%100)
}