package login

import "github.com/luobin998877/go_server/errcode"

// StartRoute cmd
func StartRoute(cmd uint32, data []byte) (code uint32, rData []byte) {
	switch cmd {
	case 1000:
		return singUp(data)
	case 1001:
		return logIn(data)
	default:
		return errcode.SYS, []byte{}
	}
}
