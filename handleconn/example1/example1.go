package main

import (
	log "github.com/sirupsen/logrus"
	"net"
	"os"
)

const (
	PORT = ":8000"
)

func handleConn(conn net.Conn) {
	defer conn.Close() // close once we are done

	log.WithFields(log.Fields{"remote addr:":conn.RemoteAddr().String()}).Info("handling connection")

	// step - handle
	for {

		/* Reader Interface Documentation (io.Reader)
		from https://golang.org/pkg/io/#Reader:

		Reader is the interface that wraps the basic Read method.

		Read reads up to len(p) bytes into p.
		It returns the number of bytes read (0 <= n <= len(p)) and any error encountered.
		Even if Read returns n < len(p), it may use all of p as scratch space during the call.
		If some data is available but not len(p) bytes, Read conventionally returns what is available instead of waiting for more.

		When Read encounters an error or end-of-file condition after successfully reading n > 0 bytes, it returns the number of bytes read.
		It may return the (non-nil) error from the same call or return the error (and n == 0) from a subsequent call.
		An instance of this general case is that a Reader returning a non-zero number of bytes at the end of the input stream may return either err == EOF or err == nil.
		The next Read should return 0, EOF.

		Callers should always process the n > 0 bytes returned before considering the error err.
		Doing so correctly handles I/O errors that happen after reading some bytes and also both of the allowed EOF behaviors.
		 */

		// buffer for reading
		msgIn := make([]byte, 100)

		// step - try read incoming message
		n, err := conn.Read(msgIn)
		if err != nil {
			log.WithFields(log.Fields{"err":err}).Error("error while reading from connection")
			continue
		}
		log.WithFields(log.Fields{"msgIn":msgIn, "numBytes":n}).Info("msg from client")

		// step - message back
		n, err = conn.Write(msgIn) // writer interface writes all msgIn buffer even though it is not full!
		if err != nil {
			log.WithFields(log.Fields{"err":err}).Error("error while writing to connection")
			continue
		}
		log.WithFields(log.Fields{"num bytes:":n}).Info("written")
	}

}

func main() {

	// step 1 - start listener
	l, err := net.Listen("tcp", PORT)
	if err != nil {
		log.WithFields(log.Fields{"err": err, "port": PORT}).Error("cant listen on Port")
		os.Exit(1)
	}

	// step 2 - accept incoming connections and handle
	for {
		conn, err := l.Accept()
		if err != nil {
			log.WithFields(log.Fields{"err": err}).Error("cant accept connection")
		}

		// step 3 - handle the connection
		handleConn(conn)
	}
}
