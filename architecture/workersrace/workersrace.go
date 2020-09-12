package main

import (
	"context"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

var (
	mirrors = []string{"mirrors.163.com", "mirror.aarnet.edu.au", "cygwin.mirror.rafal.ca", "mirror.easyname.at", "www.mirrorservice.org"}
)

// firstToFinish
// given mirrored websites, access concurrently return first to finish or error
func firstToFinish(ctx context.Context, urls []string, cResult chan string, cError chan error) {

	for i, v := range urls {
		go func (id int, url string, cResult chan string) {
			r, err := http.Get(url)
			if err != nil {
				log.WithFields(log.Fields{"err":err, "id":i}).Error()
				cError <- err
				return
			}
			defer r.Body.Close()

			cResult <- url
		} (i, v, cResult)
	}

}

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 3)

	cResult := make(chan string, len(mirrors))
	cError := make(chan error, len(mirrors))

	firstToFinish(ctx, mirrors, cResult, cError)

	outer:
	for i:=0;i<len(mirrors);i++ {
		select{
		case url := <-cResult:
			log.WithFields(log.Fields{"url":url}).Info("received response")
			cancel()
			break outer
		case err := <-cError:
			log.WithFields(log.Fields{"err":err}).Info("received error")
		case <- ctx.Done():
			log.Errorf("context timeout")
			cancel()
            break outer
		}
	}

}
