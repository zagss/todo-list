package router

import (
	"github.com/gin-gonic/gin"
	"github.com/zagss/todo-list/conf"
	"github.com/zagss/todo-list/internal/router/api/v1"
)

func New() *gin.Engine {
	gin.SetMode(conf.ServerSetting.RunMode)
	r := gin.Default()

	apiV1 := r.Group("/api/v1")
	{
		apiV1.GET("ping", v1.Ping)
		apiV1.POST("user/register", v1.UserRegisterHandler)
		apiV1.POST("user/login", v1.UserLoginHandler)
	}

	return r
}
