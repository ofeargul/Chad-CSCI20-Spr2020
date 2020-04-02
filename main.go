// Programmer name: chad farrell
// Date completed:  4/2/2020
// Description: trivial game that uses functions to give feedback

package main

import "fmt"
var anwser int
//Function if correect
func Pass(){
  fmt.Println("Correct good job!!")
}
func Fail(){
  fmt.Println("Oops wrong!!")
}
func main() {
    fmt.Println("Whats 9 + 11 =?")
    fmt.Scanln(&anwser)
    if anwser == 21{
      Pass()
    }else{
      Fail()
    }
}