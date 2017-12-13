package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gopkg.in/asaskevich/govalidator.v4"
	"net/http"
	"service_test/model"
	"service_test/util"
	"time"
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
	c.JSON(http.StatusOK, token)
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
		c.JSON(http.StatusBadRequest, model.NewResponse(trans.TransactionId, ErrorStatusDone, err.Error(), trans.DateTime.UTC()))
		return
	}

	c.JSON(http.StatusOK, model.NewResponse(trans.TransactionId, SUCCESS, "success", trans.DateTime.UTC()))
}

func Service1(c *gin.Context) {
	trans := model.Transaction{}
	if err := c.BindJSON(&trans); err != nil {
		logrus.Error("Error, parse json failed, ", err)
		c.JSON(http.StatusBadRequest, model.NewResponse(trans.TransactionId, ErrorInvalid, err.Error(), time.Now().UTC()))
		return
	}
	//check token
	if !model.CheckToken(trans.Token) {
		logrus.Error("token is invalid")
		c.JSON(http.StatusBadRequest, model.NewResponse(trans.TransactionId, ErrorToken, "token is invalid", trans.DateTime.UTC()))
		return
	}
	//check validate
	_, err := govalidator.ValidateStruct(trans)
	if err != nil {
		logrus.Error("param is invalid, ", err)
		c.JSON(http.StatusBadRequest, model.NewResponse(trans.TransactionId, ErrorInvalid, err.Error(), trans.DateTime.UTC()))
		return
	}

	trans.Status = model.StatusInsert
	if err := trans.Insert(); err != nil {
		logrus.Error("insert failed, ", err)
		c.JSON(http.StatusBadRequest, model.NewResponse(trans.TransactionId, ErrorFailed, "insert failed", trans.DateTime.UTC()))
		return
	}
	c.JSON(http.StatusOK, model.NewResponse(trans.TransactionId, SUCCESS, "success", trans.DateTime.UTC()))
}

func Service2(c *gin.Context) {
	trans := model.Transaction{}
	if err := c.BindJSON(&trans); err != nil {
		logrus.Error("parse json failed, ", err)
		c.JSON(http.StatusBadRequest, model.NewResponse(trans.TransactionId, ErrorInvalid, err.Error(), time.Now().UTC()))
		return
	}
	//check token
	if !model.CheckToken(trans.Token) {
		logrus.Error("token is invalid, ")
		c.JSON(http.StatusBadRequest, model.NewResponse(trans.TransactionId, ErrorToken, "token is invalid", trans.DateTime.UTC()))
		return
	}
	if err := trans.CheckDoneAndUpdateReply(); err != nil {
		logrus.Error("CheckDoneAndUpdateReply failed, ", err)
		c.JSON(http.StatusBadRequest, model.NewResponse(trans.TransactionId, ErrorStatusDone, err.Error(), trans.DateTime.UTC()))
		return
	}
	c.JSON(http.StatusOK, model.NewResponse(trans.TransactionId, SUCCESS, "success", trans.DateTime.UTC()))
}
