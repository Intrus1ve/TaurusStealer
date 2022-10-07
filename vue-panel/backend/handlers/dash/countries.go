package dash

import (
	"fmt"
	"net/http"

	login "../login"
	"github.com/gin-gonic/gin"
)

type CountriesData struct {
	Country string `db:"Country" json:"country"`
	Count   int    `db:"Count" json:"count"`
	Percent string `json:"percent"`
}

func GetUserCountriesData(username string) (countriesData []CountriesData) {
	if username == "" {
		return
	}

	err := db.Select(&countriesData, "SELECT DISTINCT `Country`, COUNT(*) as Count FROM `logs` WHERE `ForUsers` LIKE ? GROUP BY `Country` ORDER BY Count DESC", "%"+username+"%")
	if err != nil {
		logger.Println("dash.GetUserCountriesData: get countries data from db error: ", err)
	}

	if len(countriesData) < 1 {
		return
	}

	logsCount := GetLogsCount()
	for i := 0; i < len(countriesData); i++ {
		if countriesData[i].Count > 0 && logsCount > 0 {
			countriesData[i].Percent = fmt.Sprintf("%0.2f", (float64(countriesData[i].Count)*float64(100))/float64(logsCount))
		}
	}

	return
}

func GetTopCountriesData(ctx *gin.Context) {
	username := login.CtxIsUserValid(ctx)
	if username == "" {
		logger.Println("dash.GetTopCountriesData: unauthorized user: ", ctx.ClientIP())
		return
	}

	countriesData := GetUserCountriesData(username)
	if len(countriesData) <= 5 {
		ctx.JSON(http.StatusOK, countriesData)
		return
	}

	ctx.JSON(http.StatusOK, countriesData[0:5])
}

func GetCountriesData(ctx *gin.Context) {
	username := login.CtxIsUserValid(ctx)
	if username == "" {
		logger.Println("dash.GetCountriesData: unauthorized user: ", ctx.ClientIP())
		return
	}
	ctx.JSON(http.StatusOK, GetUserCountriesData(username))
}
