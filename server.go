package main

import (
	"log"
	"net"
	"os"
)

func main() {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", ":9000")
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	log.Println("Game Server OK.")

	for {
		conn, err := listener.AcceptTCP()

		if err != nil {
			log.Println("accept failed", err)
			continue
		}

		go handleClient(conn)
	}

}

func handleClient(conn *net.TCPConn) {
	defer conn.Close()

	var buf [4]byte
	for {
		n, err := conn.Read(buf[0:])
		if err != nil {
			return
		}

		log.Println("request messge:", string(buf[0:]))
		if n > 0 {
			n, err = conn.Write([]byte("Pone"))
			if err != nil {
				return
			}
		}
	}
}

func checkError(err error) {
	if err != nil {
		log.Println("Fatal error:", err)
		os.Exit(-1)
	}
}
