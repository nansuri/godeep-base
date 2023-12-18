package utils

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type ResponseHandler struct{}

func (rh *ResponseHandler) ResponseEncoder(c *gin.Context, statusCode int, status bool, errDetail map[string]string, dataName string, data interface{}, printLog bool) {

	response := make(map[string]interface{})
	response["success"] = status
	response["errorDetail"] = errDetail
	response["traceId"] = c.GetString("traceId")
	response[dataName] = data

	if printLog {
		LogJsonify(c, response)
	}

	c.JSON(statusCode, response)
}

func (rh *ResponseHandler) RegularResponseEncoder(c *gin.Context, statusCode int, status bool, errDetail map[string]string, printLog bool) {

	response := make(map[string]interface{})
	response["success"] = status
	response["errorDetail"] = errDetail
	response["traceId"] = c.GetString("traceId")

	if printLog {
		LogJsonify(c, response)
	}

	c.JSON(statusCode, response)
}

func LogJsonify(context *gin.Context, data map[string]interface{}) {
	resLimit := 500
	j, _ := json.Marshal(data)

	bodyRes := string(j)
	bodyResLen := len([]rune(bodyRes))

	if bodyResLen > resLimit {
		resLimit = 500
	} else {
		resLimit = bodyResLen
	}

	logrus.WithFields(logrus.Fields{
		"traceId":   context.GetString("traceId"),
		"Response ": string(j)[0:resLimit],
	}).Info("Response")
}
