package libgomessage

import(
	"bufio"
	"encoding/xml"
	"io/ioutil"
)

type XmppAndroidpnServer struct {
	TcpPlainServer
}

func NewXmppAndroidpnServer(config *TcpServerConfig) *XmppAndroidpnServer {
	xmpp := &XmppAndroidpnServer{ }
	xmpp.TcpPlainServer.Init(config, xmpp)
	return xmpp
}

func (xmpp *XmppAndroidpnServer) GetName() string {
	return "XmppAndroidpnServer"
}

func (xmpp *XmppAndroidpnServer) Parse(from *bufio.Reader) (message *Message) {
	messageBytes, e := ioutil.ReadAll(from)
	if e != nil {
		goto out
	}

	e = xml.Unmarshal(messageBytes, &message)

out:
	if e != nil {
		message.Result = FAILURE
		message.ErrorString = e.Error()
	}

	return
}
