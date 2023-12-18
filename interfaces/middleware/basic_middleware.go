package middleware

import (
	"bytes"
	"io"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func LoggerMiddleman() gin.HandlerFunc {
	return func(c *gin.Context) {
		var buf bytes.Buffer
		reqLimit := 500

		tee := io.TeeReader(c.Request.Body, &buf)
		body, _ := io.ReadAll(tee)
		c.Request.Body = io.NopCloser(&buf)

		bodyReq := string(body)
		bodyReqLen := len([]rune(bodyReq))

		if bodyReqLen > reqLimit {
			reqLimit = 500
		} else {
			reqLimit = bodyReqLen
		}

		logrus.WithFields(logrus.Fields{"Request": string(body)[0:reqLimit]}).Info("Request API : " + c.FullPath())
	}
}
