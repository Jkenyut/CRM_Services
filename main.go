package main

import (
	"crm_service/app/clients/connection"
	"crm_service/app/config"
	"crm_service/app/model"
	"crm_service/app/services/service_actor"
	"fmt"
	ratelimit "github.com/JGLTechnologies/gin-rate-limit"
	helmet "github.com/danielkov/gin-helmet"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"time"
)

func main() {
	if err := config.NewConfig("local.env"); err != nil {
		panic(err)
	}
	conf := config.GetConfig()
	conn := connection.NewConnection(conf)
	conn.Init()

	validators := validator.New()
	router := gin.New()
	router.Use(cors.Default())
	router.Use(helmet.Default())
	store := ratelimit.InMemoryStore(&ratelimit.InMemoryOptions{
		Rate:  time.Minute * 1,
		Limit: 20,
	})
	mw := ratelimit.RateLimiter(store, &ratelimit.Options{
		ErrorHandler: model.ErrorHandler,
		KeyFunc:      model.KeyFunc,
	})
	router.Use(mw)

	service_actor.NewServiceActor(router, conf, conn, validators)

	//
	//customerHandler := customer.NewRouter(db)
	//customerHandler.Handle(router)
	//
	errRouter := router.Run(":8081")
	if errRouter != nil {
		fmt.Println("error running server", errRouter)
		return
	}
}
