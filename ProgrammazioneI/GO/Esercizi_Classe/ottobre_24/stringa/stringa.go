package main

import "fmt"
import "strconv"

func main() {
  var s string
  s = "pa🥰lo"
  for i:=0; i < len(s); i++ {
    c := rune(s[i])
    fmt.Printf("%c\t%8d\t%08s\n", c, c, strconv.FormatInt(int64(c), 2))
  }
}
//🥰