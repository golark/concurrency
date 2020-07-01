package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"math/rand"
	"sync"
	"time"
)

func dBQuery() <-chan int {
	ch := make(chan int)
	go func() {
		time.Sleep(time.Second * time.Duration(rand.Intn(10)))
		ch <- 1
	}()

	return ch
}

func main() {

	rand.Seed(time.Now().Unix())
	wg := sync.WaitGroup{}

	// single db query closure with timeout
	// do not use default case in the select statement
	fmt.Printf("single db query closure\n")
	wg.Add(1)
	go func() {
		defer wg.Done()
		select {
		case <- dBQuery():
			log.Info("received db response in time")
		case <- time.After(time.Second * time.Duration(rand.Intn(10))):
			log.Error("timeout")
		}
	} ()

	wg.Wait()
}

