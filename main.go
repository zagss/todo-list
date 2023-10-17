package main

import (
	"github.com/zagss/todo-list/cmd"
	"github.com/zagss/todo-list/conf"
	"github.com/zagss/todo-list/internal/database"
)

func init() {
	conf.Setup()
	database.Setup()
}

func main() {
	cmd.RunServer()
}
