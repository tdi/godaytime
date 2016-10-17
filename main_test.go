package main

import (
	"bufio"
	"net"
	"strings"
	"testing"
	"time"
)

func TestTCPServer(t *testing.T) {
	connstring := "localhost:3333"
	go setupTCPServer(connstring)
	conn, err := net.Dial("tcp", connstring)
	if err != nil {
		t.Error("Expected to connect to ", connstring)
	}
	status, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		t.Error("No data from server")
	}
	// status = "Mon, 17 Oct 2016 12:29:15 CEST"
	parsed, err := time.Parse(time.RFC1123, strings.Trim(status, "\n"))
	if err != nil {
		t.Error("Parse problem")
	}
	now := time.Now()

	if now.Truncate(time.Hour) != parsed.Truncate(time.Hour) {
		t.Error()
	}
}
