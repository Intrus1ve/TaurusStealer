package dash

import (
	"fmt"
	"net/http"

	login "../login"
	"github.com/gin-gonic/gin"
)

type PrefixData struct {
	Prefix  string `db:"Prefix" json:"prefix"`
	Count   int    `db:"Count" json:"count"`
	Percent string `json:"percent"`
}

func GetUserPrefixData(username string) (prefixData []PrefixData) {
	if username == "" {
		return
	}

	err := db.Select(&prefixData, "SELECT DISTINCT `Prefix`, COUNT(*) as Count FROM `logs` WHERE `ForUsers` LIKE ? GROUP BY `Prefix` ORDER BY Count DESC", "%"+username+"%")
	if err != nil {
		logger.Println("dash.GetUserPrefixData: get countries data from db error: ", err)
	}

	if len(prefixData) < 1 {
		return
	}

	logsCount := GetLogsCount()
	for i := 0; i < len(prefixData); i++ {
		if prefixData[i].Count > 0 && logsCount > 0 {
			prefixData[i].Percent = fmt.Sprintf("%0.2f", (float64(prefixData[i].Count)*float64(100))/float64(logsCount))
		}
	}

	return
}

func GetTopPrefixData(ctx *gin.Context) {
	username := login.CtxIsUserValid(ctx)
	if username == "" {
		logger.Println("dash.GetTopPrefixData: unauthorized user: ", ctx.ClientIP())
		return
	}

	prefixData := GetUserPrefixData(username)
	if len(prefixData) <= 5 {
		ctx.JSON(http.StatusOK, prefixData)
		return
	}

	ctx.JSON(http.StatusOK, prefixData[0:5])
}

func GetPrefixData(ctx *gin.Context) {
	username := login.CtxIsUserValid(ctx)
	if username == "" {
		logger.Println("dash.GetPrefixData: unauthorized user: ", ctx.ClientIP())
		return
	}
	ctx.JSON(http.StatusOK, GetUserPrefixData(username))
}
