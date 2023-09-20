package router

import (
	"github.com/gin-gonic/gin"
	"github.com/zagss/todo-list/conf"
)

func New() *gin.Engine {
	gin.SetMode(conf.RunMode)
	r := gin.Default()
	return r
}
