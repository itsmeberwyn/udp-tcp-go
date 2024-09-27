package main

import (
	"net"
	"log"
	"fmt"
	"os"
)

func main() {
	port := 5001
	addr := &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: port,
		Zone: "",
	}

	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		log.Fatal(err)
	}

	_, err = conn.Write([]byte("message from client heheheelllooo\n"))
	fmt.Println("send...")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}