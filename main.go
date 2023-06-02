package main

import (
	"crm_service/modules/actor"
	db2 "crm_service/utils/db"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	_ = godotenv.Load()
	db := db2.GormMysql()
	router := gin.New()

	actorHandler := actor.NewRouter(db)
	actorHandler.Handle(router)

	errRouter := router.Run(":8081")
	if errRouter != nil {
		fmt.Println("error running server", errRouter)
		return
	}
}
