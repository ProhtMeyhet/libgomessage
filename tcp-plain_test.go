package libgomessage

import(
	"testing"
	"runtime"
	"time"
	"crypto/tls"
)

func genericSendTest(t *testing.T, client SendMessageInterface, name string) {
	select {
	case result := <-client.GetResult():
		if result.Result == FAILURE {
			t.Errorf(name + " failed! %s", result.ErrorString)
		}
	case <-time.After(time.Duration(20)*time.Second):
		t.Errorf(name + " timeout!")
	}
}

func genericServerTest(t *testing.T, server RecieveMessageInterface,
			expected *Message, name string) {
	select {
	case message := <-server.GetMessage():
		if message.Result == FAILURE {
			t.Errorf("server error: %s from %s", message.ErrorString, name)
		}
		if message.Title != expected.Title {
			t.Errorf("Got Title '%s', want '%s'", message.Title,
					expected.Title)
		}
		if message.Message != expected.Message {
			t.Errorf("Got Message '%s', want '%s'", message.Message,
					expected.Message)
		}
	case <-time.After(time.Duration(5) * time.Second):
		t.Errorf(name + " timeout!")
	}
}

func TestTcp(t *testing.T) {
	runtime.GOMAXPROCS(10)
	serverConfig := NewTcpServerConfig()
	serverConfig.SetSSL("/home/neo/go/bin/cert.pem", "/home/neo/go/bin/key.pem")
	server := NewTcpPlainServer(serverConfig)

	e := server.StartService()
	if e != nil {
		t.Errorf("couldn't start service: %s", e.Error())
		return
	}
	server.RequireAuthentication(NewUserPassAuthentication())

	authenticator := NewUserPassAuthenticator()
	authenticator.User = "Me"
	authenticator.Pass = "asdf"

	tlsConfig := tls.Config{ InsecureSkipVerify: true }
	clientConfig := NewTcpPlainConfig()
	clientConfig.SetSSLConfig(tlsConfig)
	client := NewTcpPlain(clientConfig)
	client.RequireAuthentication(authenticator)

	message := &Message{ Title: "Kenny is dead!",
				Message: "Oh, my God! They killed Kenny." }
	to := client.GetTo()
	to.AddAddress("127.0.0.1")

	go server.Receive()

	// the go command has some overhead, so sleep for 1 second
	time.Sleep(time.Duration(1) * time.Second)

	go client.Send(message, to)
	go genericSendTest(t, client, "TcpPlain")
	genericServerTest(t, server, message, "TcpPlain")
	server.StopService()
}
