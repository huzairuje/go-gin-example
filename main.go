package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-gin-example/config"
	"github.com/go-gin-example/database"
	loanDomain "github.com/go-gin-example/loan"
	"github.com/go-gin-example/response"
	"github.com/go-gin-example/utils"
	"github.com/sirupsen/logrus"
)

func main() {
	//load config from config.yml
	loadCfg := config.LoadConfig()
	//initiate database connection
	db := database.ConnectDB(loadCfg)
	//initiate gin instance
	router := gin.Default()
	//custom response for not matching routes
	router.NoRoute(func(c *gin.Context) {
		response.NotFound(c, utils.NotMatchingAnyRoute, utils.NotFound)
	})
	//initiate routes by domain (this domain just loan domain)
	loanDomain.Routes(router, db)
	//start web (gin) service
	serverPort := loadCfg.ServerPort
	if serverPort != "" {
		logrus.Info(utils.ServerPortIsSet, serverPort)
		serverPort = loadCfg.ServerPort
	} else {
		logrus.Errorf(utils.ServerPortIsNotSet)
		serverPort = utils.DefaultServerPort
	}
	err := router.Run(serverPort)
	if err != nil {
		logrus.Info(err)
		panic(err)
	}

}
