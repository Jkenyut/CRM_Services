package main

import (
	"crm_service/app/clients/connection"
	"crm_service/app/config"
	"crm_service/app/model"
	"crm_service/app/services/service_actor"
	"crm_service/app/services/services_auth"
	"fmt"
	ratelimit "github.com/JGLTechnologies/gin-rate-limit"
	helmet "github.com/danielkov/gin-helmet"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"time"
)

func main() {
	conf := config.GetConfig()
	conn := connection.NewConnection(conf, false)
	conn.Init(false)

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
	services_auth.NewServiceAuth(router, conf, conn, validators)
	service_actor.NewServiceActor(router, conf, conn, validators)

	//
	//customerHandler := contoller_customer.NewRouter(db)
	//customerHandler.Handle(router)
	//
	errRouter := router.Run(":8081")
	if errRouter != nil {
		panic(fmt.Sprint("error running server", errRouter))
	}

}
