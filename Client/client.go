package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp4", ":9000")
	defer conn.Close()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	conn.Write([]byte("Hello from the client"))
}
