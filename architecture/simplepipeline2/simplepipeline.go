package main

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"time"
)

// signalHandler
// receive signals over channel
// exit upon interrupt and kill otherwise continue as normal
func signalHandler(c <- chan os.Signal, cExitReq chan <- struct{}) {

	for sig := range c {
		if sig == os.Interrupt || sig == os.Kill {
			log.WithFields(log.Fields{"sig":sig}).Info("received signal, sending exit trigger")
			cExitReq <- struct{}{}
			return
		} else {
			log.WithFields(log.Fields{"sig":sig}).Info("received signal, continue operation")
		}
	}
}


// evaluates exit request and acts upon
func checkExit(cExitReq chan struct{}, cExit chan struct{}) {

	for {
		if _, ok := <- cExitReq; ok {
			close(cExit)
			return
		}
	}
}

func getUserRequest(c chan<- string, cExitReq chan struct{}) {
	defer func() {
		log.Info("getUserRequest exited")
	}()

	for {
		s := bufio.NewScanner(os.Stdin)
		s.Scan()

		// scan until error occurs
		if err := s.Err(); err != nil {
			log.WithFields(log.Fields{"err": err}).Error("error occurred during scan, bailing")
			cExitReq <- struct{}{}
			return
		}

		// receive text and send over channel
		t := s.Text()
		c <- t
	}
}

// computeHash
func computeHash(cRequest <-chan string, cResult chan<- string, cExit chan struct{}) {
	defer log.Info("computeHash exited")

	for {
		select {
		case req := <-cRequest:
			// compute hash and send over
			hash := sha256.Sum256([]byte(req))
			cResult <- string(hash[:])
		case <-cExit:
			return
		}
	}

}

// simply notify the user
func notifyUser(cMsg <-chan string, cExit chan struct{}) {
	defer log.Info("notifyUser exited")

	for {
		select {
		case msg := <-cMsg:
			fmt.Println(msg)
		case <-cExit:
			return
		}
	}
}

func main() {

	cRequest := make(chan string, 1)
	cResult := make(chan string, 1)
	cExitReq := make(chan struct{}, 1)
	cExit := make(chan struct{}, 1)

	cSignal := make(chan os.Signal, 1)
	signal.Notify(cSignal)

	go signalHandler(cSignal, cExitReq)
	go checkExit(cExitReq, cExit)

	go getUserRequest(cRequest, cExitReq)
	go computeHash(cRequest, cResult, cExit)
	go notifyUser(cResult, cExit)

	<- cExit // wait for exit signal
	time.Sleep(time.Second * 1)
}
