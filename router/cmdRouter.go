package router

import (
	"fmt"

	"github.com/luobin998877/go_server/errcode"
	"github.com/luobin998877/go_server/modules/login"
)

// StartCmdRouter router
func StartCmdRouter(node string, socketID uint32, ipAddr string, cmd uint32, data []byte) (code uint32, rData []byte) {
	fmt.Println("node: ", node, "socketID: ", socketID, "ipAddr: ", ipAddr)
	module := findModules(cmd)
	switch module {
	case "login":
		return login.StartRoute(cmd, data)
	default:
		return errcode.SYS, []byte{}
	}
}

func findModules(cmd uint32) string {
	var m uint8
	m = uint8(cmd / 100)
	switch m {
	case 10:
		return "login"
	default:
		return "unknow"
	}
}
