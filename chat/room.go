package chat

import (
	"log"
)

type Room struct {
	id        int
	clients   []*Client
	msgBus    chan *Message
	clientBus chan *Client
}

func (r *Room) ServerItself() {
	println("Starting room_serve wheel...for room ::: ")
	for {
		select {

		case client := <-r.clientBus:
			{
				log.Printf("Welome home, man, %s !!!", client.name)
				r.clients = append(r.clients, client)
				client.room = r
				msgSetNamePlease := Message{
					Author:  &ADMIN,
					Payload: "Introduce yourself ... ",
				}
				client.WriteMsg(&msgSetNamePlease)

			}

		case msg := <-r.msgBus:
			{
				log.Printf("MSG")

				switch msg.Type {

				case SAY:
					{

						for _, client := range r.clients {
							if client != msg.Author {
								client.WriteMsg(msg)

							} else {
								client.WriteEmptyMsg()

							}

						}

					}
				case SET_NICK_NAME:
					{

						msg.Author.SetName(msg.Payload)
						msg.Author.msgType = SAY
					}

				}

			}
		}

	}

}

func (r *Room) RegisterNewClient(client *Client) {

	r.clientBus <- client

}
