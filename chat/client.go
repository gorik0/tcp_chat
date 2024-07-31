package chat

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

type Client struct {
	conn    net.Conn
	name    string
	room    *Room
	msgType MsgType
}

var ADMIN = Client{
	name: "ADMIN",
}

func (c *Client) ServeItself() {
	for {
		msgString, err := bufio.NewReader(c.conn).ReadString('\n')
		if err != nil {
			log.Printf("Error reading from server: %v", err)
		}
		msgString = strings.Trim(msgString, "\r\n")

		msg := Message{
			Type:    c.msgType,
			Payload: msgString,
			Author:  c,
		}

		c.room.msgBus <- &msg

	}

}

func (c *Client) WriteMsg(msg *Message) {
	payload := "\r\n" + fmt.Sprintf("%s > %s", msg.Author.name, msg.Payload) + "\n"
	_, err := c.conn.Write([]byte(payload))
	if err != nil {
		log.Printf("Error writing to server: %v", err)
		return
	}
}

func (c *Client) SetName(nameToSet string) {
	c.name = nameToSet

}

func (c *Client) WriteEmptyMsg() {
	c.conn.Write([]byte("\t"))
}

func NewClient(conn net.Conn) *Client {
	return &Client{
		conn:    conn,
		msgType: SET_NICK_NAME,
	}
}
