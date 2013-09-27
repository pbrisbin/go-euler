package lib

// generate 2, 3, 5... to out
func Primes(out chan<- int) {
  ch := make(chan int)
  go generate(ch)

  for {
    prime := <-ch

    out <- prime

    ch1 := make(chan int)
    go filter(ch, ch1, prime)

    ch = ch1
  }
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
