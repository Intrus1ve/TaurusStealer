package logs

import (
	"net/http"
	"os"
	"path/filepath"

	backup "../backup"
	login "../login"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type SelectedForm struct {
	Cookie   string `json:"cookie"`
	Selected struct {
		User          string   `json:"user"`
		DisallowUser  string   `json:"disallowUser"`
		CreateBackup  bool     `json:"createBackup"`
		BackupComment string   `json:"backupComment"`
		DeleteLogs    bool     `json:"deleteLogs"`
		Uid           []string `json:"uid"`
	} `json:"selected"`
}

func SelectedAction(ctx *gin.Context) {
	var selectedForm SelectedForm
	ctx.BindJSON(&selectedForm)

	if !login.RootAuthCheck(selectedForm.Cookie) {
		logger.Println("selected.SelectedAction: unauthorized user: ", ctx.ClientIP())
		return
	}

	var err error
	if selectedForm.Selected.User != "" {
		user := selectedForm.Selected.User + ","
		var query string
		var args []interface{}
		query, args, err = sqlx.In("UPDATE `logs` SET `ForUsers` = concat(`ForUsers`, ?) WHERE `Uid` IN (?)", user, selectedForm.Selected.Uid)
		if err != nil {
			logger.Println("logs.SelectedAction: creating sql-query error: ", err)
		}
		query = db.Rebind(query)
		_, err = db.Exec(query, args...)
		if err != nil {
			logger.Println("logs.SelectedAction: update ForUsers error: ", err)
		}
	}

	if selectedForm.Selected.DisallowUser != "" {
		user := selectedForm.Selected.DisallowUser + ","
		var query string
		var args []interface{}
		query, args, err = sqlx.In("UPDATE `logs` SET `ForUsers` = replace(`ForUsers`, ?, '') WHERE `Uid` IN (?)", user, selectedForm.Selected.Uid)
		if err != nil {
			logger.Println("logs.SelectedAction: creating sql-query error: ", err)
		}
		query = db.Rebind(query)
		_, err = db.Exec(query, args...)
		if err != nil {
			logger.Println("logs.SelectedAction: update ForUsers error: ", err)
		}
	}

	if selectedForm.Selected.CreateBackup {
		backup.BackupLogs(selectedForm.Selected.BackupComment, false, selectedForm.Selected.Uid)
	}

	if selectedForm.Selected.DeleteLogs {
		query, args, err := sqlx.In("DELETE FROM `logs` WHERE `Uid` IN (?)", selectedForm.Selected.Uid)
		if err != nil {
			logger.Println("logs.SelectedAction: creating sql-query for delete log-array error: ", err)
		}
		query = db.Rebind(query)
		_, err = db.Exec(query, args...)
		if err != nil {
			logger.Println("logs.SelectedAction: delete log from db error: ", err)
		}

		for _, logUid := range selectedForm.Selected.Uid {
			var country string
			err := db.Select(&country, "SELECT `Country` FROM `logs` WHERE `Uid` = ?", logUid)
			if err != nil {
				logger.Println("logs.SelectedAction: ", err)
				continue
			}

			err = os.Remove(filepath.Join("./logs", country+"_"+logUid+".zip"))
			if err != nil {
				logger.Println("logs.SelectedAction:delete log from disk error ", err)
			}
		}
	}

	statusErr := ""
	if err != nil {
		statusErr = err.Error()
	}

	ctx.JSON(http.StatusOK, gin.H{"err": statusErr})
}
