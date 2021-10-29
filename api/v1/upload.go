package v1

import (
	"github.com/gin-gonic/gin"
	"goBlog/servers"
	"goBlog/utils/errmsg"
	"net/http"
)

func UpLoad(ctx *gin.Context)  {
	file, fileHeader, _ := ctx.Request.FormFile("file")
	fileSize := fileHeader.Size
	url, code := servers.UploadFile(file, fileSize)
	ctx.JSON(http.StatusOK, gin.H{
		"status": code,
		"message": errmsg.GetErrMsg(code),
		"url": url,
	})
}