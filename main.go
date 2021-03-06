package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"os"
	"service_test/handler"
	"service_test/model"
)

var db *gorm.DB

func setLogger() {
	logrus.SetFormatter(&logrus.TextFormatter{})
	logrus.SetOutput(os.Stdout)
	if gin.Mode() == gin.DebugMode {
		logrus.SetLevel(logrus.InfoLevel)
	}
}

func main() {
	fmt.Println("===Service start===")
	setLogger()
	var db *gorm.DB
	var err error
	if db, err = model.ConnectDB(); err != nil {
		logrus.Error("connect database failed: ", err)
		os.Exit(1)
	}
	defer db.Close()

	//Init Database
	model.InitDB()

	router := gin.Default()
	////setup router
	router.POST("/service1", handler.Service1)
	router.POST("/service2", handler.Service2)
	router.POST("/token", handler.CreateToken)
	router.POST("/process_done", handler.ProcessStatusDone)
	//
	router.Run(fmt.Sprintf(":%s", os.Getenv("API_PORT")))
	fmt.Println("===Service stop===")
}
