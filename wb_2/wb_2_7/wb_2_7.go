package main

import (
	"fmt"
	"sync"
	"time"
)

func or(channels ...<-chan interface{}) <-chan interface{} {
	var wg sync.WaitGroup
	res := make(chan interface{})

	output := func(ch <-chan interface{}) {
		for a := range ch {
			res <- a
		}
		wg.Done()
	}

	wg.Add(len(channels))
	for _, ch := range channels {
		go output(ch)
	}

	go func() {
		wg.Wait()
		close(res)
	}()

	return res
}

func main() {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})

		go func() {
			defer close(c)
			time.Sleep(after)
		}()

		return c
	}

	start := time.Now()

	<-or(
		sig(2*time.Second),
		sig(5*time.Second),
		sig(1*time.Second),
		sig(1*time.Second),
		sig(1*time.Second),
	)

	fmt.Printf("fone after %v\n", time.Since(start))
}
