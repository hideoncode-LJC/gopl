package main

import (
	"go/gopl.io/config"
	"io"
	"log"
	"net"
	"os"
)

func main() {

	ncc := config.GetNetCat1()

	conn, err := net.Dial(ncc.IP, ncc.Port)

	if err != nil {
		log.Println(err)
	}

	defer conn.Close()

	mustCopy(os.Stdout, conn)

}

func mustCopy(dst io.Writer, src io.Reader) {

	/* Copy(io.Writer, io.Reader)
	从Reader中读取，写入Writer中。
	*/
	if _, err := io.Copy(dst, src); err != nil {
		log.Println(err)
	}
}


