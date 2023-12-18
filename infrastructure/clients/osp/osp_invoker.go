package clients

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	utils "github.com/nansuri/godeep-base/utils/logger"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/tidwall/gjson"
	"gopkg.in/resty.v1"
)

func ospResponseValidator(c *gin.Context, resp *resty.Response) bool {
	if resp.StatusCode() != 200 && gjson.Get(string(resp.Body()), "success").Bool() {
		logrus.WithFields(logrus.Fields{"OSP Response ": string(resp.Body())}).Error("OSP Error")
		return false
	} else {
		return true
	}
}

func responseUnmarshaller(c *gin.Context, resp *resty.Response, response interface{}) bool {
	unmarshallError := json.Unmarshal(resp.Body(), &response)
	if unmarshallError != nil {
		utils.LogErrorWithtracer(c, "OSP Invoker Error\n", unmarshallError)
		return false
	} else {
		utils.LogInfoWithtracer(c, "OSP Invoker\n", "OSP response: ", string(resp.Body()))
		return true
	}
}

func InvokeOspGetMethod(c *gin.Context, path string, response interface{}) (invokeStatus bool) {

	// Get the token from existing
	OspToken := c.Request.Header.Get("Authorization")

	utils.LogInfoWithtracer(c, "OSP Invoker\n", "Path: ", path)

	// Invoke start here
	resp, _ := resty.R().
		SetHeader("Authorization", OspToken).
		SetHeader("Content-Type", "application/json").
		Get(viper.GetString("OSP_BE_HOST") + path)

	// Validate the response from osp
	if ospResponseValidator(c, resp) {
		// Unmarshaling the response
		return responseUnmarshaller(c, resp, &response)
	}
	return false
}

func InvokeOspPostMethod(c *gin.Context, path string, request map[string]interface{}, response interface{}) (invokeStatus bool) {
	// Get the token from existing
	OspToken := c.Request.Header.Get("Authorization")

	utils.LogInfoWithtracer(c, "OSP Invoker\n", "path: ", path, "\nRequest: ", request)

	// Marshall the map
	jsonString, _ := json.Marshal(request)

	// Invoke start here
	resp, _ := resty.R().
		SetHeader("Authorization", OspToken).
		SetHeader("Content-Type", "application/json").
		SetBody(jsonString).
		Post(viper.GetString("OSP_BE_HOST") + path)

	// Validate the response from osp
	if ospResponseValidator(c, resp) {
		// Unmarshaling the response
		return responseUnmarshaller(c, resp, &response)
	}
	return false
}
