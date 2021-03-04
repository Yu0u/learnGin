package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
	"strconv"
	"time"
)

func Upload(c *gin.Context) {

	file, _ := c.FormFile("file")
	ext := path.Ext(file.Filename)
	fileNameInt := time.Now().Unix()
	fileNameStr := strconv.FormatInt(fileNameInt, 10)
	fileName := fileNameStr + ext

	dst := fmt.Sprintf("./upload/%s%s", strconv.FormatInt(time.Now().Unix(), 10), fileName)
	err := c.SaveUploadedFile(file, dst)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "保存文件失败",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "保存成功",
			"url":     dst,
		})
	}

}
