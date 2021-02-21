package grpc

import "github.com/luobin998877/go_server/router"

type handler struct{}

func (h *handler) HandleCmd(node string, socketID uint32, ipAddr string, cmd uint32, data []byte) (code uint32, rData []byte) {
	code, rData = router.StartCmdRouter(node, socketID, ipAddr, cmd, data)
	return
}

func (h *handler) HandleJSON(path string, data []byte) (rData []byte) {
	rData = []byte{123, 34, 99, 111, 100, 101, 34, 58, 49, 44, 32, 34, 109, 115, 103, 34, 58, 34, 115, 121, 115, 32, 101, 114, 114, 111, 114, 34, 125}
	return
}
