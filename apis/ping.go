package apis

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func PingApi(c *gin.Context) {
	c.JSON(
		http.StatusOK,
		gin.H{
			"code":  http.StatusOK,
			"msg": "Welcome server",
		},
	)
}
