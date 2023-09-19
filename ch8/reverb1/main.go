package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"time"
)

/*
大多数echo服务仅仅会返回他们读取到的内容
*/



func echo(c net.Conn, shout string, delay time.Duration) {
	
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))

	time.Sleep(delay)

	fmt.Fprintln(c, "\t", shout)

	time.Sleep(delay)

	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)
	for input.Scan() {
		echo(c, input.Text(), 1 * time.Second)
	}
	c.Close()
}