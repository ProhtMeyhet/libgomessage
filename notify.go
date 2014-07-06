package libgomessage

import (
	"errors"
	"time"
	notifylib "github.com/mqu/go-notify"
)

type notify struct {
	abstractSender

	delay int32
	icon string
}


func NewNotify() *notify {
	n := &notify{ }
	n.InitAbstract()
	return n
}

func (n *notify) Send(message *Message, to ToInterface) {
	var notifyWindow *notifylib.NotifyNotification
	var e error = nil
	if e = n.verify(message); e != nil {
		goto out
	}

	notifylib.Init(message.Title)
	notifyWindow = notifylib.NotificationNew(message.Title,
				message.Message, message.Icon)

	if notifyWindow == nil {
		e = errors.New("Unable to create a new notification")
		goto out
	}

	notifyWindow.SetTimeout(n.delay)

	if glibe := notifyWindow.Show(); glibe != nil {
		// ignore empty error strings
		if glibe.Error() != "" {
			e = errors.New("Cannot show notification! '" +
					glibe.Error() + "'")
			goto out
		}
	}

out:
	if e != nil {
		n.result <- &Message{ Result: FAILURE, ErrorString: e.Error() }
	} else {
		n.result <- &Message{ Result: SUCCESS }
	}

	time.Sleep(time.Duration(n.delay) * time.Millisecond)
	notifyWindow.Close()
	notifylib.UnInit()
}

func (n *notify) GetTo() ToInterface {
	return NewTo()
}

func (n *notify) verify(message *Message) error {
	if message.Title == "" {
		return errors.New("No title given!")
	}

	if message.Message == "" {
		return errors.New("No message given!")
	}

/*
	for _, attachment := range message.Attachments {
		if x, ok := attachment.(AttachmentLibNotifyInterface); ok {
			n.delay = x.GetDelay();
		}
		if x, ok := attachment.(AttachmentIconInterface); ok {
			n.icon = x.GetIcon()
			if n.icon == "" {
				n.icon = x.GetIconPath()
			}
		}
	}*/

	if n.delay <= 0 {
		//return errors.New("Delay must be greater then 0!")
		n.delay = 10000
	}

	return nil
}
