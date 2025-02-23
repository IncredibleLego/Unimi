//Autore: Francesco Corrado

package main

import(
	"fmt"
	"time"
)

type Clock struct {
	hour, min, sec int
}

func countdown (ho int, mi int, se int) bool{ //Crea problemi se orario.sec è 0 all'inizio
	var orario Clock
	orario.hour = ho
	orario.min = mi
	orario.sec = se
	time.Sleep(time.Duration(1) * time.Second)
	fmt.Println("{", orario.hour, orario.min, orario.sec, "}")
	orario.sec -= 1
	if orario.hour == 0 && orario.min == 0 && orario.sec == 0{
		fmt.Println("{", orario.hour, orario.min, orario.sec, "}")
		return true
	}
	if orario.sec == 0{
		time.Sleep(time.Duration(1) * time.Second)
		fmt.Println("{", orario.hour, orario.min, orario.sec, "}")
		orario.hour, orario.min, orario.sec = downMin(orario.hour, orario.min, orario.sec)
	}
	countdown(orario.hour, orario.min, orario.sec)
	return false //Can't happen
}

func downMin (ho int, mi int, se int) (int, int, int){
	var orario Clock
	orario.hour = ho
	orario.min = mi
	orario.sec = se
	if orario.min == 0{
		orario.hour, orario.min, orario.sec = downHour(orario.hour, orario.min, orario.sec)
	}else{
		orario.min -= 1
		orario.sec = 59
	}
	return orario.hour, orario.min, orario.sec
}

func downHour (ho int, mi int, se int) (int, int, int){
	var orario Clock
	orario.hour = ho
	orario.min = mi
	orario.sec = se
	orario.hour -= 1
	orario.min = 59
	orario.sec = 59
	return orario.hour, orario.min, orario.sec
}


func main(){
	var h, m, s int
	fmt.Print("time (hh mm ss):")
	fmt.Scan(&h, &m, &s)
	countdown(h, m, s)
}