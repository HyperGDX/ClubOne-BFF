package system

import (
	"bff/model/common/response"

	"github.com/gin-gonic/gin"
)

type ForumApi struct{}

func (b *ForumApi) Posts(c *gin.Context) {
	response.FailWithMessage("验证码错误", c)
}
