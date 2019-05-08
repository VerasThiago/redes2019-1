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
	// Stablish connection to the current host
	s.SocketClient, err = net.Dial("tcp", strings.Join([]string{s.Ip, strconv.Itoa(s.Port)}, ":"))
	return err

}

func (s *Socket) GetWebPage() string {

	// HTTP request
	request := "GET /" + s.Path + " HTTP/1.0\r\nHost: " + s.Ip + "\r\n\r\n"

	// Show current request
	fmt.Println(request)

	// Transform request into packet and send to the server
	s.SocketClient.Write([]byte(request))

	// Create a buffer to read response from server
	buffer := make([]byte, 1024)

	// String to get response from the server
	var response string

	// Read all the packets from the server
	for true {
		// Read from server and check if does not exist and err and the size is not 0
		n, err := s.SocketClient.Read(buffer)
		if err != nil || n == 0 {
			break
		}
		// Concatenate response into the response variable
		response += string(buffer[:n])
	}

	// Parser the response to get only the html
	split := strings.SplitAfterN(response, "<html", 2)

	// Check if the split was successful
	if len(split) > 1 {
		response = "<html" + split[1]
	}

	return response
}
