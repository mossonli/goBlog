package v1

import (
	"github.com/gin-gonic/gin"
	"goBlog/model"
	"goBlog/utils/errmsg"
	"net/http"
	"strconv"
)

// 添加文章
func AddArticle(c *gin.Context) {
	var data model.Article
	_ = c.ShouldBindJSON(&data)

	code = model.CreateArt(&data)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// 查询分类下的所有文章
func GetCateArt(ctx *gin.Context)  {
	pageSize, _ := strconv.Atoi(ctx.Query("pagesize"))
	pageNum, _ := strconv.Atoi(ctx.Query("pagenum"))
	id, _ := strconv.Atoi(ctx.Param("id"))

	// 查询文章列表
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	data, code, _ := model.GetCateArt(id, pageSize, pageNum)

	ctx.JSON(http.StatusOK, gin.H{
		"status": code,
		"data":data,
		"message": errmsg.GetErrMsg(code),
	})
}

// 查询单个文章信息
func GetArtInfo(ctx *gin.Context)  {
	id, _ := strconv.Atoi(ctx.Param("id"))
	data, code := model.GetArtInfo(id)
	ctx.JSON(http.StatusOK, gin.H{
		"ststus": code,
		"data": data,
		"message": errmsg.GetErrMsg(code),
	})

}

// 查询单个文章
// todo 查询文章列表
func GetArt(ctx *gin.Context) {
	pageSize, _ := strconv.Atoi(ctx.Query("pagesize"))
	pageNum, _ := strconv.Atoi(ctx.Query("pagenum"))

	// 查询文章列表
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	data,code := model.GetArt(pageSize, pageNum)
	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// 编辑文章
func EditArt(c *gin.Context) {
	var data model.Article
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&data)

	code = model.EditArt(id, &data)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// 删除文章
func DeleteArt(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	code = model.DeleteArt(id)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
