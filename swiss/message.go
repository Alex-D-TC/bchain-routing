package swiss

type Message struct {
	rawString string
}

func MessageFromBytes(raw []byte) (*Message, error) {
	rawString := string(raw)
	return &Message{
		rawString: rawString,
	}, nil
}

func (msg *Message) ToBytes() []byte {
	return []byte(msg.rawString)
}
