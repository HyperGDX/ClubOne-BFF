package system

import (
	"bff/model/common/response"

	"github.com/gin-gonic/gin"
)

type ThirdApi struct{}

func (b *ThirdApi) GetOssPolicy(c *gin.Context) {
	ossPolicy, err := thirdService.GetOssPolicy()
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Result(200, ossPolicy, "操作成功", c)
}
