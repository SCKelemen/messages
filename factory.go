package messages

import (
	"encoding/json"
)

type MessageFactory struct {
}

func New() MessageFactory {
	return MessageFactory{}
}

func (mf MessageFactory) CreateBirthMessage(name string, address string) Message {
	data := CreateBirthData(name, address)
	return Message{Type: BirthMessageType, Data: data}
}

func (mf MessageFactory) CreateDeathMessage(name string) Message {
	return Message{Type: DeathMessageType}
}

func (mf MessageFactory) CreateSuicideMessage(name string) Message {
	return Message{Type: SuicideMessageType}
}

type BirthData struct {
	Name    string `json:"name"`
	Address string `json:"addr"`
}

func CreateBirthData(name string, address string) BirthData {
	return BirthData{Name: name, Address: address}
}
func (bd BirthData) MarshalJSON() ([]byte, error) {
	return json.Marshal(bd)
}
func (bd BirthData) UnmarshalJSON(data []byte) error {
	//data := BirthData{}
	err := json.Unmarshal(data, &bd)
	return err
}
