package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

// server: Struct for holding server info
type server struct {
	name string
	// List of clients
	clients []client
}

// client: Struct for holding client info
type client struct {
	username     string // Clients given username for global communication
	nickname     string
	password     string
	socket       net.Conn  // Connection interface
}

// addClient: Add client to list of clients on server
func (s *server) addClient(c client) {
	s.clients = append(s.clients, c)
}

func (c *client) readFromClient() string {
	clientData, err := bufio.NewReader(c.socket).ReadString('\n')
	if err != nil {
		log.Println(err)
		return "EOF"
	}

	cleanedData := strings.TrimSpace(strings.TrimRight(string(clientData), "\n"))
	return cleanedData
}

// HandleConnections: Main life cycle for every client connection
func (s *server) HandleConnection(c client) {
	defer c.socket.Close()
	// defer fmt.Println("Connection closed with client.")

	log.Println("Eternal conn loop")
	for {
		clientMsg := c.readFromClient()
		if clientMsg != "EOF" {
			log.Println(c.readFromClient())
		} else {
			return
		}
		// fmt.Println(cleanedData)

		// Do something with the incoming command here...
	}
}

// main: Starts server and allows new connections to be made
func main() {
	srv := &server{name: "TestServer"}
	fmt.Printf("Starting %s...\n", srv.name)

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
			socket: conn,
		}

		// fmt.Println("Connection made with client.")
		// Add client to the client list then begin client life cycle
		srv.addClient(newClient)
		go srv.HandleConnection(newClient)
	}
}
