package client

import (
	"bufio"
	"os"
	"fmt"
	"net"
)




func main(){
	conn, err := net.Dial("tcp", "localhost:7998")
	if err != nil {
		fmt.Println("Error in connecting to server")
		os.Exit(1)
	}
	defer conn.Close()

	reader := bufio.NewReader(conn)
	response, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error in readying", err)
		return
	}
	response = response[:len(response)-1]

	if response == "OK"{
		fmt.Println("All is good !")
	} else {
		fmt.Println("Unexpected response:", response)
	}
}