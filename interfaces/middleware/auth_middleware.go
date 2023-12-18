package middleware

import (
	"net/http"

	enum "github.com/nansuri/godeep-base/domain/enum/error_enum"
	"github.com/nansuri/godeep-base/infrastructure/auth"
	logger "github.com/nansuri/godeep-base/utils/logger"
	utils "github.com/nansuri/godeep-base/utils/response_handler"

	"github.com/gin-gonic/gin"
)

func AuthorizationCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		var errEnc utils.ErrorResponseEncoder

		uid, err := auth.TokenValid(c.Request)
		if err != nil {
			logger.LogErrorWithtracer(c, "Authorization Error", err.Error())
			errEnc.CustomErrorResponse(c, http.StatusUnauthorized, enum.CodeInvalidAccessToken, enum.MessageInvalidAccessToken)
			c.Abort()
			return
		}
		c.Set("userId", uid)
		c.Next()
	}
}

func CORSHeaderSet() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}
