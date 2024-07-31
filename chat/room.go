package chat

type Room struct {
	id      int
	clients []*Client
	msgBus  chan *Message
}

func (r *Room) ServerItself() {
	for {
		msg := <-r.msgBus
		switch msg.Type {

		case SAY:
			{

				for _, client := range r.clients {
					if client != msg.Author {
						client.WriteMsg(msg)

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
