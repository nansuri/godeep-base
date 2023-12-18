package base

import (
	"bitbucket.org/be-proj/osp-base/interfaces/handlers/base"
	"bitbucket.org/be-proj/osp-base/interfaces/middleware"
	"github.com/gin-gonic/gin"
)

func BaseRouter(r *gin.RouterGroup, bases *base.Base) {
	baseRouter := r.Group("/base")
	baseRouter.Use(middleware.AuthorizationCheck())
	baseRouter.GET("/heartBeat", bases.BaseHeartBeatCheck)
}
