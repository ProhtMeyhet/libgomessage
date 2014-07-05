package libgomessage

import(

)

type TcpPlain struct {
	TcpSender
	config *TcpSenderConfig
}

func NewTcpPlain(config *TcpSenderConfig) *TcpPlain {
	tcp := &TcpPlain{ }
	tcp.Init(config)
	return tcp
}

func (tcp *TcpPlain) Init(Config *TcpSenderConfig) {
	tcp.config = Config
	tcp.TcpSender.Init(Config)
}

func (tcp *TcpPlain) Send(message *Message, to ToInterface) {
	tcp.sendToAll(to.GetTos(), tcp.constructMessage(message))
}

func (tcp *TcpPlain) sendToAll(to []string, message []byte) {
	for _, to := range to {
		go func(gto string) {
			tcp.send(gto, message)
		}(to)
	}
}

/*func (tcp *TcpPlain) send(to string, message []byte) {
	var connection *net.TCPConn

	tcpAddress, e := net.ResolveTCPAddr(tcp.config.Type, to)
	if e != nil {
		goto out
	}

	connection, e = net.DialTCP(tcp.config.Type,
					nil,
					tcpAddress)
	if e != nil {
		goto out
	}
	defer connection.Close()

	_, e = connection.Write(message)

out:
	if e != nil {
		tcp.result <- &Message{ Result: FAILURE,
					ErrorString: e.Error(),
					To: []string{ to } }
	} else {
		tcp.result <- &Message{ Result: SUCCESS,
					To: []string{ to } }
	}
}*/

func (tcp *TcpPlain) GetTo() ToInterface {
	return NewTcpPlainTo(tcp.config.GetPort())
}

func (tcp *TcpPlain) constructMessage(message *Message) []byte {
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
	return []byte(stringMessage)
}



type TcpPlainConfig struct {
	TcpSenderConfig
}

func NewTcpPlainConfig() *TcpSenderConfig {
	return NewTcpSenderConfig()
}

