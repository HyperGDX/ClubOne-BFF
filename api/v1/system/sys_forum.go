package system

import (
	"bff/model/common/response"

	"github.com/gin-gonic/gin"
)

type ForumApi struct{}

func (b *ForumApi) Posts(c *gin.Context) {
	res, err := forumService.GetPosts();
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	}
	response.Result(200, res, "操作成功", c)
}
