package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"GoLearn/accountBank/config"
	"GoLearn/accountBank/controller"
	"GoLearn/accountBank/model"
	"GoLearn/accountBank/router"
)

// @title Service Auth
// @version 1.0
// @description Các API liên quan đến User.
// @host localhost:8080
func main() {
	r := gin.Default()

	// Doc cau hinh tu file Json
	config := config.SetupConfig()

	// Khoi tao Controller
	c := controller.NewController()
	c.Config = config

	// Ket noi CSDL
	dbConfig := config.Database
	db := model.ConnectDb(dbConfig.User, dbConfig.Password, dbConfig.Database, dbConfig.Address)
	defer db.Close()
	c.DB = db

	err := model.MigrationDb(db, config.ServiceName)
	if err != nil {
		log.Println(err)
		panic(err)
	}

	router.SetupRouter(config, r, c)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	err = r.Run(":" + port)
	if err != nil {
		panic(err)
	}
}

