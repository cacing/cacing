package socket

import (
	"fmt"
	"strings"
)

// CommandHeader structure of command header
type CommandHeader struct {
	ExecType ExecType
}

// NewCommandHeadersFromMessage create command header from given message
// the message should: EXEC_TYPE:<ExecType>
func NewCommandHeadersFromMessage(message string) []*CommandHeader {
	messageSplitted := strings.Split(message, " ")

	headers := make([]*CommandHeader, 0)
	for _, headerString := range messageSplitted {
		headerStringSplitted := strings.Split(headerString, ":")
		header := &CommandHeader{}
		if headerStringSplitted[0] == string(CommandHeaderExecType) {
			switch headerStringSplitted[1] {
			case string(ExecSet):
				header.ExecType = ExecSet
			}
		}

		headers = append(headers, header)
	}

	return headers
}

// CommandHeadersToMessage returns message built from given headers
func CommandHeadersToMessage(headers []*CommandHeader) string {
	messages := ""

	for _, header := range headers {
		if string(header.ExecType) != "" {
			headerString := fmt.Sprintf("%s:%s", CommandHeaderExecType, header.ExecType)
			messages = fmt.Sprintf("%s %s", messages, headerString)
		}
	}

	return messages
}
