// Programmer name: Chad Farrell
// Date completed: 2.13.2020
// Description: This program does math to find the total cost for a yogurt shop.

package main

import "fmt" 

func main() {
//Creating the variables
  var Oz, coupon, tax, totalWithout, totalWith, tip, toppings2 float32
  var toppings int
  var couponQ int

//asking the user for ounces of yogurt
  fmt.Println("How many ounces?")
  fmt.Scanln(&Oz)
  totalWithout = (Oz) * .17
  
//getting the toppings prics
  fmt.Println("how many toppings")
  fmt.Scanln(&toppings)
  toppings2 = float32(toppings) * .5

//getting coupon info if any
  fmt.Println("does the coustomer have a coupon? Enter 1 for yes and 2 for no")
  fmt.Scanln(&couponQ)
  if couponQ == 1 {
    fmt.Println("please enter coupon amount in decimals")
    fmt.Scanln(&coupon)
    totalWith = (coupon * totalWithout)
     }

//adding total and topping + the weight of the cup
  totalWith = totalWith + toppings2 + .00425

//finding the tax  
  tax = totalWith * .725

//finding the tip amount
  fmt.Println("Please enter tip amount")
  fmt.Scanln(&tip)
 
//final totals
totalWith = (totalWith + tax) + tip

fmt.Println("The price of just yogurt is",totalWithout)
fmt.Println("The total price is", totalWith)







}
 