package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"sync"
	"time"
)

func selectLoop(c1 chan string, c2 chan string, c3 chan string, wg *sync.WaitGroup) {
	defer func() {
		log.Info("exiting selectLoop")
		wg.Done()
	}()

	fmt.Printf("\n")

	timeout := time.After(time.Second * 20)
	for {
		select {
		case m:= <-c1:
			fmt.Printf("%v",m)
		case m:= <-c2:
			fmt.Printf("%v",m)
		case m:= <-c3:
			fmt.Printf("%v",m)
		case <-timeout:
			log.Info("timeout")
			return
		default:
			time.Sleep(time.Second * 1)
			//fmt.Printf(".")
		}
	}
}

func trafficGenerator(c1 chan string, c2 chan string, c3 chan string) {

	// generate some random traffic
	t1 := time.Tick(time.Second * 2)
	t2 := time.Tick(time.Second * 4)

	tickFreq := time.Duration(time.Second * 4)
	t := time.NewTimer(tickFreq)

	for {
		select {
		case <-t1:
			c1 <- "|\n"
		case <-t2:
			c2 <- " |\n"
		case <-t.C:
			tickFreq = tickFreq - tickFreq/4
			t.Reset(tickFreq)
			c3 <- "  |\n"
		}
	}

	t.Stop()

}

func selectPrint(s string) chan struct{} {
	fmt.Printf("selectPrint: %v\n", s)
	c := make(chan struct{})

	return c
}

func forSelectExamples(wg * sync.WaitGroup) {
	defer wg.Done()

	// loop select without default
	go func() {
		for {
			select {
			case <-selectPrint("loop select without default"):
			}
		}
	}()

	// loop select with default
	// Do not call functions and return new channel everytime when using for select with default option!!
	// time.After generates a new Timer object everytime it is called, and returns back a new channel everytime
	// hence do not use time.After in for select with default option
	go func() {
		for {
			select {
			case <-selectPrint("loop select WITH default"):
			default:
				time.Sleep(time.Second * 1)
			}
		}
	}()

	time.Sleep(10 * time.Second)
}

func main() {

	wg := sync.WaitGroup{}
	c1 := make(chan string, 10)
	c2 := make(chan string, 10)
	c3 := make(chan string, 10)

	wg.Add(1)
	go selectLoop(c1, c2, c3, &wg)
	go trafficGenerator(c1, c2, c3)
	// go forSelectExamples(&wg)
	wg.Wait()
}
