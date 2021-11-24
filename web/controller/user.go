package controller

import (
	"net/http"
	"web-server/web/utils"

	"github.com/gin-gonic/gin"
)

func GetSession(ctx *gin.Context) {
	resp := make(map[string]string)
	resp["errno"] = utils.RECODE_ERR
	resp["errmsg"] = utils.RecodeText(utils.RECODE_ERR)

	ctx.JSON(http.StatusOK, resp)
}
