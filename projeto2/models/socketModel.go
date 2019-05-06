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
	SocketClient net.Conn
}

func (s *Socket) Init() {
	var err error
	s.SocketClient, err = net.Dial("tcp", strings.Join([]string{s.Ip, strconv.Itoa(s.Port)}, ":"))
	if err != nil {
		fmt.Println("Deu merda pra criar o socket")
		fmt.Println(err)
	} else {
		fmt.Println("Deu bom pra criar o socket")
	}

}

func (s *Socket) GetMessage(handle string) string {
	message := "GET /api/user.info?handles=" + handle + " HTTP/1.0\r\nHost: codeforces.com\r\n\r\n"
	s.SocketClient.Write([]byte(message))
	buff := make([]byte, 1024)
	n, _ := s.SocketClient.Read(buff)
	aux := string(buff[:n])
	teste := strings.SplitAfterN(aux, "{", 2)
	aux = teste[1]
	aux = strings.Replace(aux, "[", "", -1)
	aux = strings.Replace(aux, "]", "", -1)
	aux = strings.Replace(aux, "{", "", -1)
	aux = strings.Replace(aux, "}", "", -1)
	aux = strings.Replace(aux, `"result":`, "", -1)
	return "{" + aux + "}"
}
