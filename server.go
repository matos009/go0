package server

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	
	listener, err := net.Listen("tcp", "localhost:7998")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	defer listener.Close()
	fmt.Println("Server launched")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		go connHandler(conn) 
	}
}

func connHandler(conn net.Conn) {
	defer conn.Close()

	writer := bufio.NewWriter(conn)
	_, err := writer.WriteString("OK\n")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	writer.Flush()
}






