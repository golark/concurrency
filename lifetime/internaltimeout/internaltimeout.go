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
	log.Info("making db Query")
	go func() {
		// time.Sleep(time.Second * time.Duration(rand.Intn(10)))
		time.Sleep(time.Second * 3)
		log.Info("send dB Query result")
		ch <- 1
	}()

	return ch
}

func internalTimeout(wg *sync.WaitGroup) {
	defer func() {
		log.Info("exiting internalTimeoutV1")
		wg.Done()
	}()

	timeout := time.After(50 * time.Second)

	chandB := dBQuery()

	tCount := 0
	for {
		select {
		case <-timeout:
			log.Info("timeout")
			return
		case i := <-chandB:
			log.WithFields(log.Fields{"val:": i}).Info("received from dB")
			return
		default:
			// idle operation
			time.Sleep(1000 * time.Millisecond)
			tCount++
			fmt.Printf("%vms\n",tCount*100)
		}
	}
}

func main() {

	rand.Seed(time.Now().Unix())
	wg := sync.WaitGroup{}


	wg.Add(1)
	go internalTimeout(&wg)

	wg.Wait()
}

