package api

import (
	"net/http"
	"stalcraftbot/configs"
	_ "stalcraftbot/docs"
	"stalcraftbot/internal/jsWorker"
	"stalcraftbot/internal/logs"
	"stalcraftbot/internal/tgBot"
	"stalcraftbot/internal/timeRes"
	"stalcraftbot/pkg/postgres"

	"strconv"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title StalcraftAPI Telegram Bot
// @version 1.2.0
// @description Telegram Bot fo getting emission info from StalcraftAPI
// @contact.name blacknoise
// @contact.email blacknoise228@gmail.com

// Starting API Server with administator tools
func StartAdminAPI(conf *configs.Config) {
	port := ":" + strconv.Itoa(conf.API.AdminAPI.PortAdminAPI)

	routerAPI := gin.Default()
	v1 := routerAPI.Group("")
	{
		users := v1.Group("/users")
		{
			users.GET("", getUsers)
			users.GET(":id", getUser)
			users.POST("", addUser)
			users.DELETE(":id", deleteUser)
		}
	}
	routerAPI.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	routerAPI.Run(port)
}

// Starting API Server for receiving emission data from crawler
func DataMessageAPI(conf *configs.Config) {
	port := ":" + strconv.Itoa(conf.API.BotAPI.PortTgBot)
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
// @success 200 {object} postgres.User
// @router /users [get]
func getUsers(ctx *gin.Context) {
	db := postgres.InitDB()
	postgres.LoadChatID(db)
	ctx.JSON(http.StatusOK, gin.H{"users": postgres.Users})
}

// @summary Retrieves user based on given ID
// @produce json
// @param id path integer true "User ID"
// @success 200 {object} postgres.User
// @router /users/{id} [get]
func getUser(ctx *gin.Context) {
	db := postgres.InitDB()
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		logs.Logger.Err(err).Msg("id user not int!")
	}
	postgres.LoadChatID(db)
	ctx.JSON(http.StatusOK, gin.H{"user": postgres.Users[id-1]})
}

// @summary Create new user
// @produce json
// @param id path integer true "User ID"
// @success 200
// @router /users/{id} [post]
func addUser(ctx *gin.Context) {
	db := postgres.InitDB()
	var newUser postgres.User
	if err := ctx.ShouldBindJSON(&newUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	postgres.SaveChatID(db, newUser)
	ctx.JSON(http.StatusOK, gin.H{"message": "user created"})
}

// @summary Delete user based on given ID
// @produce json
// @param id path integer true "User ID"
// @success 200
// @router /users/{id} [delete]
func deleteUser(ctx *gin.Context) {
	db := postgres.InitDB()
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		logs.Logger.Err(err).Msg("id user not int!")
	}
	postgres.LoadChatID(db)
	for _, val := range postgres.Users {
		if val.ID == id {

			postgres.DeleteUserInDB(db, id)
			ctx.JSON(http.StatusOK, gin.H{"message": "User deleted"})
			return
		}
	}
	ctx.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
}
