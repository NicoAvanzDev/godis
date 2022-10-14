package main

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"sync"
)

var dict sync.Map

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)

	var commands []string

	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			break
		}

		sz, _ := strconv.Atoi(string(line[1:]))

		for i := 0; i < sz; i++ {
			_, _, _ = reader.ReadLine()
			line, _, _ = reader.ReadLine()

			commands = append(commands, string(line))
		}

		reply := executeCommands(commands)

		if reply == "" {
			writer.WriteString("$-1\r\n")
		} else {
			writer.WriteString(fmt.Sprintf("$%d\r\n%s\r\n", len(reply), reply))
		}

		writer.Flush()
		commands = nil
	}
}

func executeCommands(commands []string) string {
	switch commands[0] {
	case "SET":
		dict.Store(commands[1], commands[2])

	case "GET":
		value, ok := dict.Load(commands[1])
		if ok {
			return value.(string)
		}
	}
	return ""

}
