package main

import (
	"log"
	"net"
	"os"
)

func main() {

	tcpAddr, err := net.ResolveTCPAddr("tcp4", "localhost:9000")
	checkError(err)

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)

	conn.Write([]byte("Ping"))

	var buf [4]byte
	_, err = conn.Read(buf[0:])

	checkError(err)

	log.Println("response messge:", string(buf[0:]))

	conn.Close()
}

func checkError(err error) {
	if err != nil {
		log.Println("Fatal error:", err)
		os.Exit(-1)
	}
}
