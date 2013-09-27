/*

The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.

Find the sum of all the primes below two million.

Answer: 142913828922

*/
package main

import(
  "fmt"
  "./lib"
)

func main() {
  sum := 0

  ch := make(chan int)
  go lib.Primes(ch)

  for {
    prime := <-ch

    if prime > 2000000 {
      break
    }

    sum += prime
  }

  fmt.Println(sum)
}
