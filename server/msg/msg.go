package msg

import (
	"encoding/json"
	"time"
)

type ContentType int

const (
	ContentTypeText ContentType = iota
	ContentTypeFile
	ContentTypeImage
	ContentTypeVideo
)

type Content struct {
	Data []byte      `json:"data,omitempty"`
	Type ContentType `json:"type,omitempty"`
}

type message struct {
	Sender         Sender    `json:"sender,omitempty"`
	Content        Content   `json:"content,omitempty"`
	GroupName      string    `json:"group_name,omitempty"`
	MentionedNames []string  `json:"mentioned_names"`
	SendTime       time.Time `json:"send_time,omitempty" time_format:"2006-01-02 15-04-05.000"`
	ForwardTime    time.Time `json:"forward_time,omitempty" time_format:"2006-01-02 15-04-05.000"`
}

type Message struct {
	*message
}

func (m *Message) Sender() Sender {
	return m.message.Sender
}

func (m *Message) Content() Content {
	return m.message.Content
}

func (m *Message) GroupName() string {
	return m.message.GroupName
}

func (m *Message) MentionedNames() []string {
	return m.message.MentionedNames
}

func (m *Message) SendTime() time.Time {
	return m.message.SendTime
}

func (m *Message) ForwardTime() time.Time {
	return m.message.ForwardTime
}

func (m *Message) SetSender(sender Sender) {
	m.message.Sender = sender
}

func (m *Message) SetContent(content Content) {
	m.message.Content = content
}

func (m *Message) SetGroupName(groupName string) {
	m.message.GroupName = groupName
}

func (m *Message) SetMentionedNames(mentionedNames ...string) {
	m.message.MentionedNames = mentionedNames
}

func (m *Message) SetSendTime(sendTime time.Time) {
	m.message.SendTime = sendTime
}

func (m *Message) SetForwardTime(forwardTime time.Time) {
	m.message.ForwardTime = forwardTime
}

func (m *Message) Serialize() ([]byte, error) {
	return json.Marshal(m.message)
}

func (m *Message) Parse(js []byte) error {
	return json.Unmarshal(js, m.message)
}

func (m *Message) String() string {
	if js, err := json.MarshalIndent(m.message, "", "    "); err != nil {
		return err.Error()
	} else {
		return string(js)
	}
}
