package backend

import (
	"bff/global"
	"bff/model/common/response"
	"bff/model/system"
	"bff/utils"
	"fmt"
)

type BackendForumService struct{}

//var

//@author: [piexlmax](https://github.com/piexlmax)
//@author: [SliverHorn](https://github.com/SliverHorn)
//@function: Login
//@description: 用户登录
//@param: u *model.SysUser
//@return: err error, userInter *model.SysUser

func (forumService *BackendForumService) GetPosts(channelId int) ([]system.Post, error) {
	url := fmt.Sprintf("%s/posts/channel/%d", global.GVA_CONFIG.Backend.ForumApi, channelId)
	var posts []system.Post
	res := response.Response{
    	Data: &posts,
	}
	err := utils.HttpGetJsonRes(url, res)
	if err != nil {
		return nil, err
	}
	
	return *res.Data.(*[]system.Post), nil
}
