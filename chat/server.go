package chat

import (
	"bufio"
	"log"
	"net"
	"strconv"
	"strings"
)

type Server struct {
	listener net.Listener
	rooms    map[int]*Room
}

func (s Server) Run() error {
	//	::: Strart to accept connections

	for {

		//::: Accept conn
		conn, err := s.listener.Accept()
		log.Println("New client ...")
		if err != nil {
			log.Println("Error accepting: ", err.Error())
			continue
		}
		//	::: create client

		client := NewClient(conn)

		//	::: launch client
		s.handleClient(client)
		go client.ServeItself()
	}

}

func (s Server) handleClient(client *Client) {

	msgWelcome := Message{
		Author:  &ADMIN,
		Payload: "Enter room id you want to come in\n",
	}
	client.WriteMsg(&msgWelcome)

	roomIDToEnter, err := bufio.NewReader(client.conn).ReadString('\n')
	roomIDToEnter = strings.Trim(roomIDToEnter, "\r\n")
	roomId, err := strconv.Atoi(roomIDToEnter)
	if err != nil {
		log.Printf("Couldn't convert to int: %s\n", roomIDToEnter)
		return
	}

	room := s.roomToEnterById(roomId)
	if room == nil {
		log.Printf("Couldn't find room %d\n", roomId)
		return
	}

	println("Enter room id you want to enter: ", roomIDToEnter)
	room.RegisterNewClient(client)
}

func (s Server) roomToEnterById(id int) *Room {
	for _, room := range s.rooms {
		if room.id == id {
			return room
		}

	}
	return nil
}

func NewServer(lis net.Listener) *Server {
	rooms := make(map[int]*Room)
	intialRoom := &Room{id: 0, clientBus: make(chan *Client), msgBus: make(chan *Message)}
	go intialRoom.ServerItself()
	rooms[intialRoom.id] = intialRoom
	return &Server{listener: lis, rooms: rooms}

}
