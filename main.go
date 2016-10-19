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
	version string = "0.4"
	address string = "127.0.0.1"
	port    string = "3333"
)

func printHelp() {
	fmt.Printf("go-daytime version: %s\n", version)
	fmt.Println("usage: go-daytime [-h] [-H HOST_NAME] [-p PORT]")
	os.Exit(0)
}

// Handles connection, returns date to a socket and status via a channel
func handleTCPConnection(conn *net.TCPConn, c chan string) error {
	log.Printf("New connection from %s", conn.RemoteAddr().String())
	defer conn.Close()
	dateTime := fmt.Sprintf("%s\n", time.Now().Format(time.RFC1123))
	_, err := conn.Write([]byte(dateTime))
	c <- fmt.Sprintf("done serving %s", conn.RemoteAddr().String())
	return err
}

// handleUDPClient handles a UDP daytime response
func handleUDPClient(ln *net.UDPConn, clientAddress *net.UDPAddr, c chan string) {
	dateTime := fmt.Sprintf("%s\n", time.Now().Format(time.RFC1123))
	_, err := ln.WriteToUDP([]byte(dateTime), clientAddress)
	if err != nil {
		log.Fatal(err)
	}
	c <- fmt.Sprintf("done serving %s", clientAddress)
	return
}

func setupUDPServer(connString string) {
	listenAddress, err := net.ResolveUDPAddr("udp4", connString)
	if err != nil {
		log.Fatal(err)
	}
	conn, err := net.ListenUDP("udp", listenAddress)
	defer conn.Close()

	if err != nil {
		log.Fatal(err)
	}
	log.Print("Listening on ", connString)
	c := make(chan string)
	buf := make([]byte, 1024)
	for {
		_, clientAddress, err := conn.ReadFromUDP(buf)
		log.Printf("New datagram from %s", clientAddress)
		if err != nil {
			log.Fatal(err)
		}
		go handleUDPClient(conn, clientAddress, c)
		log.Print(<-c)
	}
}

func setupTCPServer(connString string) {
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
		go handleTCPConnection(conn, c)
		log.Print(<-c)
	}
}

func main() {

	addressFlag := flag.String("H", address, "address to listen on default: localhost")
	portFlag := flag.String("p", port, "port to listen on, default: 2055")
	helpFlag := flag.Bool("h", false, "help message")
	protoFlag := flag.Bool("u", false, "listens on UDP")
	flag.Parse()

	if *helpFlag != false {
		printHelp()
	}

	connString := *addressFlag + ":" + *portFlag
	if *protoFlag != false {
		setupUDPServer(connString)
	}
	setupTCPServer(connString)
}
