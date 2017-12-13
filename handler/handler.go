package handler

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/asaskevich/govalidator.v4"
	"net/http"
	"service_test/model"
	"service_test/util"
	"time"
	"github.com/sirupsen/logrus"
)

const (
	SUCCESS         = 0
	ErrorInvalid    = -1
	ErrorToken      = -2
	ErrorFailed     = -3
	ErrorStatusDone = -4
)

func CreateToken(c *gin.Context) {
	tokenStr, _ := util.GenerateRandomString(20)
	token := model.Token{
		Value: tokenStr,
	}
	if err := token.Insert(); err != nil {
		c.JSON(http.StatusBadRequest, "failed")
		return
	}
	c.JSON(http.StatusOK, tokenStr)
}

func ProcessStatusDone(c *gin.Context) {
	res := model.ResponseType{}
	trans := model.Transaction{}
	if err := c.BindJSON(&trans); err != nil {
		logrus.Error("Error, parse json failed, ", err)
		res = model.ResponseType{DateTime: time.Now().UTC(), Code: ErrorInvalid, Description: err.Error()}
		c.JSON(http.StatusBadRequest, res)
		return
	}
	if err := trans.UpdateStatusDone(); err != nil {
		logrus.Error("UpdateStatusDone failed, ", err)
		res = model.ResponseType{DateTime: trans.DateTime.UTC(), Code: ErrorStatusDone, Description: err.Error(), TransactionId: trans.TransactionId}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	res = model.ResponseType{DateTime: trans.DateTime.UTC(), Code: SUCCESS, Description: "success", TransactionId: trans.TransactionId}
	c.JSON(http.StatusOK, res)
}

func Service1(c *gin.Context) {
	res := model.ResponseType{}
	trans := model.Transaction{}
	if err := c.BindJSON(&trans); err != nil {
		logrus.Error("Error, parse json failed, ", err)
		res = model.ResponseType{DateTime: time.Now().UTC(), Code: ErrorInvalid, Description: err.Error()}
		c.JSON(http.StatusBadRequest, res)
		return
	}
	//check token
	if !model.CheckToken(trans.Token) {
		logrus.Error("token is invalid")
		res = model.ResponseType{DateTime: trans.DateTime.UTC(), Code: ErrorToken, Description: "token is invalid"}
		c.JSON(http.StatusBadRequest, res)
		return
	}
	//check validate
	_, err := govalidator.ValidateStruct(trans)
	if err != nil {
		logrus.Error("param is invalid, ", err)
		res = model.ResponseType{DateTime: trans.DateTime.UTC(), Code: ErrorInvalid, Description: err.Error()}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	trans.Status = model.STATUS_INSERT
	if err := trans.Insert(); err != nil {
		logrus.Error("insert failed, ", err)
		res = model.ResponseType{DateTime: trans.DateTime.UTC(), Code: ErrorFailed, Description: "insert failed", TransactionId: trans.TransactionId}
		c.JSON(http.StatusBadRequest, res)
		return
	}
	res = model.ResponseType{DateTime: trans.DateTime.UTC(), Code: SUCCESS, Description: "success", TransactionId: trans.TransactionId}
	c.JSON(http.StatusOK, res)
}

func Service2(c *gin.Context) {
	res := model.ResponseType{}
	trans := model.Transaction{}
	if err := c.BindJSON(&trans); err != nil {
		logrus.Error("parse json failed, ", err)
		res = model.ResponseType{DateTime: time.Now().UTC(), Code: ErrorInvalid, Description: err.Error(), TransactionId: trans.TransactionId}
		c.JSON(http.StatusBadRequest, res)
		return
	}
	//check token
	if !model.CheckToken(trans.Token) {
		logrus.Error("token is invalid, ")
		res = model.ResponseType{DateTime: trans.DateTime.UTC(), Code: ErrorToken, Description: "token is invalid"}
		c.JSON(http.StatusBadRequest, res)
		return
	}
	if err := trans.CheckDoneAndUpdateReply(); err != nil {
		logrus.Error("CheckDoneAndUpdateReply failed, ", err)
		res = model.ResponseType{DateTime: trans.DateTime.UTC(), Code: ErrorStatusDone, Description: err.Error(), TransactionId: trans.TransactionId}
		c.JSON(http.StatusBadRequest, res)
		return
	}
	res = model.ResponseType{DateTime: trans.DateTime.UTC(), Code: SUCCESS, Description: "success", TransactionId: trans.TransactionId}
	c.JSON(http.StatusOK, res)
}
