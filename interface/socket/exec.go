package socket

import (
	"strings"
)

// Exec define execution model
type Exec struct {
	Type ExecType
	Args []string
}

// NewExecFromCommandPayload create exec using command payload
func NewExecFromCommandPayload(payload string) *Exec {
	payloadSplitted := strings.Split(payload, " ")
	var execType ExecType

	switch strings.ToLower(payloadSplitted[0]) {
	case "set":
		execType = ExecSet
	case "get":
		execType = ExecGet
	case "del":
		execType = ExecDel
	case "exp":
		execType = ExecExp
	case "exists":
		execType = ExecExists
	}

	args := make([]string, 0)
	if len(payloadSplitted) > 1 {
		args = payloadSplitted[1:]
	}

	return &Exec{
		execType,
		args,
	}
}
