package main

import (
	"fmt"
	"os"
	"text/tabwriter"
)

func lenOfChannel() {

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, '.', tabwriter.AlignRight|tabwriter.Debug)
	c := make(chan int, 5)

	fmt.Fprintf(w,"c = make(chan int)\t  len(c): %v\t cap(c): %v\t\n",len(c), cap(c))

	c <- 1
	fmt.Fprintf(w,"c <-1\t  len(c): %v\t cap(c): %v\t\n",len(c), cap(c))

	c <- 1
	fmt.Fprintf(w,"c <-1\t  len(c): %v\t cap(c): %v\t\n",len(c), cap(c))

	c <- 1
	fmt.Fprintf(w,"c <-1\t  len(c): %v\t cap(c): %v\t\n",len(c), cap(c))

	w.Flush()

}


// makingChannels
// buffered channels cap >= 1
// unbuffered channels cap = 0
// make takes only capacity argument when making channels - unlike slices there is no len argument
func makingChannels() {

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, '.', tabwriter.AlignRight|tabwriter.Debug)

	var c chan int
	fmt.Fprintf(w,"var c chan int\t  len(c): %v\t cap(c): %v\t c is nil: %v\n",len(c), cap(c), c==nil)

	c = make(chan int)
	fmt.Fprintf(w,"c = make(chan int)\t  len(c): %v\t cap(c): %v\t c is nil: %v\n",len(c), cap(c), c==nil)

	c = make(chan int, 5)
	fmt.Fprintf(w,"c = make(chan int, 5)\t  len(c): %v\t cap(c): %v\t c is nil: %v\n",len(c), cap(c), c==nil)

	w.Flush()
}

func main() {

	// point 1 - making channels
	// makingChannels()

	// point 2 - length of channels
	lenOfChannel()
}
