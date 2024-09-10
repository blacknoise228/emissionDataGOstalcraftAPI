package api

import (
	"net/http"

	"strconv"

	"stalcraftBot/configs"
	_ "stalcraftBot/docs"
	"stalcraftBot/internal/jsWorker"
	"stalcraftBot/internal/logs"
	"stalcraftBot/internal/tgBot"
	"stalcraftBot/internal/timeRes"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title StalcraftAPI Telegram Bot
// @version 1.2.0
// @description Telegram Bot fo getting emission info from StalcraftAPI
// @contact.name blacknoise
// @contact.email blacknoise228@gmail.com
func StartAdminAPI(conf *configs.Config) {
	port := ":" + strconv.Itoa(conf.PortAdminAPI)

	routerAPI := gin.Default()
	v1 := routerAPI.Group("")
	{
		users := v1.Group("/users")
		{
			users.GET("", getUsers)
			users.GET(":id", getUser)
			users.POST("", addUser)
			users.DELETE(":id", deleteUser)
			users.DELETE("block/:id", deleteUserFromBlackList)
			users.GET("block/:id", addUserToBlackList)
		}
	}
	routerAPI.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	routerAPI.Run(port)
}
func DataMessageAPI(conf *configs.Config) {
	port := ":" + strconv.Itoa(conf.PortTgBot)

	routerBot := gin.Default()

	routerBot.POST("/emdata", sendEmissionMsgFromAPI)
	routerBot.Run(port)
	logs.Logger.Info().Msg("API for sending data to tgBot started")
}

// @summary Receives a command to start sending messages about the start of emission
// @success 200
// @accept json
// @produse json
// @param b body string false "emData"
// @router /emdata [post]
func sendEmissionMsgFromAPI(ctx *gin.Context) {
	var newEmission jsWorker.EmissionInfo
	if err := ctx.ShouldBindJSON(&newEmission); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	data, err := timeRes.CurrentEmissionResult(newEmission)
	if err != nil {
		logs.Logger.Err(err).Msg("send current emission info to tgAPI error")
	}
	lastEm, err := timeRes.TimeResult(newEmission)
	if err != nil {
		logs.Logger.Err(err).Msg("send current emission info to tgAPI error")
	}
	tgBot.SendMessageTG(data + lastEm)
}

// @summary Retrives all users
// @produce json
// @success 200 {object} jsWorker.User
// @router /users [get]
func getUsers(ctx *gin.Context) {
	jsWorker.LoadChatID()
	ctx.JSON(http.StatusOK, gin.H{"users": jsWorker.Users})
}

// @summary Retrieves user based on given ID
// @produce json
// @param id path integer true "User ID"
// @success 200 {object} jsWorker.User
// @router /users/{id} [get]
func getUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		logs.Logger.Err(err).Msg("id user not int!")
	}
	jsWorker.LoadChatID()
	ctx.JSON(http.StatusOK, gin.H{"user": jsWorker.Users[id-1]})
}

// @summary Create new user
// @produce json
// @param id path integer true "User ID"
// @success 200
// @router /users/{id} [post]
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

// @summary Delete user based on given ID
// @produce json
// @param id path integer true "User ID"
// @success 200
// @router /users/{id} [delete]
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

// @summary Add user to blacklist based on given ID
// @produce json
// @param id path integer true "User ID"
// @success 200
// @router /users/block/{id} [get]
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

// @summary Delete user from blacklist based on given ID
// @produce json
// @param id path integer true "User ID"
// @success 200
// @router /users/block/{id} [delete]
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
