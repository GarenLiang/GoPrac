package main

import "fmt"
func main(){
  doSomething()
  add(2, 3, 5)
}

func doSomething() {
  fmt.Println("Performing the doSomething function")
}
func add(a int, b int, c int) {
  fmt.Println(a + b + c)
}
//-------------------
package main
import "fmt"
func main() {
  result := superComplexCalculation()
  fmt.Println(result)
}
func superComplexCalculation() string {
  str := "some string"


  return str
}
