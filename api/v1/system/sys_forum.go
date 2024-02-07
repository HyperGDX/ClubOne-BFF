package system

import (
	"bff/model/common/response"
	sysReq "bff/model/system"
	sysRes "bff/model/system/response"
	"encoding/json"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ForumApi struct{}

func (b *ForumApi) GetPostsByChannel(c *gin.Context) {
	channelIdStr := c.Param("channelId")
	channelId, err := strconv.Atoi(channelIdStr)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	offsetStr := c.Query("pageIndex")
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

func (b *ForumApi) InsertPost(c *gin.Context) {
	bodyBytes, err := c.GetRawData()
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	post := new(sysReq.Post)
	err = json.Unmarshal(bodyBytes, &post)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	insertNum, err := forumService.InsertPost(post)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.Result(200, insertNum, "操作成功", c)
}

