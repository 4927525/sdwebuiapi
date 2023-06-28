package server

import (
	"sdwebuiapi/config"
	"sdwebuiapi/crontab"
	"sdwebuiapi/router"
)

var App = &Server{}

type Server struct {
}

func (*Server) Start() {
	config.Init()
	go func() {
		crontab.Init()
	}()
	engine := router.Router()
	_ = engine.Run(config.Config.GetString("server.port"))
}
