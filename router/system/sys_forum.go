package system

import (
	v1 "bff/api/v1"

	"github.com/gin-gonic/gin"
)

type ForumRouter struct{}

func (s *ForumRouter) InitForumRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	forumRouter := Router.Group("forum")
	forumApi := v1.ApiGroupApp.SystemApiGroup.ForumApi
	{
		forumRouter.GET("posts/channel/:channelId", forumApi.GetPostsByChannel)
		forumRouter.GET("posts/OSS/Policy", forumApi.GetOssPolicy)
	}
	return forumRouter
}
