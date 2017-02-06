package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
)

// defining how big the packets will be
const BUFFERSIZE = 1024

func main() {
	// defining two functions, one listens of tcp, the other port 27001
	server, err := net.Listen("tcp", "localhost:27001")
	// if there is something on 27001, it throughs an error
	if err != nil {
		fmt.Println("Error listetning: ", err)
		os.Exit(1)
	}
	// if the search goes out of scope the operation is shut down
	defer server.Close()
	// this loop wait until someone connects, then sets up a go routine to send the file to the client then starts the loop again
	for {
		connection, err := server.Accept()
		if err != nil {
			fmt.Println("Error: ", err)
			os.Exit(1)
		}
		fmt.Println("Reciever connected")
		go sendFileToClient(connection)

	}
}

// this is the juicy part (the part "go" called)
func sendFileToClient(connection net.Conn) {
	fmt.Println("A reciever has connected!")
	defer connection.Close()
	file, err := os.Open("dummyfile.dat")
	if err != nil {
		fmt.Println(err)
		return
	}
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return
	}
	fileSize := fillString(strconv.FormatInt(fileInfo.Size(), 10), 10)
	fileName := fillString(fileInfo.Name(), 64)
	fmt.Println("Sending filename and filesize!")
	connection.Write([]byte(fileSize))
	connection.Write([]byte(fileName))
	sendBuffer := make([]byte, BUFFERSIZE)
	fmt.Println("Start sending file!")
	for {
		_, err = file.Read(sendBuffer)
		if err == io.EOF {
			break
		}
		connection.Write(sendBuffer)
	}
	fmt.Println("File has been sent, closing connection!")
	return
}
func fillString(retunString string, toLength int) string {
	for {
		lengtString := len(retunString)
		if lengtString < toLength {
			retunString = retunString + ":"
			continue
		}
		break
	}
	return retunString
}
