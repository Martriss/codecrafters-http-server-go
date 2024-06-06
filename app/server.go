package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

const CRLF = "\r\n"

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

		requestLine := strings.Split(req, CRLF)[0]
		headers := strings.Split(strings.Split(req, CRLF)[1], CRLF)
		// body := strings.Split(req, CRLF+CRLF)[1]

		// method := strings.Split(requestLine, " ")[0]
		path := strings.Split(requestLine, " ")[1]
		// HTTPVersion := strings.Split(requestLine, " ")[2]

		pathFragments := strings.Split(path, "/")

		res := []byte("HTTP/1.1 404 Not Found\r\n\r\n")
		if path == "/" {
			res = []byte("HTTP/1.1 200 OK\r\n\r\n")
		} else if pathFragments[1] == "echo" {
			msg := pathFragments[2]
			res = []byte(fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: %d\r\n\r\n%s", len(msg), msg))
		} else if pathFragments[1] == "user-agent" {

		}

		conn.Write(res)
		conn.Close()
	}
}
