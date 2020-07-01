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

func externalTimeout(timeout chan struct{}, wg *sync.WaitGroup) {
	defer func() {
		log.Info("exiting externalTimeout")
		wg.Done()
	} ()

	select {
		case <-timeout:
			log.Info("timeout")
			return
		case i :=  <- dBQuery():
			log.WithFields(log.Fields{"val:":i}).Info("received from dB")
	}

}

func main() {


	rand.Seed(time.Now().Unix())

	wg := sync.WaitGroup{}
	timeout := make(chan struct{})

	wg.Add(1)
	go externalTimeout(timeout, &wg)

	go func() {
		time.Sleep(time.Second * time.Duration(rand.Intn(10)))
		timeout <- struct{}{}
	} ()

	wg.Wait()

}
