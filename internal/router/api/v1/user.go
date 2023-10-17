package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/zagss/todo-list/internal/types"
)

func UserRegisterHandler(c *gin.Context) {
	var req types.UserServiceReq
	if err := c.ShouldBind(&req); err == nil {

	} else {
		//c.JSON(http.StatusBadRequest, ErrorR)
	}
}

func UserLoginHandler(c *gin.Context) {
	var req types.UserServiceReq
	if err := c.ShouldBind(&req); err == nil {

	}
}
