package main

import "net"
import "fmt"
import (
	"crypto/sha1"
	"io"
	"bufio"
	"os"
)

func main() {

	// connect to this socket
	conn, _ := net.Dial("tcp", "127.0.0.1:6969")
	mess , _ := bufio.NewReader(conn).ReadString('\n')
	pseudo := bufio.NewReader(os.Stdin)
	fmt.Print(mess)

	fmt.Print("User Name: ")
	user, _ := pseudo.ReadString('\n')
	fmt.Print("Password: ")

	password, _ := pseudo.ReadString('\n')

	h := sha1.New()
	io.WriteString(h, user)
	io.WriteString(h, password)
	io.WriteString(h, mess)
	auth := (string(h.Sum(nil)))
	fmt.Fprintf(conn, user+":"+auth)

	//for {
	//	reader := bufio.NewReader(os.Stdin)
	//	fmt.Print("Text to send: ")
	//	text, _ := reader.ReadString('\n')
	//	fmt.Fprintf(conn, text + "\n")
	//	message, _ := bufio.NewReader(conn).ReadString('\n')
	//	fmt.Print("Message from server: "+message)
	//}
}