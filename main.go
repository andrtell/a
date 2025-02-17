package main

import (
	"context"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
)

func client(_ context.Context, conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			switch err {
			case io.EOF:
				log.Printf("Client closed connection (EOF) from %v", conn.RemoteAddr())
			default:
				log.Print(err)
			}
			return
		}
		if n > 0 {
			conn.Write(buf[:n])
		}
	}
}

func server(ctx context.Context) {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		log.Printf("Client connected from %v", conn.RemoteAddr())
		go client(ctx, conn)
	}
}

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()
	go server(ctx)
	select {
	case <-ctx.Done():
		log.Print("Server shutting down (SIGINT)")
	}
}
