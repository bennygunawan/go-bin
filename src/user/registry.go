package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewUserController(server *gin.Engine) {

	routes := server.Group("/user")
	{
		routes.POST("", health)
		routes.GET("", health)
		routes.HEAD("", health)
	}

}

func health(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, ".")
}
