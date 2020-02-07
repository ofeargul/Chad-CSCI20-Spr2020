// Programmer name: Chad Farrell
// Date completed: 2.6.2020
// Description: names and ages

package main

import "fmt" 

func main() {
    //declare variable for favorite food and store your favorite food.
    var Food = "pizza"


    //declare variables for name and age (make sure they are appropriate data types)
    var age int
    var name string 

    //ask the user to enter their answer for name and age.
    fmt.Println("please enter your name")
    fmt.Scanln(&name)
    fmt.Println("Please enter your age") 
    fmt.Scanln(&age)
    //output a welcome statement using their name
    fmt.Println("hi", name )
    //output a statement that says At their age you enjoyed the favorite food
    fmt.Println("when I was", age, "i liked", Food, "alot")
}
 