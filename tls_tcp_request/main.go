package main

import (
	"bufio"
	"crypto/tls"
	"fmt"
)

const (
	request = "GET /get?foo=bar1 HTTP/1.1\r\n" +
		"Host: postman-echo.com\r\n" +
		"Accept: */*\r\n" +
		"Content-Type: application/json\r\n" +
		"Connection: close\r\n" +
		"\r\n"
)

func main() {
	conf := &tls.Config{}
	dial, err := tls.Dial("tcp", "postman-echo.com:443", conf)
	if err != nil {
		fmt.Println("Error dialing: ", err.Error())
		return
	}
	defer dial.Close()
	dial.Write([]byte(request))
	reader := bufio.NewReader(dial)
	for {
		line, err := reader.ReadString('\n')
		fmt.Print(line)
		if err != nil {
			break
		}
	}
}
