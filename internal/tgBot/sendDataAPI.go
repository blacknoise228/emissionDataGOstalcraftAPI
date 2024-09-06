package tgBot

import (
	"net/http"
	"stalcraftBot/internal/emissionInfo"
	"stalcraftBot/internal/jsWorker"
	"stalcraftBot/internal/logs"
	"time"

	"github.com/gin-gonic/gin"
)

func DataMessageAPI() {
	routerBot := gin.Default()
	routerBot.GET("/emdata", sendEmissionMsg)
	routerBot.Run(":1234")
	logs.Logger.Info().Msg("API for sending data to tgBot started")
}
func sendEmissionMsg(ctx *gin.Context) {
	info := jsWorker.LoadEmData(emissionInfo.CurrentEmissionDataFile)
	SendMessageTG(info)
	ctx.JSON(http.StatusOK, "")
	time.Sleep(3 * time.Minute)
	SendMessageTG("Еще немного и можно будет собирать артефакты!")
}
