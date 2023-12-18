package base

import (
	"github.com/nansuri/godeep-base/interfaces/handlers/base"

	"github.com/gin-gonic/gin"
)

func BaseRouter(r *gin.RouterGroup, bases *base.Base) {
	baseRouter := r.Group("/base")
	baseRouter.GET("/heartBeat", bases.BaseHeartBeatCheck)
	baseRouter.GET("/games", bases.BaseHeartBeatCheck)
}
