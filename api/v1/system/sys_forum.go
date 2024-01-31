package system

import (
	"bff/model/common/response"
	sysRes "bff/model/system/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ForumApi struct{}

func (b *ForumApi) Posts(c *gin.Context) {
	channelIdStr := c.Param("channelId")
	channelId, err := strconv.Atoi(channelIdStr)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	offsetStr := c.Query("offset")
	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	posts, err := forumService.GetPosts(channelId, offset);
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	PostRes := make([]sysRes.Post, len(posts))
	for i, post := range posts {
		PostRes[i] = sysRes.Post{
			PostID:      post.PostID,
			UserName:    "user.UserName",
			UserAvatar:  "user.UserAvatar",
			PostTitle:   post.PostTitle,
			PostContent: post.PostContent,
			PostPics:    post.PostPics,
			UpdateTime:  post.UpdateTime,
			LikeCount:   post.LikeCount,
			ViewCount:   post.ViewCount,
		}
	}
	response.Result(200, PostRes, "操作成功", c)
}
