package CLI

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/websocket"
)

var num int

func SetupRouter(port int, route string) {
	//fmt.Println("CLI has successfully connected with your local app")
	fmt.Print("Webhook tester is hosted at port :8000\n\n")

	// Calls whtestServerConnection which attempts to connect to the online hosted
	// server through websockets through which data is transferred between servers
	go whtestServerConnection("ws://whtest.pc-1827.online/whtest", port, route)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

var subdomainReceived = false

func SubdomainHandler(conn *websocket.Conn, port int, route string) {
	//fmt.Print("Attempting to receive Subdomain.\n")
	_, subdomain, err := conn.ReadMessage()
	if err != nil {
		fmt.Print("Error receiving Subdomain:\n\n", err)
		return
	}

	if string(subdomain) == "None" {
		fmt.Print("No subdomain available\n\n")
	} else {
		conn.Close()
		subdomainReceived = true
		subdomainURL := "ws://" + string(subdomain) + route
		go whtestServerConnection(string(subdomainURL), port, route)
		// go whtestServerConnection("ws://localhost:2001/subdomain", port, route)

		localServerURL := "http://localhost:" + strconv.Itoa(port) + route

		fmt.Printf("WebSocket traffic will be transferred from %s ---> %s\n\n", subdomain, localServerURL)
	}

}

// After successfully establishing a websocket connection between CLI and server hosted
// MessageTransfer is used to send an encoded message to the hosted server, which helps in
// identifying if the message is received by the CLI or not.
func MessageTransfer(conn *websocket.Conn) {
	//fmt.Println("Inside MessageTransfer function")
	// Log the current connection state
	if conn == nil {
		log.Print("WebSocket connection is nil\n\n")
		return
	}

	encodedMessage := "afjhfa793n&%$$kjhfah8H&h88h78uHY&99yhyfauhh8YUIH98Hh9``hhhre9rfhh93%4&"
	//fmt.Print("Encoded message is being transferred.\n")

	// Log before sending the message
	//log.Println("Attempting to send message:", encodedMessage)

	if err := conn.WriteMessage(websocket.TextMessage, []byte(encodedMessage)); err != nil {
		log.Print("Error connecting withe the online server\n\n", err)
		return
	}

	// Log after successful send
	//log.Println("Message sent successfully")
}

func whtestServerConnection(URL string, port int, route string) {
	//fmt.Print("Hello, trying to connect to whtest_server.\n")

	conn, _, err := websocket.DefaultDialer.Dial(URL, nil)
	if err != nil {
		log.Print("WebSocket dial error:\n\n", err)
		return
	}

	fmt.Print("Successfully connected with whtest server\n\n")

	// Call MessageTransfer function
	//fmt.Println("Calling MessageTransfer function")
	MessageTransfer(conn)

	if !subdomainReceived {
		SubdomainHandler(conn, port, route)
	} else {
		DataTransferHandler(conn, port, route)
	}
}
