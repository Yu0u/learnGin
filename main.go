package main

import (
	"learnGin/model"
	"learnGin/router"
)

func main() {
	model.InitDb()
	router.InitRouter()
}
