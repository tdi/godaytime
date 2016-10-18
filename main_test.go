package main

import (
	"bufio"
	"net"
	"strings"
	"testing"
	"time"
)

const connstring = "localhost:3333"

func TestUDPServer(t *testing.T) {
	go setupUDPServer(connstring)
	conn, err := net.Dial("udp4", connstring)
	if err != nil {
		t.Error("Expected to connect to ", connstring)
	}
	_, err = conn.Write([]byte("Give me date server!\n"))
	if err != nil {
		t.Error("Sending the datagram failed")
	}
	status, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		t.Error("No data received from server", err)
	}
	parsed, err := time.Parse(time.RFC1123, strings.Trim(status, "\n"))
	if err != nil {
		t.Error("Parse problem", parsed)
	}
	now := time.Now()
	if !now.Truncate(time.Hour).Equal(parsed.Truncate(time.Hour)) {
		t.Error("Dates differ", now.Truncate(time.Hour), parsed.Truncate(time.Hour))
	}
}

func TestTCPServer(t *testing.T) {
	go setupTCPServer(connstring)
	conn, err := net.Dial("tcp", connstring)
	if err != nil {
		t.Error("Expected to connect to ", connstring)
	}
	status, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		t.Error("No data from server")
	}
	parsed, err := time.Parse(time.RFC1123, strings.Trim(status, "\n"))
	if err != nil {
		t.Error("Parse problem", parsed)
	}
	now := time.Now()
	if !now.Truncate(time.Hour).Equal(parsed.Truncate(time.Hour)) {
		t.Error("Dates differ", now.Truncate(time.Hour), parsed.Truncate(time.Hour))
	}
}
