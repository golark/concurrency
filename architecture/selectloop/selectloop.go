package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"math/rand"
	"sync"
	"time"
)

func selectLoop(c1 chan int, c2 chan struct{}, c3 chan string, wg *sync.WaitGroup) {
	defer func() {
		log.Info("exiting selectLoop")
		wg.Done()
	} ()

	fmt.Printf("\n")
	for {
		select {
		case <-c1:
			log.Info("received from c1")
		case <-c2:
			log.Info("received from c2")
		case <-c3:
			log.Info("received from c3")
		case <-time.After(time.Second * 10):
			log.Info("10 seconds elapsed")
		default:
			time.Sleep(time.Second * 1)
			fmt.Printf(".")
		}
	}
}



func trafficGenerator(c1 chan int, c2 chan struct{}, c3 chan string) {

	// generate some random traffic
	for {
		select {
			case <-time.After(time.Second * time.Duration(rand.Intn(10))):
				c1 <- rand.Intn(10) // send random int
			case <-time.After(time.Second * time.Duration(rand.Intn(10))):
				c2 <- struct{}{}
			case <-time.After(time.Second * time.Duration(rand.Intn(10))):
				c3 <- "hey"
		}
	}

}

func main() {

	wg := sync.WaitGroup{}
	c1 := make(chan int)
	c2 := make(chan struct{})
	c3 := make(chan string)

	wg.Add(1)
	go selectLoop(c1, c2, c3, &wg)
	go trafficGenerator(c1, c2, c3)
	wg.Wait()
}
