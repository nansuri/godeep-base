package base

import (
	"github.com/nansuri/godeep-base/interfaces/handlers/base"
	"github.com/nansuri/godeep-base/interfaces/middleware"

	"github.com/gin-gonic/gin"
)

func BaseRouter(r *gin.RouterGroup, bases *base.Base) {
	baseRouter := r.Group("/base")
	baseRouter.Use(middleware.AuthorizationCheck())
	baseRouter.GET("/heartBeat", bases.BaseHeartBeatCheck)
}
