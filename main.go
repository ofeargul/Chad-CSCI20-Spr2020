// Programmer name: __________
// Date completed:  __________
// Description: ___________________________

package main

import (
    "fmt"
    "math/rand"
    "time"
) //adding the ability to do random numbers

func main() {
    //create two variables - one for the computer and one for the user
    s1 := rand.NewSource(time.Now().UnixNano())
    r1 := rand.New(s1)
    var user int
    var cpu = r1.Intn(2)
    //use a random integer value representing the computer's choice in a game of Rock, Scissors, Paper. 0=rock, 1=scissors, 2=paper
    if cpu == 0 {
      fmt.Println("The computer choose Rock")
    }
    if cpu == 1 {
      fmt.Println("The computer choose Scissors")
    }
    if cpu == 2 {
      fmt.Println("The computer choose paper")
    }
    //prompt the user for an integer value representing the player's choice
    fmt.Println("Please enter 0 for Rock 1 for Scissors and 2 for Paper")
    fmt.Scan(&user)
    if user == 0 {
      fmt.Println(" You Rock")
    }
    if user == 1 {
      fmt.Println(" You Scissors")
    }
    if user == 2 {
      fmt.Println(" You paper")
    }
    
    
    
    //Print out the values using the words rock, scissors, paper.  ie. "Computer chose rock and player chose paper"
    //You will need to use decisions for this
}