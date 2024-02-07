package system

import (
	v1 "bff/api/v1"

	"github.com/gin-gonic/gin"
)

type ThirdRouter struct{}

func (s *ThirdRouter) InitThirdRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	thirdRouter := Router.Group("third")
	thirdApi := v1.ApiGroupApp.SystemApiGroup.ThirdApi
	{
		thirdRouter.GET("oss/policy", thirdApi.GetOssPolicy)
	}
	return thirdRouter
}
