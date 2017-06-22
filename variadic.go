package main
import "fmt"
func main(){
  r := add(6, 8, 10)
  fmt.Println(r)
}

func add(nums ...int) int {
   var total int

   for _, n := range nums {
    total += n
   }
   return total
}
