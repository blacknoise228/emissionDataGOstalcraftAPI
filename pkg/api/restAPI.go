package api

import (
	"net/http"
	"stalcraftBot/internal/jSon"

	"github.com/gin-gonic/gin"
)

func StartAdminAPI() {

	routerAPI := gin.Default()
	routerAPI.GET("/users", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"users": jSon.Users})
	})
	routerAPI.Run()
}
