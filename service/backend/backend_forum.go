package backend

import (
	"bff/global"
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

func (forumService *BackendForumService) GetPosts(channelId int, offset int) ([]system.Post, error) {
	url := fmt.Sprintf("%s/posts/channel/%d?offset=%d", global.GVA_CONFIG.Backend.ForumApi, channelId, offset)

	res, err := utils.HttpGetJsonRes(url)
	if err != nil {
		return nil, err
	}
	if res.Data != nil {
		return *res.Data.(*[]system.Post), nil
	} else {
		return []system.Post{}, nil
	}
	//return res.Data.([]system.Post), nil
}
