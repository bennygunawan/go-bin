package health

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func health(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, ".")
}