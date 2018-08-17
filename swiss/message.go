package swiss

import "fmt"

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

func DefaultMessageProcessor(msg *Message) {
	fmt.Println("Processing message with the default processor...")
	fmt.Println(msg.RawString)
}
