package messages

type MessageType string

const (
	BirthMessageType   MessageType = "birth"
	SuicideMessageType MessageType = "suicide"
	DeathMessageType   MessageType = "death"
)
