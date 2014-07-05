package libgomessage

import(
	"errors"
)


// TODO
type TcpPlainTo struct {
	To
	port string
}

func NewTcpPlainTo(Port string) *TcpPlainTo {
	return &TcpPlainTo{ port: Port }
}

func (tcp *TcpPlainTo) AddAddress(addresses ...string) error {
	for _, address := range addresses {
		if tcp.IsValid(address) {
			tcp.To.To = append(tcp.To.To, address + ":" + tcp.port)
		} else {
			return errors.New("invalid address: " + address)
		}
	}
	return nil
}

func (tcp *TcpPlainTo) IsValid(address string) bool {
	return true
}
