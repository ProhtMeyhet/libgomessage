package libgomessage

import(

)

type abstractSender struct {
	result chan *Message
}

func (abstract *abstractSender) InitAbstract() {
	abstract.result = make(chan *Message, 1)
}

func (abstract *abstractSender) GetResult() chan *Message {
	return abstract.result
}


type abstractServer struct {
	messages chan *Message
	control chan uint8
}

func (abstract *abstractServer) InitAbstract(numberOfConnections int) {
	if numberOfConnections < 1 {
		numberOfConnections = 1
	}
	abstract.messages = make(chan *Message, numberOfConnections)
	abstract.control = make(chan uint8, 1)
}

func (abstract *abstractServer) GetMessage() chan *Message {
	return abstract.messages
}
