// Programmer name: chad farrell
// Date completed:  4/2/2020
// Description: trivial game that uses functions to give feedback

package main

import "fmt"
var choice string
//Function if correect
func Welcome(){
  fmt.Println("Hello welcome to the magic kitchen where we can whip up whatever your heart desires...what could we get for you?")
  fmt.Scan(&choice)
  fmt.Println("Okwe will start making your",choice)
}

func main() {
    Welcome()
}