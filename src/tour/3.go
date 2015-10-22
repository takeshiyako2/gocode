package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	fmt.Println("Welcome to the playground!")

	fmt.Println("The time is", time.Now())

	fmt.Println("And if you try to open a file:")
	fmt.Println(os.Open("filename"))

	fmt.Println("Or access the network:")

	// net.Dial
	conn, err := net.Dial("tcp", "www.yahoo.co.jp:80")
	if err != nil {
		log.Printf("Failed to connect to echoserver")
	}
	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	line, err := bufio.NewReader(conn).ReadString('\n')
	log.Printf(line)
}
