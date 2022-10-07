package gate

import (
	"io/ioutil"
	"net/http"

	utils "../../utils"

	"github.com/gin-gonic/gin"
)

func DllsHandler(ctx *gin.Context) {
	data, err := ioutil.ReadFile("./dlls/ff_dlls.bin")
	if err != nil {
		logger.Println("gate.dlls.DllsHandler:", err)
		ctx.Status(http.StatusOK)
		return
	}

	ctx.String(http.StatusOK, utils.EncryptData(string(data)))
}
