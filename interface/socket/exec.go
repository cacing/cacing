package socket

import (
	"fmt"
	"strings"
)

// ExecType type of execution command
type ExecType string

// const for execution types
const (
	ExecSet ExecType = "EXEC_SET"
	ExecGet          = "EXEC_GET"
	ExecDel          = "EXEC_DEL"
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
	}

	args := make([]string, 0)
	if len(payloadSplitted) > 1 {
		args = payloadSplitted[1:]
	}

	fmt.Println(len(payloadSplitted))

	return &Exec{
		execType,
		args,
	}
}
