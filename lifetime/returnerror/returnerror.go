package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"golang.org/x/sync/errgroup"
)

// fetchConcurrently
// start a fetch operation with go routines working concurrently
// n is the number of go routines
func fetchConcurrently(ctx context.Context, n int) error {

	errs, ctx := errgroup.WithContext(ctx)

	for i := 0; i < n; i++ {
		errs.Go(func() error {
			defer fmt.Printf("exit\n")

			// wait for a random time and error
			tWork := time.NewTimer(time.Millisecond * time.Duration(rand.Intn(20)))
			tErr := time.NewTimer(time.Millisecond * time.Duration(rand.Intn(100)))
			for {
				select {
				case <-tErr.C:
					tErr.Stop()
					fmt.Printf("errorred: ")
					return fmt.Errorf("exited go routine with error")
				case <-tWork.C:
					fmt.Printf("work completed: ")
					tWork.Stop()
					return nil
				}
			}
		})
	}

	// Wait for completion and return the first error (if any)
	return errs.Wait()
}

func main() {

	err := fetchConcurrently(context.Background(), 5)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	time.Sleep(time.Second * 1)
}
