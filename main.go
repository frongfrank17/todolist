package main

import (
	"github.com/gin-gonic/gin"

	configs "Hugeman/configs"
	"Hugeman/database"
	repository "Hugeman/repository"
	"Hugeman/route"
)

func main() {
	run := gin.Default()
	cfg, err := configs.NewConfig()
	if err != nil {
		panic(err)
	}
	dbConnector, err := database.Connector(cfg.Postgress_host, cfg.Postgress_port, cfg.Postgress_username, cfg.Postgress_password, cfg.Postgress_db)
	if err != nil {
		panic(err)
	}
	err = dbConnector.AutoMigrate(&repository.Todo{})
	if err != nil {
		panic(err)
		//return
	}
	route.ApirountingInterfaces(run, dbConnector)
	//fmt.Println(dbConnector)
	//fmt.Println(cfg.PORT)
	run.Run(":" + cfg.PORT) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
