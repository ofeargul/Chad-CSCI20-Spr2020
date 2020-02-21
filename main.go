// Programmer name: __________
// Date completed:  __________
// Description: ___________________________

package main

import "fmt" 

func main() {
  n := 1234
  for 0 < n && n > 9{
    var sum, a, b, c int
    //dividing 1234 to get the 1 isolated and passing the 234 on
    a = n / 1000
    //fmt.Println(a)
    sum = n % 1000
    //fmt.Println(sum)
    //dividing 234 to get the 2 isolated and then passing 34 on
    b = sum / 100
    //fmt.Println(b)
    sum = sum % 100
    //fmt.Println(sum)
    //
    c = sum / 10
    //fmt.Println(c)
    sum = sum % 10
    //fmt.Println(sum)
    //
    n = a + b + c + sum

  

  }
 fmt.Println(n)
}