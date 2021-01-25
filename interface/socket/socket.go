package socket

import (
	"strings"
)

// const for connection
const (
	ConnHost = "localhost"
	ConnType = "tcp"
)

// Signal type
type Signal string

// const for signal type
const (
	SignalConnect Signal = "SIGNAL_CONNECT"
	SignalExec           = "SIGNAL_EXEC"
	SignalSuccess        = "SIGNAL_SUCCESS"
	SignalError          = "SIGNAL_ERROR"
)

// Command struct for client and server data
type Command struct {
	Type    Signal
	User    string
	Payload string
}

// NewCommandFromMessage create command from message string
// message format: cmd=>uid=>payload
func NewCommandFromMessage(message string) *Command {
	messageSplitted := strings.Split(message, "=>")
	var signalType Signal = SignalError

	switch strings.ToLower(messageSplitted[0]) {
	case "connect":
		signalType = SignalConnect
	case "exec":
		signalType = SignalExec
	}

	user := messageSplitted[1]

	payload := ""
	if len(messageSplitted) > 2 {
		payload = messageSplitted[2]
	}
	return &Command{
		signalType,
		user,
		payload,
	}
}
