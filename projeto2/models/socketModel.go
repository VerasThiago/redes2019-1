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
	n, err := s.SocketClient.Read(buff)
	if err != nil {
		fmt.Print("Deu merda pra ler do server = ")
		fmt.Println(err)
	}

	aux := string(buff[:n])
	teste := strings.SplitAfterN(aux, "{", 2)
	if len(teste) > 1 {
		aux = teste[1]
		aux = strings.Replace(aux, "[", "", -1)
		aux = strings.Replace(aux, "]", "", -1)
		aux = strings.Replace(aux, "{", "", -1)
		aux = strings.Replace(aux, "}", "", -1)
		aux = strings.Replace(aux, `"result":`, "", -1)
	} else {
		fmt.Println("Len eh < 1")
		fmt.Println("AUX = " + aux)
	}
	return "{" + aux + "}"
}
