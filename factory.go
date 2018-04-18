package messages

import (
	"encoding/json"
	"fmt"
)

type MessageFactory struct {
}

func New() MessageFactory {
	return MessageFactory{}
}

func (mf MessageFactory) CreateBirthMessage(name string, address string) Message {
	data := CreateBirthData(name, address)
	json, err := data.MarshalJSON()
	if err != nil {
		fmt.Println(err)
	}
	return Message{Type: BirthMessageType, Data: json}
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
