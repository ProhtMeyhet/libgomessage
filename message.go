package libgomessage

import(

)

type Message struct {
	Id string //IdInterface
	AnswerTo []string //[]IdInterface

	Priority string

	Message string
	Title string

	Command string

	Result uint8
	ErrorCode uint8
	ErrorString string

	// attachements
	//Attachments []AttachmentInterface



	Uri string
	// used for receivers, not senders
	From string
	FromRemote string
	To []string

	// icon name (freedesktop)
	Icon string
	IconPath string
	IconUri string

	// taken from libnotify
	Delay int32

	// if one wants to set up a service
	ApiKey string
}

