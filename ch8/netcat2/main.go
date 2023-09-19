package main

import (
	"net"
	"os"
)


func main() {
	conn, err := net.Dial(network, address)

	if err != nil {
		log.Println(err)
	}

	defer conn.Close()

	go mustCopy(os.Stdout, conn)

	mustCopy(conn, os.Stdin)
}