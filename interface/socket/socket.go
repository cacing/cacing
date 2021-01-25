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
