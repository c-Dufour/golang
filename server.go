package main

import "net"
import "fmt"
import "bufio"
import "io/ioutil"
import (
	"strings"
	"strconv"
	"math/rand"
) // only needed below for sample processing

func main() {

	fmt.Println("Launching server...")

// read file
    b, err := ioutil.ReadFile("tmp/login.txt") // just pass the file name
    if err != nil {
        fmt.Print(err)
    }
    str := string(b) // convert content to a 'string'

    fmt.Println(str) // print the content as a 'string'
	// listen on all interfaces
	rand := strconv.Itoa(rand.Intn(10000000000))
	ln, _ := net.Listen("tcp", ":6969")
	ln.Addr().Network()
	fmt.Print("Random:", string(rand))

	// accept connection on port
	conn, _ := ln.Accept()

	

	// run loop forever (or until ctrl-c)
		// will listen for message to process ending in newline (\n)
		message, _ := bufio.NewReader(conn).ReadString('\n')
		// output message received
		fmt.Print("Message Received:", string(message), "from ")
		// sample process for string received
		newmessage := strings.ToUpper(message)
		// send new string back to client
		conn.Write([]byte(newmessage + "\n"))

}