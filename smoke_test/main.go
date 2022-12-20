package main

import (
	"io"
	"log"
	"net"
)

func main() {
	log.Println("smoke test server v1.0.0")

	serv, err := net.Listen("tcp", "0.0.0.0:31337")
	if err != nil {
		log.Fatalf("ERROR: Listen: %v\n", err)
	}

	defer serv.Close()
	log.Println("Listening on 0.0.0.0:31337")

	for {
		conn, err := serv.Accept()
		if err != nil {
			log.Fatalf("ERROR: Accept: %v\n", err)
		}

		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	buf := make([]byte, 16*1024*1024) // 16MiB buffer, should be capable of holding anything

	len, err := conn.Read(buf)
	if err != nil {
		if err != io.EOF {
			log.Fatalf("ERROR: Read: %v\n", err)
		}
	}

	log.Printf("Received message of %d bytes\n", len)

	conn.Write(buf[:len])

	conn.Close()
}
