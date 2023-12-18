package clients

import (
	"github.com/gin-gonic/gin"
	"github.com/nansuri/godeep-base/infrastructure/clients/osp/entity"
)

type OspUserServiceInterface interface {
	GetUserInfoByAccessToken(context *gin.Context) entity.GetUserInfoResponse
	QueryUserList(context *gin.Context, filterRequest map[string]interface{}) entity.QueryUserInfoResponse
}

type OspUserService struct {
}

var _ OspUserServiceInterface = &OspUserService{}

func NewOspUserService() *OspUserService {
	return &OspUserService{}
}

func (ous *OspUserService) GetUserInfoByAccessToken(context *gin.Context) entity.GetUserInfoResponse {
	var userInfo entity.GetUserInfoResponse
	InvokeOspGetMethod(context, "/v1/user/info", &userInfo)
	return userInfo
}

func (ous *OspUserService) QueryUserList(context *gin.Context, filterRequest map[string]interface{}) entity.QueryUserInfoResponse {
	var queryUserInfo entity.QueryUserInfoResponse

	InvokeOspPostMethod(context, "/v1/users/query", filterRequest, &queryUserInfo)
	return queryUserInfo
}
