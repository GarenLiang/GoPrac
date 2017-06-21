package main

import "fmt"
func main(){
  command := "new file"

  switch command {
  case "create":
    fmt.Println("creating...")
  case "open":
    fmt.Println("Opening...")
  case "new file":
    fmt.Println("new file...")
  default:
    fmt.Println("Unrecognised")
  }
}
