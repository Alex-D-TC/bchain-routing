package swiss

type Message struct {
}

func MessageFromBytes([]byte) (*Message, error) {
	return &Message{}, nil
}

func (msg *Message) ToBytes() []byte {
	return []byte{}
}
