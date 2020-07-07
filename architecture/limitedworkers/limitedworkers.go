package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	FIBRANGE = 50
)

func fib(n int) int {

	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}

	return fib(n-1) + fib(n-2)
}

func randomWorkGenerator(numGens int, cWork chan<- int, wg *sync.WaitGroup) {
	defer func() {
		close(cWork)
		wg.Done()
	} ()

	sTime := rand.NewSource(1010)
	sNum  := rand.NewSource(1010)

	rTime := rand.New(sTime)
	rNum  := rand.New(sNum)

	t := time.NewTimer(time.Millisecond * time.Duration(rTime.Intn(100)))


	for numGens > 0 {
		<-t.C
		cWork <- rNum.Intn(FIBRANGE)
		t.Reset(time.Millisecond * time.Duration(rTime.Intn(100)))

		numGens--
	}
}

func uniformWorkGenerator(numGens int, cWork chan<- int, wg *sync.WaitGroup) {
	defer func() {
		close(cWork)
		wg.Done()
	} ()

	fibStart := 30

	for numGens > 0 {
		cWork <- fibStart
		fibStart++
		numGens--
	}
}

func getResults(numRequests int, cRes <-chan [2]int, wg *sync.WaitGroup) {

		defer wg.Done()

		res := make([][2]int, numRequests)
		for i:=0; i<numRequests;i++ {
			if r, ok := <- cRes; ok {
				res[i] = r
			}
		}

		// print the results
		for _,v := range res {
			fmt.Printf("fib(%v) = %v\n", v[0],v[1])
		}
}

func main() {

	numRequests := 10
	cWork := make(chan int, numRequests)
	cRes := make(chan [2]int, numRequests)

	wg := sync.WaitGroup{}

	// step 1- work generation
	wg.Add(1)
	// go randomWorkGenerator(numRequests, cWork, &wg)
	go uniformWorkGenerator(numRequests, cWork, &wg)

	// step 2 - collect the results
	wg.Add(1)
	go getResults(numRequests, cRes, &wg)


	// step 3 - workers
	wgWorker := sync.WaitGroup{}
	start := time.Now()
	cWorkers := make(chan struct{}, 10)
	for i:=0;i<cap(cWorkers);i++ {
		cWorkers <- struct{}{}
	}

	for i := range cWork{
		wgWorker.Add(1)

		// func(i int, cWorker chan struct{}, wg *sync.WaitGroup) {
		go func(i int, cWorker chan struct{}, wg *sync.WaitGroup) {
			defer func() {
				cWorkers <- struct{}{}
				wg.Done()
			} ()
			<- cWorkers
			cRes <- [2]int{i,fib(i)}
		} (i, cWorkers, &wgWorker)
	}
	wgWorker.Wait()
    elapsed := time.Since(start)
    fmt.Printf("time elaped: %s\n", elapsed)

	wg.Wait()
}
