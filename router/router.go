package router

import (
	"github.com/gin-gonic/gin"
	"learnGin/controller"
	"learnGin/middleware"
	"learnGin/utils"
)

const (
	_  = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.New()
	r.Use(middleware.Log(),gin.Recovery())
	// 最大上传文件大小
	r.MaxMultipartMemory = 16 * MB

	router := r.Group("api")
	{
		// 增加
		router.POST("student/add",controller.AddStudent)
		// 根据学号得到
		router.GET("student/:id",controller.GetStudentInfo)
		// 查询所有成绩
		router.GET("students",controller.GetStudent)
		// 根据学号修改
		router.PUT("student/:id",controller.EditStudent)
		// 根据学号删除
		router.DELETE("student/:id",controller.DeleteStudent)
		// 文件上传
		router.POST("upload", controller.Upload)
		// 文件下载
		router.GET("download", controller.Download)

	}

	_ = r.Run(utils.HttpPort)
}
