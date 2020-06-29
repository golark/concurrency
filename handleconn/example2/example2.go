package main

import (
	"bufio"
	log "github.com/sirupsen/logrus"
	"net"
	"os"
)

const (
	PORT = ":8000"
)

func handleConn(conn net.Conn) {

	defer func() {
		log.WithFields(log.Fields{"remote:":conn.RemoteAddr().String()}).Info("closing conn")
		conn.Close()
	} ()

	log.WithFields(log.Fields{"remote addr:":conn.RemoteAddr().String()}).Info("handling connection")

	buffReader := bufio.NewReader(conn)

	// step - handle
	for {

		// step - buffered reader
		bytes, err := buffReader.ReadBytes('\n')
		if err != nil {
			log.WithFields(log.Fields{"err":err}).Error("error while reading from connection")
			return
		}

		// step - message back
		n, err := conn.Write(bytes)
		if err != nil {
			log.WithFields(log.Fields{"err":err}).Error("error while writing to connection")
			return
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
