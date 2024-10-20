package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	l, err := net.Listen("tcp4", ":9000")
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	for true {
		conn, err := l.Accept()
		defer conn.Close()
		if err != nil {
			fmt.Println(err)
			continue
		}
		buff := make([]byte, 1024)
		conn.Read(buff)
		fmt.Println(buff)
	}
}
