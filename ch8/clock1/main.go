package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {

	log.SetFlags(log.Lshortfile | log.Lmicroseconds | log.Ldate)

	log.SetPrefix("\033[34m[info]\033[0m ")

	// 创建listener对象，监听IP:HOST端口
	listener, err := net.Listen("tcp", "localhost:8080")

	log.Println("开始监听端口")

	if err != nil {
		log.Fatal(err)
	}

	/*
	Accept()方法listener对象的Accept方法会直接阻塞，
	直到一个新的连接被创建，
	然后会返回一个net.Conn对象来表示这个连接。
	*/ 
	for {
		conn, err := listener.Accept()
	
		log.Println("获取到一个新的连接")

		log.Println("LocalAddr", conn.LocalAddr())

		log.Println("RemoteAddr", conn.RemoteAddr())

		if err != nil {
			log.Println(err)
			continue
		}

		// 处理一个连接
		handleConn(conn)
	}

}

func handleConn(conn net.Conn) {
	
	defer conn.Close()

	for {

		/* 	conn也实现了Write方法，conn也是一个io.Writer接口
		 	time.Now().Format("15:04:05\n")返回的是各种样式的时间字符中
		 	这个死循环会一直执行，直到写入失败
			最可能的原因是客户端主动断开连接。
			这种情况下handleConn函数会用defer调用关闭服务器侧的连接，
			然后返回到主函数，
			继续等待下一个连接请求。
		*/

		_, err := io.WriteString(conn, time.Now().Format("15:04:05\n"))

		if err != nil {
			log.Println(err)
		}

		time.Sleep(time.Second * 1)
	}
}