package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func LogInfoWithtracer(c *gin.Context, logInfo ...interface{}) {
	logrus.WithFields(logrus.Fields{
		"\ntraceId": c.GetString("traceId"),
	}).Info(logInfo...)
}

func LogErrorWithtracer(c *gin.Context, context string, logInfo ...interface{}) {
	logrus.WithFields(logrus.Fields{
		"\ncontext": logInfo,
		"\ntraceId": c.GetString("traceId"),
	}).Info(context)
}
