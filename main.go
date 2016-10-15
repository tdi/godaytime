//Simple daytime server with go routines
package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

const (
	version string = "0.3"
	address string = "127.0.0.1"
	port    string = "3333"
)

func print_help() {
	fmt.Printf("go-daytime version: %s\n", version)
	fmt.Println("usage: go-daytime [-h] [-H HOST_NAME] [-p PORT]")
	os.Exit(0)
}

// Handles connection, returns date to a socket and status via a channel
func handleConnection(conn *net.TCPConn, c chan string) error {
	log.Printf("New connection from %s", conn.RemoteAddr().String())
	defer conn.Close()
	dateTime := fmt.Sprintf("%s\n", time.Now().Format(time.RFC1123))
	_, err := conn.Write([]byte(dateTime))
	c <- fmt.Sprintf("done serving %s", conn.RemoteAddr().String())
	return err
}

func main() {

	addressFlag := flag.String("H", address, "address to listen on default: localhost")
	portFlag := flag.String("p", port, "port to listen on, default: 2055")
	helpFlag := flag.Bool("h", false, "help message")
	flag.Parse()

	if *helpFlag != false {
		print_help()
	}

	connString := *addressFlag + ":" + *portFlag
	listenAddress, err := net.ResolveTCPAddr("tcp4", connString)

	if err != nil {
		log.Fatal(err)
	}

	ln, err := net.ListenTCP("tcp", listenAddress)
	defer ln.Close()

	if err != nil {
		log.Fatal(err)
	}

	log.Print("Listening on ", connString)
	c := make(chan string)
	for {
		conn, err := ln.AcceptTCP()
		if err != nil {
			log.Fatal(err)
		}
		go handleConnection(conn, c)
		log.Print(<-c)
	}
}
