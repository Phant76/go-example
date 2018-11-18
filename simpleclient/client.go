package simpleclient

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"
)

// RunClient запускает клиент
func RunClient() {
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	s := make(chan string)
	e := make(chan error)
	reader := bufio.NewReader(os.Stdin)
	go readWithWait(reader, s, e)

	for {
		var source string
		select {
		case line := <-s:
			source = line
			s = make(chan string)
			e = make(chan error)
			go readWithWait(reader, s, e)
		case _ = <-e:
			source = "\r\n"
		case <-time.After(1 * time.Second):
			source = "\r\n"
		}
		if len(source) == 0 {
			continue
		}
		if n, err := conn.Write([]byte(source[:len(source)-1])); n == 0 || err != nil {
			fmt.Println(err)
			return
		}
		checkMessages(conn)
	}
}

func checkMessages(conn net.Conn) {
	buff := make([]byte, 1024)
	n, err := conn.Read(buff)
	if err != nil {
		return
	}
	resultStr := string(buff[0:n])
	if resultStr != " " {
		fmt.Print("\r" + string(buff[0:n]))
		fmt.Println()
	}
}

func readWithWait(reader *bufio.Reader, s chan string, e chan error) {
	line, err := reader.ReadString('\n')
	if err != nil {
		e <- err
	} else {
		s <- line
	}
	defer close(e)
	defer close(s)
}
