/*

A Pythagorean triplet is a set of three natural numbers, a < b < c, for
which,

  a^2 + b^2 = c^2

For example, 32 + 42 = 9 + 16 = 25 = 52.

There exists exactly one Pythagorean triplet for which a + b + c = 1000.

Find the product abc.

Answer: 31875000

*/
package main

import "fmt"

func main() {
  for a := 1; ; a++ {
    for b := a+1; b < 998; b++ {
      c := 1000 - b - a

      if pythagorean(a, b, c) {
        fmt.Println(a*b*c)
        return
      }
    }
  }
}

func pythagorean(a int, b int, c int) bool {
  return a*a + b*b == c*c
}
