package main

import (
	"context"
	log "github.com/sirupsen/logrus"
	"math/rand"
	"sync"
	"time"
)

// makeOrBreak
// demonstrates timeout with context
func makeOrBreak(ctx context.Context, cRes chan int, cErr chan error) {
	t := time.NewTimer(time.Second * time.Duration(rand.Intn(3)))

	select {
	case <-t.C:
		log.Info("completed task")
		cRes <- 123

	case <-ctx.Done():
		log.Error("timeout")
		cErr <- ctx.Err()
	}
}

func main() {

	wg := sync.WaitGroup{}
	cRes := make(chan int)
	cError := make(chan error)
	ctx, _ := context.WithTimeout(context.Background(), time.Second*2)

	go makeOrBreak(ctx, cRes, cError)

	wg.Add(1)
	go func() {
		defer wg.Done()
		select {
		case i := <-cRes:
			log.WithFields(log.Fields{"i": i}).Info("received resp")
			return
		case err := <- cError:
			log.WithFields(log.Fields{"err":err}).Error("received error")
			return

		}
	}()

	wg.Wait()
}

