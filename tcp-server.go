package libgomessage

import(
	"errors"
	"net"
	"strconv"
	"time"
        "crypto/tls"
)

type TcpServer struct {
	abstractServer
	config *TcpServerConfig
	tcpAddr *net.TCPAddr
	listener net.Listener
	acceptError error
	connectionHandler ConnectionHandlerInterface

	requireAuthentication bool
	authentication AuthenticationInterface
}

func (tcp *TcpServer) Init(config *TcpServerConfig, handler ConnectionHandlerInterface) {
	tcp.config = config
	tcp.connectionHandler = handler
	tcp.InitAbstract(config.MaxConnections)
}

func (tcp *TcpServer) RequireAuthentication(Authentication AuthenticationInterface) {
	tcp.requireAuthentication = true
	tcp.authentication = Authentication
}

func (tcp *TcpServer) StartService() error {
	if tcp.config.SSL {
		return tcp.startSSLService()
	}

	if e := tcp.config.Valid(); e != nil {
		return e
	}

	serviceString := ":" + tcp.config.GetPort()
	tcpAddr, e := net.ResolveTCPAddr(tcp.config.Type, serviceString)
	if e != nil {
		return e
	}
	tcp.tcpAddr = tcpAddr

	listener, e := net.ListenTCP("tcp", tcpAddr)
	if e != nil {
		return e
	}
	tcp.listener = listener

	return nil
}

func (tcp *TcpServer) startSSLService() error {
	if e := tcp.config.Valid(); e != nil {
		return e
	}

	cert, e := tls.LoadX509KeyPair(tcp.config.CertificateFile, tcp.config.KeyFile)
        if e != nil {
		return e
        }

	config := tls.Config{Certificates: []tls.Certificate{cert}}

	serviceString := ":" + tcp.config.GetPort()
	/*tcpAddr, e := net.ResolveTCPAddr(tcp.config.Type, serviceString)
	if e != nil {
		return e
	}
	tcp.tcpAddr = tcpAddr*/

	listener, e := tls.Listen("tcp", serviceString, &config)
	if e != nil {
		return e
	}
	tcp.listener = listener

	return nil
}

func (tcp *TcpServer) StopService() {
	tcp.control <- SERVER_STOP
}

func (tcp *TcpServer) accept(connectionChannel chan net.Conn) chan net.Conn {
	var tempDelay time.Duration

	connection, e := tcp.listener.Accept()

	if e != nil {
		if ne, ok := e.(net.Error); ok && ne.Temporary() {
			if tempDelay == 0 {
				tempDelay = 5 * time.Millisecond
			} else {
				tempDelay *= 2
			}
			if max := 1 * time.Second; tempDelay > max {
				tempDelay = max
			}
			time.Sleep(tempDelay)
		}
		goto out
		//TODO
		//tcp.acceptError = e
		//break infinite
	}

	connectionChannel <-connection

out:
	return connectionChannel
}

func (tcp *TcpServer) Receive() {
	connectionChannel := make(chan net.Conn, tcp.config.MaxConnections)
	connectionTempChannel := make(chan net.Conn, tcp.config.MaxConnections)

	//start handle threads
	for i := 0; i < tcp.config.MaxConnections; i++ {
		go tcp.handleConnection(connectionChannel)
	}

infinite:
	for {
		select {
		case command := <-tcp.control:
			tcp.control <-command
			if command == SERVER_STOP {
				//TODO
				//wait for go threads to finish
				break infinite
			}
		case connection := <-tcp.accept(connectionTempChannel):
			connectionChannel <-connection
		}
	}
}

func (tcp *TcpServer) handleConnection(connectionChannel chan net.Conn) {
infinite:
	for {
		select {
		case command := <-tcp.control:
			tcp.control <-command
			if command == SERVER_STOP {
				break infinite
			}
		case connection := <-connectionChannel:
			if tcp.acceptError != nil {
				tcp.messages <- &Message{ Result: FAILURE,
						ErrorString: tcp.acceptError.Error() }
				break infinite
			}

			if tcp.requireAuthentication {
				if !tcp.authentication.IsAllowed(connection) {
					connection.Close()
					continue
				}
			}

			tcp.messages <- tcp.connectionHandler.Handle(connection)
		}
	}
}

type TcpServerConfig struct {
	Type		string
	Port		int
	MaxConnections	int
	ReadTimeOut	uint32
	SSL		bool
	CertificateFile	string
	KeyFile		string
}

func NewTcpServerConfig() *TcpServerConfig {
	return &TcpServerConfig{ Port: 65222,
					MaxConnections: 20,
					Type: "tcp",
					ReadTimeOut: 30 }
}

func (tcp *TcpServerConfig) Valid() error {
	if tcp.Port < 1 {
		return errors.New("Port is smaller than 1!")
	}

	if tcp.MaxConnections < 1 {
		return errors.New("MaxConnections is smaller than 1!")
	}

	if tcp.ReadTimeOut < 1 {
		return errors.New("ReadTimeOut is smaller than 1!")
	}

	return nil
}

func (tcp *TcpServerConfig) GetPort() string {
	return strconv.Itoa(int(tcp.Port))
}

func(tcp *TcpServerConfig) SetSSL(certificateFile, keyFile string) {
	tcp.SSL = true
	tcp.CertificateFile = certificateFile
	tcp.KeyFile = keyFile
}
