package base

import (
	"context"
	"net/http"

	baseapp "github.com/nansuri/godeep-base/application/base_app"
	clients "github.com/nansuri/godeep-base/infrastructure/clients/osp"
	logger "github.com/nansuri/godeep-base/utils/logger"
	utils "github.com/nansuri/godeep-base/utils/response_handler"

	"github.com/gin-gonic/gin"
)

type Base struct {
	ba baseapp.BaseAppInterface
	us clients.OspUserServiceInterface
}

func NewBase(ba baseapp.BaseAppInterface, us clients.OspUserServiceInterface) *Base {
	return &Base{
		ba: ba,
		us: us,
	}
}

func (b *Base) BaseHeartBeatCheck(c *gin.Context) {

	// define the utils and error detail as empty map
	var rh utils.ResponseHandler
	errDetail := map[string]string{}

	ctx := context.WithValue(context.Background(), "traceId", "Test")

	key := c.Query("key")

	// call the logic here
	theAnswer, err := b.ba.BaseQuery(ctx, key)
	if err != nil {
		logger.LogErrorWithtracer(c, "BaseHeartBeatCheck", err.Error())
	}

	// return the response
	rh.ResponseEncoder(c, http.StatusAccepted, true, errDetail, "theAnswer", theAnswer, true)

}
