package messages

type Message struct {
	Type MessageType `json:"type"`
	Data interface{} `json:"data"`
}
