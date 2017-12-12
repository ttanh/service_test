package handler

import (
	"github.com/gin-gonic/gin"
	"service_test/model"
	"net/http"
	"time"
	"gopkg.in/asaskevich/govalidator.v4"
	"log"
)

func Service1(c *gin.Context) {
	res := model.ResponseType{}
	//check token
	trans := model.TransactionInfo{}
	if err := c.BindJSON(&trans); err != nil {
		res = model.ResponseType{DateTime: time.Now().Unix(), Code: -1, Description: "Param invalid"}
		c.JSON(http.StatusOK, res)
		return
	}
	//check validate
	_, err := govalidator.ValidateStruct(trans)
	if err != nil {
		log.Println("param is invalid, ", err)
		res = model.ResponseType{DateTime: time.Now().Unix(), Code: -1, Description: "Param invalid"}
		c.JSON(http.StatusOK, res)
		return
	}


	c.JSON(http.StatusOK, trans)
}

func Service2(c *gin.Context) {

}