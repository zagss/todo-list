package cmd

import (
	"github.com/zagss/todo-list/conf"
	"github.com/zagss/todo-list/router"
)

func RunServer() {
	r := router.New()
	_ = r.Run(conf.HttpPort)
}
