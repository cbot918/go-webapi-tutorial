package main

import (
	"app/internal/config"
	"app/internal/util"
	"log"

	"app/internal/controller"

	"github.com/gin-gonic/gin"
)

func main() {

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	db, err := util.NewDB(cfg)
	if err != nil {
		log.Fatal(err)
	}

	dao, err := dao.NewDao(db)
	if err != nil {
		log.Fatal(err)
	}

	usecase, err := usecase.NewUseCase(dao)
	if err != nil {
		log.Fatal(err)
	}

	controller, err := controller.NewController(usecase)
	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()

	router = setupRouter(router, controller)

}

func setupRouter(router *gin.Engine, controller controller.Controller) *gin.Engin {

	return router
}
