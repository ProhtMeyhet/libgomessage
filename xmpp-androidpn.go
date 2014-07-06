package libgomessage

import(

)

/*
send messages to androidpn:
https://github.com/dannytiehui/androidpn

the message format:

<notification namespace="androidpn:iq:notification">
	<id></id>
	<apiKey></apiKey>
	<title></title>
	<message></message>
	<uri></uri>
</notification>
*/

type XmppAndroidpn struct {
	TcpPlain
}

func NewXmppAndroidpn(config *TcpSenderConfig) *XmppAndroidpn {
	xmpp := &XmppAndroidpn{ }
	xmpp.TcpPlain.Init(config)
	return xmpp
}

func (Xmpp *XmppAndroidpn) Send(message *Message, to ToInterface) {
	Xmpp.sendToAll(to.GetTos(), Xmpp.constructXmpp(message))
}

func (Xmpp *XmppAndroidpn) GetTo() ToInterface {
	return NewTcpPlainTo(Xmpp.config.GetPort())
}

func (Xmpp *XmppAndroidpn) constructXmpp(message *Message) []byte {
	xmpp := "<Notification namespace=\"androidpn:iq:notification\">"
	xmpp += "<Id>" + message.Id + "</Id>"
/*	for _, attachment := range message.Attachments {
		if x, ok := attachment.(AttachmentApiKeyInterface);  ok {
		        xmpp += "<ApiKey>" + x.GetApiKey() + "</ApiKey>"
		}
		if x, ok := attachment.(AttachmentUriInterface); ok {
		        xmpp += "<Uri>" + x.GetUri() + "</Uri>"
		}
	}*/
        xmpp += "<Title>" + message.Title + "</Title>"
        xmpp += "<Message>" + message.Message + "</Message>"
	xmpp += "</Notification>"
	return []byte(xmpp)
}
