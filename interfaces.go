package libgomessage

import(
	"bufio"
	"io"
	"net"
)

const(
	SUCCESS = 0
	FAILURE = 1

	LOW_PRIORITY =		"LOW"
	NO_PRIORITY =		""
	HIGH_PRIORITY =		"HIGH"
	URGENT_PRIORITY =	"URGENT"

	SERVER_STOP = 1
)

type Authentication struct {
	//Username:Password:Random
}

type MessageInterface interface {
	ConstructMessage() []byte
}

type AuthenticationInterface interface {
	//SetLoginInformation(loginInformation LoginInformationInterface)
	//Acceptable(headers map[string]string) bool

	IsAllowed(connection io.ReadWriter) bool
}

type AuthenticatorInterface interface {
	//IsAuthenticated(reader io.Reader, writer io.Writer) bool
	IsAuthenticated(connection io.ReadWriter) bool
}

// A ConnectionHandler just gets the connection
// and can thus read the way it wants
// like the whole thing at one
// or byte by byte ...
type ConnectionHandlerInterface interface {
	Handle(connection net.Conn) *Message
}

//type AuthenticationCheckInterface interface {
//	Acceptable(
//}

type LoginInformationInterface interface {
	getUser() string
	getPassword() string
	getSalt() string
}

type AttachmentInterface interface {
	GetAttachmentName() string
}

type AttachmentToInterface interface {
	AttachmentInterface
	GetFrom() string
	GetTo() []string
}

type AttachmentIconInterface interface {
	AttachmentInterface
	GetIcon() string
	GetIconPath() string
	GetIconUri() string
}

type AttachmentUriInterface interface {
	AttachmentInterface
	GetUri() string
}

type AttachmentLibNotifyInterface interface {
	AttachmentInterface
	GetDelay() int32
}

type AttachmentApiKeyInterface interface {
	AttachmentInterface
	GetApiKey() string
}

type ToInterface interface {
	AddAddress(address ...string) error
	IsValid(address string) bool
	// return The Original Series
	GetTos() []string
}

type IdInterface interface {
	GetId() string
	SetId(newId string)

	// generate new ID
	GetNewId() string
	// validate given id
	Valid(id string) bool
}

type MessageGeneratorInterface interface {
	GenerateMessage() []byte
}

type ParseInterface interface {
	Parse(from *bufio.Reader) *Message
}

type SendMessageInterface interface {
	GetTo() ToInterface
	GetResult() chan *Message
	/* go */ Send(message *Message, to ToInterface)
}

type RecieveMessageInterface interface {
	StartService() error
	/* go */ Receive()
	GetMessage() chan *Message
	StopService()
}
