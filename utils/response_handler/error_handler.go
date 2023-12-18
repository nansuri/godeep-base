package utils

import (
	"encoding/json"
	"net/http"

	enum "github.com/nansuri/godeep-base/domain/enum/error_enum"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type ErrorResponseEncoder struct{}

func (errEnc *ErrorResponseEncoder) BadJSONResponse(c *gin.Context) {
	errDetail := map[string]string{}
	errDetail["errorCode"] = enum.CodeInvalidJson
	errDetail["errorMessage"] = enum.MessageInvalidJson

	errorLogJsonify(errDetail)

	c.JSON(http.StatusUnprocessableEntity, gin.H{
		"success":      false,
		"errorDetails": errDetail,
		"traceId":      c.GetString("traceId"),
	})
}

func (errEnc *ErrorResponseEncoder) InternalServerErrorResponse(c *gin.Context) {
	errDetail := map[string]string{}
	errDetail["errorCode"] = enum.CodeInternalServerError
	errDetail["errorMessage"] = enum.MessageInternalServerError

	errorLogJsonify(errDetail)

	c.JSON(http.StatusUnprocessableEntity, gin.H{
		"success":      false,
		"errorDetails": errDetail,
		"traceId":      c.GetString("traceId"),
	})
}

func (errEnc *ErrorResponseEncoder) CustomErrorResponse(c *gin.Context, statusCode int, errCode string, errMessage string) {
	errDetail := map[string]string{}
	errDetail["errorCode"] = errCode
	errDetail["errorMessage"] = errMessage

	errorLogJsonify(errDetail)

	c.JSON(statusCode, gin.H{
		"success":      false,
		"errorDetails": errDetail,
		"traceId":      c.GetString("traceId"),
	})
}

func (errEnc *ErrorResponseEncoder) CustomGeneralErrorResponse(c *gin.Context, statusCode int, errCode string, errMessage string, isSuccess bool) {
	errDetail := map[string]string{}
	errDetail["errorCode"] = errCode
	errDetail["errorMessage"] = errMessage

	errorLogJsonify(errDetail)

	c.JSON(statusCode, gin.H{
		"success":      isSuccess,
		"errorDetails": errDetail,
		"traceId":      c.GetString("traceId"),
	})
}

func (errEnc *ErrorResponseEncoder) CustomErrorStructResponse(c *gin.Context, statusCode int, errDetail map[string]string) {

	errorLogJsonify(errDetail)

	c.JSON(statusCode, gin.H{
		"success":      false,
		"errorDetails": errDetail,
		"traceId":      c.GetString("traceId"),
	})
}

func errorLogJsonify(data map[string]string) string {
	j, _ := json.Marshal(data)

	logrus.WithFields(logrus.Fields{"Response ": string(j)}).Error("Request Error")
	return string(j)
}
