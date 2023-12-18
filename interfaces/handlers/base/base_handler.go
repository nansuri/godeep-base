package base

import (
	"context"
	"net/http"

	baseapp "bitbucket.org/be-proj/osp-base/application/base_app"
	clients "bitbucket.org/be-proj/osp-base/infrastructure/clients/osp"
	"bitbucket.org/be-proj/osp-base/infrastructure/clients/osp/entity"
	logger "bitbucket.org/be-proj/osp-base/utils/logger"
	utils "bitbucket.org/be-proj/osp-base/utils/response_handler"
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
	var userInfo entity.GetUserInfoResponse
	var rh utils.ResponseHandler
	errDetail := map[string]string{}

	ctx := context.WithValue(context.Background(), "traceId", "Test")

	// call the logic here
	contextValue, err := b.ba.BaseQuery(ctx, "test")
	if err != nil {
		logger.LogErrorWithtracer(c, "BaseHeartBeatCheck", err.Error())
	}

	// call getuserinfo from other service
	userInfo = b.us.GetUserInfoByAccessToken(c)

	// call Query User list
	request := make(map[string]interface{})
	request["fullName"] = ""
	request["email"] = ""
	request["domain"] = "Member"
	request["isActive"] = true
	request["pageNum"] = 1
	request["pageSize"] = 10

	listUser := b.us.QueryUserList(c, request)

	response := make(map[string]interface{})
	response["contextValue"] = contextValue
	response["userInfo"] = userInfo.UserData
	response["listUsers"] = listUser.UserData
	// return the response
	rh.ResponseEncoder(c, http.StatusAccepted, true, errDetail, "testResponse", response, true)

}
