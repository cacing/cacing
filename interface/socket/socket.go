package socket

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

// ExecType type of execution command
type ExecType string

// const for execution types
const (
	ExecSet    ExecType = "EXEC_SET"
	ExecGet             = "EXEC_GET"
	ExecDel             = "EXEC_DEL"
	ExecExp             = "EXEC_EXP"
	ExecExists          = "EXEC_EXISTS"
)

// const for command header
const (
	CommandHeaderExecType ExecType = "EXEC_TYPE"
)
