package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"

)

// server: Struct for holding server info
type server struct {
	// List of clients
	clients []client
}

// client: Struct for holding client info
type client struct {
	username         string    // Clients given username for global communication
	clientWriter     io.Writer // Uses conn as it's io.Writer
	clientConnection net.Conn  // Connection interface
}

// addClient: Add client to list of clients on server
func (s *server) addClient(c client) {
	s.clients = append(s.clients, c)
}


// HandleConnections: Main life cycle for every client connection
func (s *server) HandleConnection(c client) {
	defer c.clientConnection.Close()
	// defer fmt.Println("Connection closed with client.")

	for {
		clientData, err := bufio.NewReader(c.clientConnection).ReadString('\n')
		if err != nil {
			log.Println(err)
      return
		}

		cleanedData := strings.TrimSpace(strings.TrimRight(string(clientData), "\n"))
		fmt.Println(cleanedData)
		// fmt.Println(cleanedData)

    // Do something with the incoming command here...
	}
}

// main: Starts server and allows new connections to be made 
func main() {
	fmt.Println("Starting server...")
	srv := &server{}

	listener, err := net.Listen("tcp", ":6667")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Listening on port 6667...")
	defer listener.Close()

	// Accept incoming connections forever -- This will need to be changed in the future...
	for {
		// Accept a new connection
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		// Create a client
		newClient := client{
			clientWriter:     conn,
			clientConnection: conn,
		}

		// fmt.Println("Connection made with client.")
		// Add client to the client list then begin client life cycle
		srv.addClient(newClient)
		go srv.HandleConnection(newClient)
	}
}
