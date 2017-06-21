package main
import "fmt"

func main(){
  condition := 5 < 2

  if condition {
    fmt.Println("5 is bigger")
  } else {
    fmt.Println("This is false")
  }
}
