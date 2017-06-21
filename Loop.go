package main
import "fmt"
func main(){
  for i :=0; i<= 10; i++ {
    fmt.Println(i)
  }
  myName := "Julian"
  for _, char := range myName {
    fmt.Print(string(char) + " ")
  }
  fmt.Println()
}
//---------------------
package main
import "fmt"
func main(){
  i := 0

  for i<5 {
    fmt.Println("golang")
    i++
  }
}
