package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s string
	s = "pa🥰lo"
	fmt.Println(s)
	for i, c := range s {
		fmt.Printf("%3d\t%c\t%8d\t%08s\n", i, c, c, strconv.FormatInt(int64(c), 2))
	}
}
