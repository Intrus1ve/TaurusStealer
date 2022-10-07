package domaind

import (
	"net/http"

	database "../../db"
	lggr "../../logger"
	login "../login"

	"github.com/gin-gonic/gin"
)

var db = database.Connect()
var logger = lggr.Log()

type DdData struct {
	Id      int    `db:"Id" json:"id"`
	Group   string `db:"Group" json:"group"`
	Color   string `db:"Color" json:"color"`
	Domains string `db:"Domains" json:"domains"`
}

func GetDd() (linksData []DdData) {
	err := db.Select(&linksData, "SELECT * FROM `domain_detect` ORDER BY `Id` DESC")
	if err != nil {
		logger.Println("domaind.GetDd: get domain detect rules from bd error: ", err)
	}
	return linksData
}

func GetDdData(ctx *gin.Context) {
	if !login.CtxRootAuthCheck(ctx) {
		logger.Println("domaind.GetDdData: unauthorized user: ", ctx.ClientIP())
		return
	}

	ddData := GetDd()
	if len(ddData) < 1 {
		ctx.JSON(http.StatusOK, "")
		return
	}

	ctx.JSON(http.StatusOK, ddData)
}

type DdRuleAction struct {
	Cookie string `json:"cookie"`
	Data   DdData `json:"form"`
}

func Create(ctx *gin.Context) {
	var rule DdRuleAction
	ctx.BindJSON(&rule)

	if !login.RootAuthCheck(rule.Cookie) {
		logger.Println("domaind.Create: unauthorized user: ", ctx.ClientIP())
		return
	}

	_, err := db.Exec("INSERT INTO `domain_detect` (`Group`, `Color`, `Domains`) VALUES (?, ?, ?)",
		rule.Data.Group,
		rule.Data.Color,
		rule.Data.Domains,
	)

	if err != nil {
		logger.Println("domaind.Create: ", err)
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
		logger.Println("domaind.Edit: unauthorized user: ", ctx.ClientIP())
		return
	}

	_, err := db.Exec("UPDATE `domain_detect` SET `Group` = ?, `Color` = ?, `Domains` = ? WHERE `Id` = ?",
		rule.Data.Group,
		rule.Data.Color,
		rule.Data.Domains,
		rule.Data.Id,
	)

	if err != nil {
		logger.Println("domaind.Edit: ", err)
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
		logger.Println("domaind.Delete: unauthorized user: ", ctx.ClientIP())
		return
	}

	_, err := db.Exec("DELETE FROM `domain_detect` WHERE `Id` = ?", rule.Data.Id)
	if err != nil {
		logger.Println("domaind.AddDdRule: delete domain detect rule error:", err)
	}

	statusErr := ""
	if err != nil {
		statusErr = err.Error()
	}

	ctx.JSON(http.StatusOK, gin.H{"err": statusErr})
}
