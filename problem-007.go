/*

By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can
see that the 6th prime is 13.

What is the 10,001st prime number?

Answer: 104743

*/
package main

import(
  "fmt"
  "./lib"
)

func main() {
  i := 1

  ch := make(chan int)
  go lib.Primes(ch)

  for {
    prime := <-ch

    if i == 10001 {
      fmt.Println(prime)
      break
    }

    i++
  }
}
