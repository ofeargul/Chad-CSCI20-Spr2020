// Programmer name: Chad Farrell
// Date completed: 2.6.2020
// Description: names and ages

package main

import "fmt" 

func main() {
  //stating the variables
  var NumOne int
  var NumTwo int
  //getting input for the variables
  fmt.Println("Please enter a whole number")
  fmt.Scanln(&NumOne)
  fmt.Println("Please enter another whole number")
  fmt.Scanln(&NumTwo)
  // doing math....
  fmt.Println("Addition", NumOne + NumTwo )
  fmt.Println("Subtraction", NumOne - NumTwo )
  fmt.Println("Multiplication", NumOne * NumTwo )
  fmt.Println("Divison", NumOne / NumTwo )
  fmt.Println("remainder", NumOne % NumTwo )







}
 