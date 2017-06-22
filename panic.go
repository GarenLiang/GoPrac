package main
import (
  "io/ioutil"
  "fmt"
)
func main(){
  b, err := ioutil.ReadFile("names.txt")
  if err != nil {
    panic(err)
  }
  for _, e := range b {
    fmt.Print(string(e))
  }
}
