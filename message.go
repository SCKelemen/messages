package messages

type Message struct {
	Type MessageType `json:"type"`
	Data []byte      `json:"data"`
}
