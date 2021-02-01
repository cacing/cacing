package socket

import (
	"encoding/json"
)

// Command struct for client and server data
type Command struct {
	Type    Signal        `json:"t"`
	User    string        `json:"u"`
	Payload []string      `json:"p,omitempty"`
	Headers CommandHeader `json:"h,omitempty"`
}

// NewCommandFromMessage create command from message string
// message format: cmd=>uid=>payload=>header
func NewCommandFromMessage(message []byte) (*Command, error) {
	command := &Command{}
	err := json.Unmarshal(message, command)
	if err != nil {
		return nil, err
	}

	return command, nil
}

// CommandToMessage generate message string from given command
func CommandToMessage(command *Command) ([]byte, error) {
	message, err := json.Marshal(command)
	if err != nil {
		return nil, err
	}
	return message, nil
}
