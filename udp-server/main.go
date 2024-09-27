package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	port := 5001
	addr := &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: port,
		Zone: "",
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	fmt.Println("Listening...")

	for {
		fmt.Println("hello")
		var buf [512]byte
		_, addr, err := conn.ReadFromUDP(buf[0:])
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Print("> ", string(buf[0:]), addr)

		// Write back the message over UPD
		// conn.WriteToUDP([]byte("Hello UDP Client\n"), addr)
	}
}
