package main

import "time"

func chain(in chan int) chan int {
	out := make(chan int)
	go func() {
		for v := range in {
			out <- 1 + v
		}
		close(out)
		println("going to sleep 1s in chain")
		time.Sleep(1 * time.Second)
	}()
	return out
}
func main() {
	in := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			in <- i
		}
		close(in)
	}()
	out := chain(chain(chain(in)))
	for v := range out {
		println(v)
		time.Sleep(3 * time.Second)
	}
}
