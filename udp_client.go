package main

import (
	"fmt"
	"net"
	"time"
)

// General function to check different errors
func CheckError(err error) {
    if err  != nil {
        fmt.Println("Error: " , err)
    }
}

func listenServerBroadcastUDP(){

	///////////////////////////
	// FINDING THE SERVER IP //
	///////////////////////////

	// Listen for incoming connections.
	serverAddress, err := net.ResolveUDPAddr("udp", ":30000")
	CheckError(err)
	conn, err := net.ListenUDP("udp", serverAddress)
	CheckError(err)

	fmt.Println("Before ReadFromUDP")
	message := make([]byte, 1024)
	conn.ReadFromUDP(message)					// Will not be read if server is down
	fmt.Println(string(message))
	fmt.Println("After ReadFromUDP")

	// Close the listener when the application closes.
	defer conn.Close()
}

func tranceiveServerUDP(){
	/////////////////////////
	// SENDING UDP PACKETS //
	/////////////////////////

	// Setup server (host)
	ServerAddr,err := net.ResolveUDPAddr("udp",":20018") // 129.241.187.38, // 255.255.255.255
	CheckError(err)

	// Setup local computer (client)
	//LocalAddr, err := net.ResolveUDPAddr("udp", "129.241.187.151:20018")
	//CheckError(err)

	Conn_wr, err := net.DialUDP("udp", nil, ServerAddr)
	CheckError(err)

	Conn_rd, err := net.ListenUDP("udp", ServerAddr)
	CheckError(err)

	// End the connection when the application closes
	defer Conn_wr.Close()
	defer Conn_rd.Close()

	// Read message space
	buf_rd := make([]byte, 1024)
	
	for {
		// Write message
		msg := "Hello World !"
		buf_wr := []byte(msg)

		// Send message to server
		_,err := Conn_wr.Write(buf_wr)
		CheckError(err)

		// Read message from server
		_,err = Conn_rd.ReadFromUDP(buf_rd)
		fmt.Println(string(buf_rd))
		CheckError(err)

		// Print message
		fmt.Println(buf_rd)
		time.Sleep(time.Second * 1)
	}
}

func main() {

	listenServerBroadcastUDP()
	tranceiveServerUDP()
}
