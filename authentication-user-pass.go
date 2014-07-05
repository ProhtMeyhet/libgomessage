package libgomessage

import(
	"bytes"
	"io"
)

var	ALLOWED =	[]byte("y")
var	NOT_ALLOWED =	[]byte("n")




type UserPassAuthentication struct {
	//database UserPasswordDatabaseInterface
}

func NewUserPassAuthentication() *UserPassAuthentication {
	return &UserPassAuthentication{ }
}

func (userPass *UserPassAuthentication) IsAllowed(connection io.ReadWriter) bool {
	userBuffer := make([]byte, 6)

	if _, e := connection.Read(userBuffer); e == nil {
		if bytes.Equal(userBuffer, []byte("User: ")) {
			go connection.Write(ALLOWED)
			return true
		}
	}

	go connection.Write(NOT_ALLOWED)

	return false
}



type UserPassAuthenticator struct {
	User string
	Pass string

	tokenUser []byte
	tokenPass []byte
	}

func NewUserPassAuthenticator() *UserPassAuthenticator {
	return &UserPassAuthenticator{ tokenUser: []byte("User: "),
					tokenPass: []byte("Pass: ") }
}

func (userPass *UserPassAuthenticator) IsAuthenticated(connection io.ReadWriter) bool {
	connection.Write(userPass.tokenUser)

	resultBuffer := make([]byte, 1)

	if _, e := connection.Read(resultBuffer); e == nil {
		if bytes.Equal(resultBuffer, ALLOWED) {
			return true
		}
	}

	return false
}
