package chat

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

type Client struct {
	conn    net.Conn
	name    string
	room    *Room
	msgType MsgType
}

func (c *Client) ServeItself() {
	for {
		msgString, err := bufio.NewReader(c.conn).ReadString('\n')
		if err != nil {
			log.Printf("Error reading from server: %v", err)
		}

		msg := Message{
			Type:    c.msgType,
			Payload: msgString,
			Author:  c,
		}

		c.room.msgBus <- &msg

	}

}

func (c *Client) WriteMsg(msg *Message) {
	payload := fmt.Sprintf("%s > %s", msg.Author.name, msg.Payload)
	_, err := c.conn.Write([]byte(payload))
	if err != nil {
		log.Printf("Error writing to server: %v", err)
		return
	}
}

func (c *Client) SetName(nameToSet string) {
	c.name = nameToSet

}

func NewClient(conn net.Conn) *Client {
	return &Client{
		conn:    conn,
		msgType: SET_NICK_NAME,
	}
}
