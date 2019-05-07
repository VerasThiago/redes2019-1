package models

import (
	"fmt"
	"net"
	"strconv"
	"strings"
)

type Socket struct {
	Ip           string
	Port         int
	Path         string
	SocketClient net.Conn
}

func (s *Socket) Init() error {
	var err error
	s.SocketClient, err = net.Dial("tcp", strings.Join([]string{s.Ip, strconv.Itoa(s.Port)}, ":"))
	return err

}

func (s *Socket) GetWebPage() string {

	message := "GET /" + s.Path + " HTTP/1.0\r\nHost: " + s.Ip + "\r\n\r\n"

	fmt.Println(message)

	s.SocketClient.Write([]byte(message))
	buffer := make([]byte, 1024)

	response := ""

	for true {
		n, err := s.SocketClient.Read(buffer)
		if err != nil || n == 0 {
			break
		}
		response += string(buffer[:n])
	}

	split := strings.SplitAfterN(response, "<html", 2)

	if len(split) > 1 {
		response = "<html" + split[1]
	}

	//response = strings.Replace(response, "href="+`"`+"/", "href="+`"`+s.Ip+"/", -1)

	return response
}
