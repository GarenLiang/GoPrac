package main

import "fmt"
func main() {
  const goldenRatio float64 = 1.666
  fmt.Println(goldenRatio)

  const (
    First = iota + 1
    Second
    Third
  )
  fmt.Println(First, Second, Third)
}
