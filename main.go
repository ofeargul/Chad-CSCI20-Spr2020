// Programmer name: __________
// Date completed:  __________
// Description: ___________________________

package main

import "fmt" 
var tax float32 
var status int
func main() {
	fmt.Println("TAX TIME")
  fmt.Println("Please Enter what best applies to you")
  fmt.Println("1 for Single, 2 for Married, 3 for Head of household and 4 for Divorced")
  fmt.Scan(&status)
  fmt.Println("Now please enter your income")
  fmt.Scanln(&tax)
  //single IF
  if status == 1 {
    if tax < 9700{
      fmt.Println("your in the 10% bracket and you owe", tax * .10)
    }
    if 9701 < tax && tax < 39475{
      fmt.Println("your in the 12% bracket and you owe",tax * .12)
    }
    if 39476 < tax && tax < 84200{
      fmt.Println("your in the 22% bracket and you owe",
      tax * .22)
    } 
    if 84201 < tax && tax < 160725{
      fmt.Println("your in the 24% bracket and you owe", tax * .24)
    }
    if 160726 < tax && tax < 204100{
      fmt.Println("your in the 32% bracket and you owe", tax * .32)
    }
    if 204101 < tax && tax < 510300{
      fmt.Println("your in the 35% bracket and you owe", tax * .35 )
    }
    if 510301 < tax{
      fmt.Println("your in the 37% tax bracket and owe", tax * .37)
    }
  }
  //married if
  
  if status == 2 {
    if tax < 19400{
      fmt.Println("your in the 10% bracket and you owe", tax * .10)
    }
    if 19401 < tax && tax < 78950{
      fmt.Println("your in the 12% bracket and you owe", tax * .12)
    }
    if 78951 < tax && tax < 168400{
      fmt.Println("your in the 22% bracket and you owe", tax * .22)
    } 
    if 168401 < tax && tax < 321450{
      fmt.Println("your in the 24% bracket and you owe", tax * .24)
    }
    if 321451 < tax && tax < 408200{
      fmt.Println("your in the 32% bracket and you owe", tax * .32)
    }
    if 408201 < tax && tax < 612350{
      fmt.Println("your in the 35% bracket and you owe", tax * .35 )
    }
    if 612351 < tax{
      fmt.Println("your in the 37% tax bracket and owe", tax * .37)
    }
  }
  if status == 3 {
    if tax < 13850{
      fmt.Println("your in the 10% bracket and you owe", tax * .10)
    }
    if 13851 < tax && tax < 52850{
      fmt.Println("your in the 12% bracket and you owe", tax * .12)
    }
    if 52851 < tax && tax < 84200{
      fmt.Println("your in the 22% bracket and you owe", tax * .22)
    } 
    if 84201 < tax && tax < 160700{
      fmt.Println("your in the 24% bracket and you owe", tax * .24)
    }
    if 160701 < tax && tax < 204100{
      fmt.Println("your in the 32% bracket and you owe", tax * .32)
    }
    if 204101 < tax && tax < 510300{
      fmt.Println("your in the 35% bracket and you owe", tax * .35 )
    }
    if 510301 < tax{
      fmt.Println("your in the 37% tax bracket and owe", tax * .37)
    }
  }
  if status == 4 {
    if tax < 9700{
      fmt.Println("your in the 10% bracket and you owe", tax * .10)
    }
    if 9701 < tax && tax < 39475{
      fmt.Println("your in the 12% bracket and you owe", tax * .12)
    }
    if 39476 < tax && tax < 84200{
      fmt.Println("your in the 22% bracket and you owe", tax * .22)
    } 
    if 84201 < tax && tax < 160725{
      fmt.Println("your in the 24% bracket and you owe", tax * .24)
    }
    if 160726 < tax && tax < 204100{
      fmt.Println("your in the 32% bracket and you owe", tax * .32)
    }
    if 204101 < tax && tax < 306175{
      fmt.Println("your in the 35% bracket and you owe", tax * .35 )
    }
    if 206176 < tax{
      fmt.Println("your in the 37% tax bracket and owe", tax * .37)
    }
  }
}