package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func check(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	ln, err := net.Listen("tcp", ":8080")
	check(err)
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		check(err)
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	request(conn)
	respond(conn)
}

func request(conn net.Conn) {
	i := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		if i == 0 {
			url := strings.Fields(ln)[1]
			fmt.Println(url)
		}
		if ln == "" {
			break
		}
		i++
	}
}

func respond(conn net.Conn) {
	body := `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title></title></head><body><strong>HelloThisIsFromTheServer</strong></body></html>`
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}
