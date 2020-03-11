package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() { 
  player1 := 0
  player2 := 0
  var roll1, roll2 int
  var dice int
  var decsion string
  var turncounter int
  var pot, pot2 int
  s1 := rand.NewSource(time.Now().UnixNano())
  rand := rand.New(s1)
  //intro and getting dice size
  fmt.Println("WELCOME TO PIG")
  fmt.Println("how large of a dice would you like to use?")
  fmt.Scanln(&dice)
  //starting a nested loop
  for player1 < 100 && player2 < 100{
    switch turncounter {
      //switch for player
      case 0:
      roll1 = (rand.Intn(dice)) + 1
      fmt.Println("You rolled a", roll1)
      if roll1 == 1 {
        fmt.Println("you lose your turn and points from this turn")
        turncounter++
      }else{
        pot = pot + roll1
        fmt.Println("Your points for this turn are", pot)
        fmt.Println("do you want to roll agian? Yes or No")
        fmt.Scanln(&decsion)
        if decsion == "No" || decsion == "NO" || decsion == "no"{
          player1 = player1 + pot
          fmt.Println("Your total score is", player1)
          turncounter++
        }
      }
      //switch for computer
      case 1:
      roll1 = (rand.Intn(dice)) + 1
      fmt.Println("The computer rolled a", roll1)
      if roll1 == 1 {
        fmt.Println("The computer lost its turn")
        turncounter--
      }else{
        pot2 = pot2 + roll1
        roll2 =(rand.Intn(dice)) + 1
        if roll2 < 3{
          player2 = pot2 + player2
          fmt.Println("The computers score is", player2)
          turncounter--
        }

      }
    }
  }
  if player1 >= 100{
    fmt.Println("You won your score is!", player1)
  }
  if player2 >= 100{
    fmt.Println("The computer score is", player2)
  }
}
     
    
