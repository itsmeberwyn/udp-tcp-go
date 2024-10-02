package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	port := 5001
	addr := &net.TCPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: port,
	}

	listener, err := net.ListenTCP("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	fmt.Println("Listening...")

	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()

		for {
			data, err := bufio.NewReader(conn).ReadString('\n')
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println("> ", string(data))

			conn.Write([]byte("Reply back hello data received!"))
		}

	}
}
