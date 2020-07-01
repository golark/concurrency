package main

import (
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
	} ()

	return ch
}

func internalTimeout(wg *sync.WaitGroup) {
	defer func() {
		log.Info("exiting internalTimeout")
		wg.Done()
	} ()

	select {
		case <-time.After(time.Second * 5):
			log.Info("timeout")
			return
		case i :=  <- dBQuery():
			log.WithFields(log.Fields{"val:":i}).Info("received from dB")
	}

}

func main() {


	rand.Seed(time.Now().Unix())

	wg := sync.WaitGroup{}

	wg.Add(1)
	go internalTimeout(&wg)

	wg.Wait()

}
