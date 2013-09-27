/*

By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can
see that the 6th prime is 13.

What is the 10,001st prime number?

Answer: 104743

*/
package main

import "fmt"

func main() {
  ch := make(chan int)
  go generate(ch)

  prime := 0

  // concurrent prime sieve. not sure yet how this works...
  for i := 0; i < 10001; i++ {
    prime = <-ch

    ch1 := make(chan int)
    go filter(ch, ch1, prime)

    ch = ch1
  }

  fmt.Println(prime)
}

// generate 1, 2, 3... to ch
func generate(ch chan<- int) {
  for i := 2; ; i++ {
    ch <- i
  }
}

// pass from in to out values not divisable by prime
func filter(in <-chan int, out chan<- int, prime int) {
  for {
    i := <-in

    if i % prime != 0 {
      out <- i
    }
  }
}
