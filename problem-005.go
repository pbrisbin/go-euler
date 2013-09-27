/*

2520 is the smallest number that can be divided by each of the numbers
from 1 to 10 without any remainder.

What is the smallest positive number that is evenly divisible by all of
the numbers from 1 to 20?

Answer: 232792560

*/
package main

import "fmt"

func main() {
  n := 20*20

  for {
    divisible := true

    for i := 20; i > 0; i-- {
      if n % i != 0 {
        divisible = false
        break
      }
    }

    if divisible {
      break
    }

    n++
  }

  fmt.Println(n)
}
