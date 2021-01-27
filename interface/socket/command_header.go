package socket

import (
	"fmt"
	"strings"
)

// CommandHeader type
type CommandHeader map[string]interface{}

// NewCommandHeadersFromMessage create command header from given message
// the message should: EXEC_TYPE:<ExecType>
func NewCommandHeadersFromMessage(message string) CommandHeader {
	headers := map[string]interface{}{}
	messageSplitted := strings.Split(message, " ")

	for _, headerString := range messageSplitted {
		headerStringSplitted := strings.Split(headerString, ":")
		if len(headerStringSplitted) > 1 {
			headers[headerStringSplitted[0]] = headerStringSplitted[1]
		}
	}

	return headers
}

// CommandHeadersToMessage returns message built from given headers
func CommandHeadersToMessage(headers CommandHeader) string {
	messages := ""

	for key, val := range headers {
		messages = fmt.Sprintf("%s:%s %s", key, val, messages)
	}

	return messages
}
