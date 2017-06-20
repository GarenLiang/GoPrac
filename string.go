package main

import "fmt"

func main() {
  p1 := "Every thing will be ok"
  p2 := "day day up"

  fmt.Println(p1+p2)

  aslice := p1[1:3]
  fmt.Println(aslice)

  character := p1[0]
  fmt.Println(string(character))
}
