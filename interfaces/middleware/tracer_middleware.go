package middleware

import (
	"crypto/rand"
	"encoding/hex"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func randomHex(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func TracerInterceptor() gin.HandlerFunc {
	return func(c *gin.Context) {

		traceId, err := randomHex(20)
		if err != nil {
			logrus.Error("Error generating traceId")
		}

		c.Set("traceId", traceId)
		logrus.Debug("traceId: ", traceId)
		c.Next()
	}
}
