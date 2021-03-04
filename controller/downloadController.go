package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Download(c *gin.Context) {
	content := c.Query("content")
	c.File("./upload/" + content)
	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", content))
	return
}
