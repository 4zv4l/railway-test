package main

import (
	"net"
	"bufio"
)

func handle(client net.Conn) {
	defer client.Close()
	r := bufio.NewReader(client)
	w := bufio.NewWriter(client)

	for {
		line, err := r.ReadString('\n')
		if err != nil {
			return
		}
		w.WriteString(line)
		w.Flush()
	}
}

func main() {
	server, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {panic(err)}
	for {
		client, err := server.Accept()
		if err != nil {panic(err)}
		go handle(client)

	}
}
