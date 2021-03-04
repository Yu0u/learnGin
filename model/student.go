package model

import (
	"gorm.io/gorm"
	"learnGin/utils/errmsg"
)

type Student struct {
	gorm.Model
	StudentId int `gorm:"type:int;not null;index;unique" json:"student_id" validate:"required" label:"学号"`
	Name      string  `gorm:"type:varchar(20);not null " json:"name" validate:"required" label:"姓名"`
	Score     float64 `gorm:"type:float;" json:"score" validate:"gte=0,lte=100" label:"分数"`
}

func CheckStudent(StudentId int) (code int) {
	var student Student
	db.Select("id").Where("student_id = ?",StudentId).First(&student)
	if student.ID > 0 {
		return errmsg.STUDENT_EXIST
	}
	return errmsg.SUCCSE
}

// 更新查询
func CheckUpUser(sid int, name string) (code int) {
	var student Student
	db.Select("student_id, name").Where("name", name).First(&student)
	if student.ID == uint(sid) {
		return errmsg.SUCCSE
	}
	if student.ID > 0 {
		return errmsg.STUDENT_EXIST
	}
	return errmsg.SUCCSE
}

func AddStudent(data *Student) (code int) {
	err := db.Create(&data).Error
	if err != nil{
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}
// 查询学生
func GetStudent(sid int) (Student, int) {
	var data Student
	err := db.Where("student_id = ?", sid).First(&data).Error
	if err != nil {
		return data, errmsg.ERROR
	}
	return data, errmsg.SUCCSE
}

// 编辑学生信息
func EditStudent(sid int, data *Student) int {
	var student Student
	var maps = make(map[string]interface{})
	maps["name"] = data.Name
	maps["score"] = data.Score
	err = db.Model(&student).Where("student_id = ?", sid).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// 查询所有学生信息
func GetAllStudent() ([]Student,int64) {
	var students []Student
	var total int64
	db.Find(&students)
	db.Model(&students).Count(&total)
	if err != nil {
		return students,0
	}
	return students,total
}

func Delete(sid int) int {
	var student Student
	err = db.Where("student_id = ?",sid).Delete(&student).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}