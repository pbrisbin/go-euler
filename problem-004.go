/*

A palindromic number reads the same both ways. The largest palindrome
made from the product of two 2-digit numbers is 9009 = 91 × 99.

Find the largest palindrome made from the product of two 3-digit
numbers.

Answer: 906609

*/
package main

import (
  "strconv"
  "fmt"
)

func main() {
  max := 0
  limit := 999 // three-digit

  for i := 1; i <= limit; i++ {
    for j := 1; j <= limit; j++ {
      n := i * j

      if n > max && palindrome(n) {
        max = n
      }
    }
  }

  fmt.Println(max)
}

func palindrome(n int) bool {
  s := strconv.Itoa(n)

  return s == reverse(s)
}

func reverse(s string) string {
  runes := []rune(s)

  for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
    runes[i], runes[j] = runes[j], runes[i]
  }

  return string(runes)
}
