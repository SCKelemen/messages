package messages

type MessageType string

const (
	BirthMessageType   MessageType = "birth"
	SuicideMessageType MessageType = "suicide"
	DeathMessageType   MessageType = "death"
)

type IMessage interface {
	Type() MessageType
	Data() interface{}
}

type BirthMessage struct {
	data BirthData
}

func (bm BirthMessage) Type() MessageType {
	return BirthMessageType
}
func (bm BirthMessage) Data() interface{} {
	return bm.data
}
