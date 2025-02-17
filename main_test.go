package main

import (
	"testing"
	"time"
	"net"
)

func TestMain(t *testing.T) {
	go main()

	time.Sleep(1 * time.Second)

	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	msg := "1234"

	n, err := conn.Write([]byte(msg))

	if err != nil {
		t.Fatal(err)
	}

	buf := make([]byte, 1024)

	m, err := conn.Read(buf)

	if err != nil {
		t.Fatal(err)
	}

	if n != m {
		t.Fatalf("length of message %v does not match length of echo %v", n , m)
	}

	if msg != string(buf[:m]) {
		t.Fatalf("message \"%v\" does not match echo \"%v\"", msg, string(buf))
	}
}
