/*

The sum of the squares of the first ten natural numbers is,

  1^2 + 2^2 + ... + 10^2 = 385

The square of the sum of the first ten natural numbers is,

  (1 + 2 + ... + 10)2 = 55^2 = 3025

Hence the difference between the sum of the squares of the first ten
natural numbers and the square of the sum is 3025 − 385 = 2640.

Find the difference between the sum of the squares of the first one
hundred natural numbers and the square of the sum.

Answer: 25164150

*/
package main

import "fmt"

func main() {
  s1 := sum_of_squares(100)
  s2 := square_of_sum(100)

  fmt.Println(s2 - s1)
}

func sum_of_squares(n int) int {
  s := 0

  for i := 1; i <= n; i++ {
    s += i*i
  }

  return s
}

func square_of_sum(n int) int {
  s := 0

  for i := 1; i <= n; i++ {
    s += i
  }

  return s*s
}
