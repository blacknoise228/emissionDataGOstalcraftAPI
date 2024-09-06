package api

import (
	"net/http"
	"stalcraftBot/internal/jsWorker"
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
	routerAPI.DELETE("/users/block/:id", deleteUserFromBlackList)
	routerAPI.GET("/users/block/:id", addUserToBlackList)
	routerAPI.Run()
}
func getUsers(ctx *gin.Context) {
	jsWorker.LoadChatID()
	ctx.JSON(http.StatusOK, gin.H{"users": jsWorker.Users})
}
func getUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		logs.Logger.Err(err).Msg("id user not int!")
	}
	jsWorker.LoadChatID()
	ctx.JSON(http.StatusOK, gin.H{"user": jsWorker.Users[id-1]})
}
func addUser(ctx *gin.Context) {
	var newUser jsWorker.User
	if err := ctx.ShouldBindJSON(&newUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	jsWorker.Users = append(jsWorker.Users, newUser)
	jsWorker.SaveChatID()
	ctx.JSON(http.StatusOK, gin.H{"message": "user created"})
}
func deleteUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		logs.Logger.Err(err).Msg("id user not int!")
	}
	jsWorker.LoadChatID()
	for i, val := range jsWorker.Users {
		if val.ID == id {
			jsWorker.Users = append(jsWorker.Users[:i], jsWorker.Users[i+1:]...)
			jsWorker.SaveChatID()
			ctx.JSON(http.StatusOK, gin.H{"message": "User deleted"})
			return
		}
	}
	ctx.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
}
func addUserToBlackList(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		logs.Logger.Err(err).Msg("id user not int!")
	}
	jsWorker.LoadChatID()

	for i := range jsWorker.Users {
		if jsWorker.Users[i].ID == id {
			jsWorker.Users[i].Blocked = true
			jsWorker.SaveChatID()
			ctx.JSON(http.StatusOK, gin.H{"message": "User is blocked"})
			return
		}
	}
	ctx.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
}
func deleteUserFromBlackList(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		logs.Logger.Err(err).Msg("id user not int!")
	}
	jsWorker.LoadChatID()

	for i := range jsWorker.Users {
		if jsWorker.Users[i].ID == id {
			jsWorker.Users[i].Blocked = false
			jsWorker.SaveChatID()
			ctx.JSON(http.StatusOK, gin.H{"message": "User is unlocked"})
			return
		}
	}
	ctx.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
}
