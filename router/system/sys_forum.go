package system

import (
	v1 "bff/api/v1"

	"github.com/gin-gonic/gin"
)

type ForumRouter struct{}

func (s *ForumRouter) InitForumRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	forumRouter := Router.Group("forum")
	baseApi := v1.ApiGroupApp.SystemApiGroup.BaseApi
	{
		forumRouter.GET("posts", baseApi.Login)
		//baseRouter.POST("captcha", baseApi.Captcha)
	}
	return forumRouter
}
