/*

The prime factors of 13195 are 5, 7, 13 and 29.

What is the largest prime factor of the number 600851475143 ?

Answer: 6857

*/
package main

import "fmt"

func main() {
  n := 0

  for _, n = range factors(600851475143) { }

  fmt.Println(n)
}

func factors(n int) ([]int) {
  z := 2
  primes := make([]int, 0)

  for z*z <= n {
    if n % z == 0 {
      primes = append(primes, z)
      n = n / z
    } else {
      z++
    }
  }

  if n > 1 {
    primes = append(primes, n)
  }

  return primes
}
