package controller

import (
	"image/png"
	"net/http"
	"web-server/web/utils"

	"github.com/afocus/captcha"
	"github.com/gin-gonic/gin"
)

func GetSession(ctx *gin.Context) {
	resp := make(map[string]string)
	resp["errno"] = utils.RECODE_ERR
	resp["errmsg"] = utils.RecodeText(utils.RECODE_ERR)

	ctx.JSON(http.StatusOK, resp)
}

func GetImageCode(ctx *gin.Context) {
	var uuid = ctx.Param("uuid")
	cap := captcha.New()
	cap.SetFont("./conf/comic.ttf")
	img, str := cap.Create(6, captcha.ALL)
	png.Encode(ctx.Writer, img)
	println(str)
	println(uuid)
	// ctx.JSON(http.StatusOK, uuid)
}
