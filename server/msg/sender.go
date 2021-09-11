package msg

import "time"

type Sender struct {
	UUId string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
}

func (s *Sender) CreateMessage(contentType ContentType, contentData []byte, groupName string, mentionedNames ...string) Message {
	return Message{
		&message{
			Sender: *s,
			Content: Content{
				Data: contentData,
				Type: contentType,
			},
			GroupName:      groupName,
			MentionedNames: mentionedNames,
			SendTime:       time.Now().Local(),
		},
	}
}
