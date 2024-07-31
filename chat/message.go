package chat

type Message struct {
	Author  *Client
	Type    MsgType
	Payload string
}

type MsgType string

const (
	SAY           MsgType = "SAY"
	SET_NICK_NAME MsgType = "SET_NICK_NAME"
	MAY_I_COME    MsgType = "MAY_I_COME"
)
