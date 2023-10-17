package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong!")
}
