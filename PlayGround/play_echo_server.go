package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

const (
	CONN_HOST = ""
	CONN_PORT = "3333"
	CONN_TYPE = "tcp"
)

func main() {
	// Listen of incoming connections
	listen, err := net.Listen(CONN_TYPE, ":"+CONN_PORT)
	if err != nil {
		// log.Fatal makes an os call to exit the program after logging
		// the issue
		log.Fatal(err)
	}
	// Ensures that the connection is closed properly
	defer listen.Close()

	fmt.Println("Listening on: " + CONN_HOST + ":" + CONN_PORT)

	// infinite loop to listen of incoming connections
	for {
		// listen of incoming connections
		conn, err := listen.Accept()
		if err != nil {
			log.Fatal(err)
		}

		// print info about incoming request
		fmt.Printf("Received connection from: %s -> %s \n", conn.RemoteAddr(),
			conn.LocalAddr())

		// Handle Connection in goroutine
		go HandleConn(conn)
	}

}

func HandleConn(conn net.Conn) {
	defer conn.Close()
	input := bufio.NewScanner(conn)
	for input.Scan() {
		msg := input.Text()
		fmt.Println("Text: " + msg)
		fmt.Fprintln(conn, "Received: "+msg)
	}
}
