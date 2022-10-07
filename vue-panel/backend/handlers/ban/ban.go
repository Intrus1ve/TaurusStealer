package ban

import (
	"net/http"

	database "../../db"
	lggr "../../logger"
	login "../login"

	"github.com/gin-gonic/gin"
)

var db = database.Connect()
var logger = lggr.Log()

type BannedData struct {
	Id   int    `db:"Id" json:"id"`
	Data string `db:"Data" json:"data"`
}

func GetBannedData(ctx *gin.Context) {
	if !login.CtxRootAuthCheck(ctx) {
		logger.Println("ban.GetBannedData: unauthorized user: ", ctx.ClientIP())
		return
	}

	var bannedData []BannedData
	err := db.Select(&bannedData, "SELECT * FROM `banned`")
	if err != nil {
		logger.Println("ban.GetBannedData: get banned data from db error: ", err)
	}

	if len(bannedData) < 1 {
		ctx.JSON(http.StatusOK, "")
		return
	}

	ctx.JSON(http.StatusOK, bannedData)
}

type DdRuleAction struct {
	Cookie string     `json:"cookie"`
	Data   BannedData `json:"form"`
}

func Create(ctx *gin.Context) {
	var rule DdRuleAction
	ctx.BindJSON(&rule)

	if !login.RootAuthCheck(rule.Cookie) {
		logger.Println("ban.Create: unauthorized user: ", ctx.ClientIP())
		return
	}

	_, err := db.Exec("INSERT INTO `banned` (`data`) VALUES (?)",
		rule.Data.Data,
	)

	if err != nil {
		logger.Println("ban.Create error: ", err)
	}

	statusErr := ""
	if err != nil {
		statusErr = err.Error()
	}

	ctx.JSON(http.StatusOK, gin.H{"err": statusErr})
}

func Edit(ctx *gin.Context) {
	var rule DdRuleAction
	ctx.BindJSON(&rule)

	if !login.RootAuthCheck(rule.Cookie) {
		logger.Println("ban.Edit: unauthorized user: ", ctx.ClientIP())
		return
	}

	_, err := db.Exec("UPDATE `banned` SET `data` = ? WHERE `Id` = ?",
		rule.Data.Data,
		rule.Data.Id,
	)

	if err != nil {
		logger.Println("ban.Edit error: ", err)
	}

	statusErr := ""
	if err != nil {
		statusErr = err.Error()
	}

	ctx.JSON(http.StatusOK, gin.H{"err": statusErr})
}

func Delete(ctx *gin.Context) {
	var rule DdRuleAction
	ctx.BindJSON(&rule)

	if !login.RootAuthCheck(rule.Cookie) {
		logger.Println("ban.Delete: unauthorized user: ", ctx.ClientIP())
		return
	}

	_, err := db.Exec("DELETE FROM `banned` WHERE `Id` = ?", rule.Data.Id)
	if err != nil {
		logger.Println("ban.BannedDelete error: ", err)
	}

	statusErr := ""
	if err != nil {
		statusErr = err.Error()
	}

	ctx.JSON(http.StatusOK, gin.H{"err": statusErr})
}
