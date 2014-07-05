package libgomessage

import(
	"testing"
	"time"
)

func TestNotify(t *testing.T) {
	notify := NewNotify()
	message := &Message{ Title: "Kenny is dead!",
				Message: "Oh, my God! They killed Kenny." }
	to := &To {}

	go notify.Send(message, to)

	select {
	case result := <-notify.GetResult():
		if result.Result == FAILURE {
			t.Errorf("could't notify! %s", result.ErrorString)
		}
	case <-time.After(time.Duration(20)*time.Second):
		t.Errorf("notify timeout!")
	}
}
