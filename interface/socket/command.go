package socket

import (
	"fmt"
	"strings"
)

// Command struct for client and server data
type Command struct {
	Type    Signal
	User    string
	Payload string
	Headers CommandHeader
}

// NewCommandFromMessage create command from message string
// message format: cmd=>uid=>payload=>header
func NewCommandFromMessage(message string) *Command {
	messageSplitted := strings.Split(strings.ReplaceAll(message, "\n", ""), "=>")
	var signalType Signal = Signal(messageSplitted[0])

	user := messageSplitted[1]

	payload := ""
	if len(messageSplitted) > 2 {
		payload = messageSplitted[2]
	}

	headers := CommandHeader{}
	if len(messageSplitted) > 3 {
		headers = NewCommandHeadersFromMessage(messageSplitted[3])
	}

	return &Command{
		signalType,
		user,
		payload,
		headers,
	}
}

// CommandToMessage generate message string from given command
func CommandToMessage(command *Command) string {
	message := fmt.Sprintf("%s=>%s=>%s=>%s\n", command.Type, command.User, command.Payload, CommandHeadersToMessage(command.Headers))
	return message
}
