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

func (forumService *BackendForumService) GetPosts(channelId int) (*[]system.Post, error) {
	url := fmt.Sprintf("%s/posts/channel/%d", global.GVA_CONFIG.Backend.ForumApi, channelId)
	posts := &[]system.Post{}
	err := utils.HttpGetJsonRes(url, &posts)
	if err != nil {
		return nil, err
	}
	fmt.Println(posts)
	return posts, nil
}
