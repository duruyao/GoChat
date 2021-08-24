package server

import (
	"encoding/json"
	"time"
)

type message struct {
	Time    string `json:"time,omitempty"`
	Sender  string `json:"sender,omitempty"`
	Content string `json:"content,omitempty"`
}

type Message struct {
	*message
}

func NewMessage(sender string, content string) *Message {
	return &Message{
		&message{
			Time:    time.Now().Format("2006-01-02 03-04-05"),
			Sender:  sender,
			Content: content},
	}
}

func (m *Message) String() string {
	if js, err := json.MarshalIndent(m.message, "", "    "); err != nil {
		return err.Error()
	} else {
		return string(js)
	}
}

func (m *Message) Serialize() string {
	if js, err := json.Marshal(m.message); err != nil {
		return err.Error()
	} else {
		return string(js)
	}
}

func (m *Message) Parse(js []byte) error {
	if len(js) == 0 {
		return nil
	}
	return json.Unmarshal(js, m.message)
}
