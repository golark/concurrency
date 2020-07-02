package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"sync"
	"time"
)

var (
	NUMGOROUTINES = 5
)

func main() {

	wg := sync.WaitGroup{}
	cExit := make(chan struct{})

	cSink := make(chan int)
	cTop := cSink
	for i := 0; i < NUMGOROUTINES; i++ {
		wg.Add(1)

		//cSource, cSink := cSink, make(chan int)
		cSource := cSink
		cSink = make(chan int)
		fmt.Printf("source[%v] -> sink[%v]\n", cSource, cSink)

		// spin go routine
		go func(id int, wg *sync.WaitGroup, cExit <-chan struct{}, cSource <-chan int, cSink chan<- int) {
			defer func() {
				log.WithFields(log.Fields{"id:": id}).Info("exiting")
				wg.Done()
			}()

			for {
				select {
				case i := <-cSource:
					log.WithFields(log.Fields{"i:": i, "id:": id}).Info("source -> sink")
					cSink <- i + 1
				case <-cExit:
					return
				}
			}

		}(i, &wg, cExit, cSource, cSink)
	}

	log.Info("sending to top Sink")
	//close(cExit)
	cTop <- 0

	res := <-cSink
	log.WithFields(log.Fields{"res:": res}).Info("final sink")

	// exit all go routines
	close(cExit)
	time.Sleep(time.Second * 1)


	wg.Wait()

}
