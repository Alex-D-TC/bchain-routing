package swiss

type Message struct {
	RawString string
}

func MessageFromBytes(raw []byte) (*Message, error) {
	rawString := string(raw)
	return &Message{
		RawString: rawString,
	}, nil
}

func (msg *Message) ToBytes() []byte {
	return []byte(msg.RawString)
}
