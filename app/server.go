package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221: ", err)
		os.Exit(1)
	}
	defer l.Close()
	fmt.Println("Listening...")

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}

		buf := make([]byte, 512)
		_, err = conn.Read(buf)
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}

		req := string(buf)

		target := strings.Split(req, " ")[1]
		res := []byte("HTTP/1.1 404 Not Found\r\n\r\n")
		if target == "/" {
			res = []byte("HTTP/1.1 200 OK\r\n\r\n")
		}
		conn.Write(res)
		conn.Close()
	}
}
