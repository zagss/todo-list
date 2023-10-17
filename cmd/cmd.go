package cmd

import (
	"github.com/zagss/todo-list/conf"
	"github.com/zagss/todo-list/internal/router"
)

func RunServer() {
	r := router.New()
	_ = r.Run(conf.ServerSetting.HttpPort)
}
