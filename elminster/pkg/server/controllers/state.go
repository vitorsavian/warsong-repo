package controllers

import "github.com/vitorsavian/warsong-repo/elminster/pkg/server/connections/tcp"

type StateController struct {
	TcpConfig *tcp.Config
}

func CreateStateController() *StateController {
	return nil
}
