package main

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"time"
)

func getUserRequest(c chan<- string) {

	for {
		s := bufio.NewScanner(os.Stdin)
		s.Scan()

		// scan until error occurs
		if err := s.Err(); err != nil {
			log.WithFields(log.Fields{"err": err}).Error("error occurred during scan, bailing")
			return
		}

		// receive text and send over channel
		t := s.Text()
		c <- t
	}
}

// computeHash
func computeHash(cRequest <-chan string, cResult chan<- string) {

	for req := range cRequest {
		// compute hash and send over
		hash := sha256.Sum256([]byte(req))
		cResult <- string(hash[:])
	}

	log.Info("computeHash exiting")
}

// simply notify the user
func notifyUser(cMsg <-chan string) {

	for msg := range cMsg {
		fmt.Println(msg)
	}

	log.Info("notifyUser exiting")
}

func main() {

	cRequest := make(chan string, 1)
	cResult := make(chan string, 1)

	go getUserRequest(cRequest)
	go computeHash(cRequest, cResult)
	go notifyUser(cResult)

	time.Sleep(time.Minute * 5)
}

