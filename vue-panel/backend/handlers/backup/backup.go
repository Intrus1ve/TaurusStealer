package backup

import (
	"archive/zip"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	database "../../db"
	lggr "../../logger"
	utils "../../utils"
	login "../login"

	"github.com/gin-gonic/gin"
)

var db = database.Connect()
var logger = lggr.Log()

type BackupData struct {
	Id      int    `db:"Id" json:"id"`
	Date    string `db:"Date" json:"date"`
	Size    string `db:"Size" json:"size"`
	Comment string `db:"Comment" json:"comment"`
}

func GetBackupsData(ctx *gin.Context) {
	if !login.CtxRootAuthCheck(ctx) {
		logger.Println("backup.GetBackupData: unauthorized user: ", ctx.ClientIP())
		return
	}

	var backupsData []BackupData
	err := db.Select(&backupsData, "SELECT * FROM `backup` ORDER BY `Id` DESC")
	if err != nil {
		logger.Println("backup.GetBackupData: get backups data from bd error: ", err)
	}

	if len(backupsData) < 1 {
		ctx.JSON(http.StatusOK, "")
		return
	}

	ctx.JSON(http.StatusOK, backupsData)
}

func ZipPath(path, output string, uids []string) bool {
	if _, err := os.Stat("./backups/"); os.IsNotExist(err) {
		os.MkdirAll("./backups/", 0777)
	}

	outFile, err := os.Create(output)
	if err != nil {
		logger.Println("backup.CreateZip error: ", err)
		return false
	}
	defer outFile.Close()

	writer := zip.NewWriter(outFile)

	files, err := ioutil.ReadDir(path)
	if err != nil {
		logger.Println("backup.CreateZip error: ", err)
		return false
	}

	for _, file := range files {
		if !file.IsDir() {
			fileName := file.Name()
			addFile := true

			if len(uids) > 1 {
				addFile = false
				for _, uid := range uids {
					if strings.Contains(fileName, uid) {
						addFile = true
					}
				}
			}

			if !addFile {
				continue
			}

			dat, err := ioutil.ReadFile(path + fileName)
			if err != nil {
				logger.Println("backup.CreateZip error: ", err)
				return false
			}

			f, err := writer.Create(fileName)
			if err != nil {
				logger.Println("backup.CreateZip error: ", err)
				return false
			}
			_, err = f.Write(dat)
			if err != nil {
				logger.Println("backup.CreateZip error: ", err)
				return false
			}
		}
	}
	writer.Close()

	return true
}

func BackupLogs(comment string, del_logs bool, uids []string) (err error) {
	currentDate := time.Now().Format("15-04-2006-01-02")
	backupName := fmt.Sprintf("./backups/Taurus_%s.zip", currentDate)

	if ZipPath("./logs/", backupName, uids) {
		fileInfo, _ := os.Stat(backupName)
		_, err = db.Exec("INSERT INTO `backup` (`Date`, `Size`, `Comment`) VALUES (?, ?, ?)",
			currentDate,
			utils.ByteCountSI(fileInfo.Size()),
			comment)
		if err != nil {
			logger.Println("backup.BackupLogs error: ", err)
			return err
		}

		if del_logs {
			_, err = db.Exec("DELETE FROM `logs`")
			if err != nil {
				logger.Println("backup.BackupLogs: error ", err)
				return err
			}

			err = os.RemoveAll("./logs/")
			if err != nil {
				logger.Println("backup.BackupLogs: ", err)
				return err
			}
		}
	}

	return nil
}

type BackupForm struct {
	Cookie string `json:"cookie"`
	Form   struct {
		Date    string `json:"date"`
		Comment string `json:"comment"`
		DelLogs bool   `json:"delete"`
	} `json:"form"`
	Uids []string `json:"uids"`
}

func Create(ctx *gin.Context) {
	var backupForm BackupForm
	ctx.BindJSON(&backupForm)

	if !login.RootAuthCheck(backupForm.Cookie) {
		logger.Println("backup.Create: unauthorized user: ", ctx.ClientIP())
		return
	}

	err := BackupLogs(backupForm.Form.Comment, backupForm.Form.DelLogs, backupForm.Uids)
	if err != nil {
		logger.Println("backup.CreateBackup error: ", err)
	}

	statusErr := ""
	if err != nil {
		statusErr = err.Error()
	}

	ctx.JSON(http.StatusOK, gin.H{"err": statusErr})
}

func Delete(ctx *gin.Context) {
	var backupForm BackupForm
	ctx.BindJSON(&backupForm)

	if !login.RootAuthCheck(backupForm.Cookie) {
		logger.Println("backup.Delete: unauthorized user: ", ctx.ClientIP())
		return
	}

	_, err := db.Exec("DELETE FROM `backup` WHERE `Date` = ?", backupForm.Form.Date)
	if err != nil {
		logger.Println("backup.DeleteBackup: delete backup from db error: ", err)
	}

	err = os.Remove(fmt.Sprintf("./backups/Taurus_%s.zip", backupForm.Form.Date))
	if err != nil {
		logger.Println("backup.DeleteBackup: delete backup from disk error: ", err)
	}

	statusErr := ""
	if err != nil {
		statusErr = err.Error()
	}

	ctx.JSON(http.StatusOK, gin.H{"err": statusErr})
}

func Download(ctx *gin.Context) {
	var backupForm BackupForm
	ctx.BindJSON(&backupForm)

	if !login.RootAuthCheck(backupForm.Cookie) {
		logger.Println("backup.Download: unauthorized user: ", ctx.ClientIP())
		return
	}

	ctx.File("./backups/" + "Taurus_" + backupForm.Form.Date + ".zip")
}
