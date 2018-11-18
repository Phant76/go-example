package simpleserver

import (
	"fmt"
	"net"
	"strings"
)

var messages []string

// RunServer запускает сервер
func RunServer() {
	listener, err := net.Listen("tcp", ":8888")

	if err != nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()
	fmt.Println("Сервер запущен...")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			conn.Close()
			continue
		}
		go handleConnection(conn)
	}
}

func checkNewMessages(curMsg int) []string {
	if curMsg < len(messages) {
		return messages[curMsg:]
	}
	return make([]string, 0)
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	curMsg := 0
	for {
		input := make([]byte, (1024 * 4))
		n, err := conn.Read(input)
		if n == 0 || err != nil {
			fmt.Println("Соединение закрыто:", err)
			break
		}
		msg := strings.Replace(string(input[0:n]), "\r", " ", -1)
		if msg != " " {
			messages = append(messages, conn.RemoteAddr().String()+": "+msg)
		}
		result := checkNewMessages(curMsg)
		curMsg = curMsg + len(result)
		strResult := strings.Join(result, "\n")
		if len(strResult) > 0 {
			fmt.Println(conn.RemoteAddr().String() + ": " + msg + " -> " + strResult)
		}
		conn.Write([]byte(strResult + " "))
	}
}
