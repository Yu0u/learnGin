package controller

import (
	"github.com/gin-gonic/gin"
	"learnGin/utils/errmsg"
	"learnGin/utils/validator"
	"net/http"
	"strconv"

	"learnGin/model"
)

var code int

func AddStudent(c *gin.Context) {
	var data model.Student
	var msg string
	_ = c.ShouldBindJSON(&data)
	msg, code = validator.Validate(&data)
	if code != errmsg.SUCCSE {
		c.JSON(http.StatusOK,gin.H{
			"code":code,
			"message":msg,
		})
		c.Abort()
		return
	}
	code = model.CheckStudent(data.StudentId)
	if code == errmsg.SUCCSE {
		model.AddStudent(&data)
	}
	if code == errmsg.STUDENT_EXIST {
		code = errmsg.STUDENT_EXIST
	}
	c.JSON(http.StatusOK,gin.H{
		"code":code,
		"message":errmsg.GetErrMsg(code),
	})
}

func GetStudentInfo(c *gin.Context) {
	id,_ := strconv.Atoi(c.Param("id"))
	data,code := model.GetStudent(id)
	c.JSON(http.StatusOK,gin.H{
		"code":code,
		"data":data,
		"message":errmsg.GetErrMsg(code),
	})
}

// 编辑用户
func EditStudent(c *gin.Context) {
	var data model.Student
	id, _ := strconv.Atoi(c.Param("student_id"))
	_ = c.ShouldBindJSON(&data)

	code = model.CheckUpUser(id, data.Name)
	if code == errmsg.SUCCSE {
		model.EditStudent(id, &data)
	}
	if code == errmsg.STUDENT_EXIST {
		c.Abort()
	}

	c.JSON(
		http.StatusOK, gin.H{
			"code":  code,
			"message": errmsg.GetErrMsg(code),
		},
	)
}

func GetStudent(c *gin.Context) {
	student, total := model.GetAllStudent()
	if total == 0 {
		c.JSON(http.StatusOK,gin.H{
			"code":errmsg.ERROR,
			"message":"没有学生",
		})
	}
	c.JSON(http.StatusOK,gin.H{
		"code":errmsg.SUCCSE,
		"data":student,
		"message":"查询成功",
	})
}

func DeleteStudent(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code = model.Delete(id)
	c.JSON(http.StatusOK,gin.H{
		"code":code,
		"message":errmsg.GetErrMsg(code),
	})
}