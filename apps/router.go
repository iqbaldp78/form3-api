package apps

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func setRouterHandler(router *gin.Engine) {

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	web := router.Group("accounts")
	{
		accountRepo := repository.
			web.GET("/")
	}

}
