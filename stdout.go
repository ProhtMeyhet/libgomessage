package libgomessage

import(
	"fmt"
)

type Stdout struct {
	abstractSender
}

func NewStdout() *Stdout {
	return &Stdout{}
}

func (stdout *Stdout) Send(message *Message, to ToInterface) {
	fmt.Println(stdout.constructMessage(message))
	stdout.result <- &Message{ Result: SUCCESS,
					To: to.GetTos() }
}

func (stdout *Stdout) GetTo() ToInterface {
	return NewTo()
}

func (stdout *Stdout) constructMessage(message *Message) string {
	stringMessage := "Title: " + message.Title + "\r\n"
	stringMessage += "From: " + message.From + "\r\n"
	if message.Uri != "" {
		stringMessage += "Uri: " + message.Uri + "\r\n"
	}
	if message.Command != "" {
		stringMessage += "Command: " + message.Command + "\r\n"
	}
	stringMessage += "\r\n"
	stringMessage += message.Message
	return stringMessage
}
