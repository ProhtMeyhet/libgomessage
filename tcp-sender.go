package libgomessage

import(
	"errors"
	"crypto/tls"
	"net"
	"strconv"
)

type TcpSender struct {
	abstractSender
	config *TcpSenderConfig

	requireAuthentication bool
	authenticator AuthenticatorInterface
}

func NewTcpSender(config *TcpSenderConfig) *TcpSender {
	tcp := &TcpSender{ }
	tcp.Init(config)
	return tcp
}

func (tcp *TcpSender) Init(Config *TcpSenderConfig) {
	tcp.config = Config
	tcp.abstractSender.InitAbstract()
}

func (tcp *TcpSender) RequireAuthentication(Authenticator AuthenticatorInterface) {
	tcp.requireAuthentication = true
	tcp.authenticator = Authenticator
}

func (tcp *TcpSender) Send(generator MessageGeneratorInterface, to ToInterface) {
	tcp.sendToAll(to.GetTos(), generator.GenerateMessage())
}

func (tcp *TcpSender) sendToAll(to []string, message []byte) {
	for _, to := range to {
		go func(gto string) {
			tcp.send(gto, message)
		}(to)
	}
}

func (tcp *TcpSender) dial(address string) (net.Conn, error) {
	var connection net.Conn
	var e error
	if tcp.config.SSL {
		connection, e = tls.Dial(tcp.config.Type,
					 address, &tcp.config.SSLConfig)
	} else {
		tcpAddress, e := net.ResolveTCPAddr(tcp.config.Type, address)
		if e != nil {
			goto out
		}
		connection, e = net.DialTCP(tcp.config.Type,
                                             nil,
                                             tcpAddress)
	}

out:
	return connection, e
}

func (tcp *TcpSender) send(to string, message []byte) {
	connection, e := tcp.dial(to)

	if e != nil {
		goto out
	}
	defer connection.Close()

	if tcp.requireAuthentication {
		if !tcp.authenticator.IsAuthenticated(connection) {
			e = errors.New("Couldn't Authenticate!")
			goto out
		}
	}

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
}

func (tcp *TcpSender) GetName() string {
	return "TcpSender"
}




type TcpSenderConfig struct {
	Type		string // ipv4, ipv6
	Port		int
	SSL		bool
	SSLConfig	tls.Config
}

func NewTcpSenderConfig() *TcpSenderConfig {
	return &TcpSenderConfig{ Type: "tcp4", Port: 65222 }
}

func (tcp *TcpSenderConfig) SetSSLConfig(config tls.Config) {
	tcp.SSL = true
	tcp.SSLConfig = config
}

func (tcp *TcpSenderConfig) GetPort() string {
	return strconv.Itoa(int(tcp.Port))
}

// TODO
type TcpSenderTo struct {
	To
}

func NewTcpSenderTo() *TcpSenderTo {
	return &TcpSenderTo{}
}

func (tcp *TcpSenderTo) AddAddress(address string) error {
	if tcp.IsValid(address) {
		tcp.To.To = append(tcp.To.To, address)
	} else {
		return errors.New("invalid address: " + address)
	}

	return nil
}

func (tcp *TcpSenderTo) IsValid(address string) bool {
	return true
}
