// Programmer name: Chad Farrell
// Date completed: 2.6.2020
// Description: names and ages

package main

import "fmt" 

func main() {
  //stating the variables
  var tankSize float32 
  var milesTravled float32
  //getting input for the variables
  fmt.Println("How large is the gas tank?")
  fmt.Scanln(&tankSize)
  fmt.Println("How many miles travled?")
  fmt.Scanln(&milesTravled)
  // outputing MPG of trip
  fmt.Println(tankSize / milesTravled)





}
 