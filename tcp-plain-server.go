package libgomessage

import(
	"bufio"
	"io/ioutil"
	"net"
	"net/textproto"
	"time"
)

type TcpPlainServer struct {
	TcpServer
	parser ParseInterface
}

func NewTcpPlainServer(config *TcpServerConfig) *TcpPlainServer {
	tcpPlain := &TcpPlainServer{}
	// we are and can be our own parser
	tcpPlain.Init(config, tcpPlain)
	return tcpPlain
}

func (tcp *TcpPlainServer) Init(config *TcpServerConfig, parser ParseInterface) {
	tcp.parser = parser
	tcp.TcpServer.Init(config, tcp)
}

func (tcp *TcpPlainServer) Handle(connection net.Conn) *Message {
	defer connection.Close()

	select {
	case message := <-tcp.read(connection):
		return message
	case <-time.After(time.Duration(tcp.config.ReadTimeOut) * time.Second):
	}

	return &Message{ Result: FAILURE,
			ErrorString: "Read timeout!",
			FromRemote: connection.RemoteAddr().String() }
}

func (tcp *TcpPlainServer) read(connection net.Conn) <-chan *Message {
	returnChannel := make(chan *Message, 1)
	reader := bufio.NewReader(connection)
	returnMessage := tcp.parser.Parse(reader)
	returnMessage.FromRemote = connection.RemoteAddr().String()
	returnChannel <- returnMessage
	return returnChannel
}

func (plain *TcpPlainServer) Parse(from *bufio.Reader) (message *Message) {
	message = new(Message)
	var messageBytes []byte

	fromProto := textproto.NewReader(from)
	headers, e := fromProto.ReadMIMEHeader()
	if e != nil {
		goto out
	}

	for key, value := range headers {
		if len(value) < 1 {
			continue
		}
		switch key {
		case "Title":
			message.Title = value[0]
		case "From":
			message.From = value[0]
		case "Uri":
			message.Uri = value[0]
		case "Command":
			message.Command = value[0]
		}
	}

	messageBytes, e = ioutil.ReadAll(from)
	message.Message = string(messageBytes)

out:
	if e != nil {
		message.Result = FAILURE
		message.ErrorString = e.Error()
	}
	return
}
