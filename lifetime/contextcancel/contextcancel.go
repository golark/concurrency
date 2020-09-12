package main

import (
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"math/rand"
	"sync"
	"time"
)

func makeOrBreak(ctx context.Context, cRes chan int, cErr chan error, id int) {
	defer log.WithFields(log.Fields{"id": id}).Trace("exiting")

	t := time.NewTimer(time.Millisecond * time.Duration(rand.Intn(1000)))

	select {
	case <-t.C:
		log.WithFields(log.Fields{"id": id}).Info("completed task")
		cRes <- 123

	case <-ctx.Done():
		log.WithFields(log.Fields{"id": id}).Error("cancelled")
		cErr <- ctx.Err()
	}
}

func main() {

	const SimultaneousRequests = 10

	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*100)
	cRes := make(chan int, SimultaneousRequests)
	cErr := make(chan error, SimultaneousRequests)

	for i := 0; i < SimultaneousRequests; i++ {
		go makeOrBreak(ctx, cRes, cErr, i)
	}

	wg := sync.WaitGroup{}
	wg.Add(1)
	// receive first response and cancel
	go func() {
		defer wg.Done()
		for {
			select {
			case res := <-cRes:
				cancel() // cancel the rest of the routines
				log.WithFields(log.Fields{"res": res}).Info("received first response")
				return
			case err := <-cErr:
				log.WithFields(log.Fields{"err": err}).Error("received error")
			case <-ctx.Done():
				cancel()
				return
			}
		}
	}()

	wg.Wait()

	time.Sleep(time.Second * 1 )

}
