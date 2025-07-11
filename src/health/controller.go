package health

import (
	"github.com/gin-gonic/gin"
)

func NewHealthController(server *gin.Engine) {

	routes := server.Group("/health")
	{
		routes.POST("", health)
		routes.GET("", health)
		routes.HEAD("", health)
	}

}
