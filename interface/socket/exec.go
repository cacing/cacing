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
func NewExecFromCommandPayload(payload []string) *Exec {
	var execType ExecType

	switch strings.ToLower(payload[0]) {
	case "set":
		execType = ExecSet
	case "get":
		execType = ExecGet
	case "del":
		execType = ExecDel
	case "exists":
		execType = ExecExists
	}

	args := payload[1:]

	return &Exec{
		execType,
		args,
	}
}
