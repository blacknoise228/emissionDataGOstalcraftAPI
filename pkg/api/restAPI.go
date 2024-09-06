package api

import (
	"net/http"
	"stalcraftBot/internal/jSon"
	"stalcraftBot/internal/logs"
	"strconv"

	"github.com/gin-gonic/gin"
)

func StartAdminAPI() {

	routerAPI := gin.Default()
	routerAPI.GET("/users", getUsers)
	routerAPI.GET("/users/:id", getUser)
	routerAPI.POST("/users", addUser)
	routerAPI.DELETE("/users/:id", deleteUser)
	routerAPI.Run()
}
func getUsers(ctx *gin.Context) {
	jSon.LoadChatID()
	ctx.JSON(http.StatusOK, gin.H{"users": jSon.Users})
}
func getUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		logs.Logger.Err(err).Msg("id user not int!")
	}
	jSon.LoadChatID()
	ctx.JSON(http.StatusOK, gin.H{"user": jSon.Users[id-1]})
}
func addUser(ctx *gin.Context) {
	var newUser jSon.User
	if err := ctx.ShouldBindJSON(&newUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	jSon.Users = append(jSon.Users, newUser)
	jSon.SaveChatID()
	ctx.JSON(http.StatusOK, gin.H{"message": "user created"})
}
func deleteUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		logs.Logger.Err(err).Msg("id user not int!")
	}
	jSon.LoadChatID()
	for i, val := range jSon.Users {
		if val.ID == id {
			jSon.Users = append(jSon.Users[:i], jSon.Users[i+1:]...)
			jSon.SaveChatID()
			ctx.JSON(http.StatusOK, gin.H{"message": "User deleted"})
			return
		}
	}
	ctx.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
}
