package controllers

import "github.com/vitorsavian/warsong-repo/elminster/pkg/server/connections/tcp"

type StateController struct {
	Database  *DataseController
	TcpConfig *tcp.Config
}

func CreateStateController() *StateController {

}
