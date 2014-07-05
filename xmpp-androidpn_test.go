package libgomessage

import(
	"testing"
	"runtime"
	"time"
)

func TestXmppAndroidpn(t *testing.T) {
	runtime.GOMAXPROCS(10)
	serverConfig := NewTcpServerConfig()
	serverConfig.Port = 65256
	server := NewXmppAndroidpnServer(serverConfig)

	e := server.StartService()
	if e != nil {
		t.Errorf("couldn't start service: %s", e.Error())
		return
	}

	config := NewTcpPlainConfig()
	config.Port = serverConfig.Port
	client := NewXmppAndroidpn(config)
	message := &Message{ Title: "Kenny is dead!",
				Message: "Oh, my God! They killed Kenny." }
	to := client.GetTo()
	to.AddAddress("127.0.0.1")

	go server.Receive()

	// the go command has some overhead, so sleep for 1 second
	time.Sleep(time.Duration(1) * time.Second)

	go client.Send(message, to)
	go genericSendTest(t, client, "XmppAndroidpn")
	genericServerTest(t, server, message, "XmppAndroidpn")
	server.StopService()
}
