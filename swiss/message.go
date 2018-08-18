package swiss

import "fmt"

type Message struct {
	RawString string
}

func DefaultMessageProcessor(msg *Message) {
	fmt.Println("Processing message with the default processor...")
	fmt.Println(msg.RawString)
}
