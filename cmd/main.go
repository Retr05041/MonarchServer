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
	name string
	// List of clients
	clients []client
}

// client: Struct for holding client info
type client struct {
	username         string // Clients given username for global communication
	nickname         string
	password         string
	clientWriter     io.Writer // Uses conn as it's io.Writer
	clientConnection net.Conn  // Connection interface
}

// addClient: Add client to list of clients on server
func (s *server) addClient(c client) {
	s.clients = append(s.clients, c)
}

func (c *client) readFromClient() string {
	clientData, err := bufio.NewReader(c.clientConnection).ReadString('\n')
	if err != nil {
		log.Println(err)
		return "EOF"
	}

	cleanedData := strings.TrimSpace(strings.TrimRight(string(clientData), "\n"))
	return cleanedData
}

func (c *client) registerClient() {
	log.Println("Entered registerClient")
	for {
		cmd := strings.Split(c.readFromClient(), " ")
		log.Println(cmd)
		switch cmd[0] {
		case "PASS":
			c.password = cmd[1]
		case "NICK":
			c.nickname = cmd[1]
		case "USER":
			c.username = cmd[1]
		case "CAP":
			if cmd[1] == "LS" {
			}
			if cmd[1] == "END" {
				return
			}
		default:
			return
		}
	}
}

func (s *server) sendWelcome(cl client) {
	rplWelcome := "Your host is " + s.name + ", running version 1.0"
	rplCreated := "This server was created at the beginning of time"
	rplMyInfo := "Blah blah blah"
	rplIsSupport := "Nothing is supported"
	_, err := cl.clientWriter.Write([]byte(rplWelcome))
	if err != nil {
		log.Println(err)
	}
	_, err = cl.clientWriter.Write([]byte(rplCreated))
	if err != nil {
		log.Println(err)
	}
	_, err = cl.clientWriter.Write([]byte(rplMyInfo))
	if err != nil {
		log.Println(err)
	}
	_, err = cl.clientWriter.Write([]byte(rplIsSupport))
	if err != nil {
		log.Println(err)
	}
	return
}

// HandleConnections: Main life cycle for every client connection
func (s *server) HandleConnection(c client) {
	defer c.clientConnection.Close()
	// defer fmt.Println("Connection closed with client.")

	c.registerClient()
	s.sendWelcome(c)

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
	fmt.Println("Starting server...")
	srv := &server{name: "TestServer"}

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
