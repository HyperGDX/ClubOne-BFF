package backend

import (
	"bff/global"
	"bff/model/system"
	"bff/utils"
)

type BackendForumService struct{}

//var

//@author: [piexlmax](https://github.com/piexlmax)
//@author: [SliverHorn](https://github.com/SliverHorn)
//@function: Login
//@description: 用户登录
//@param: u *model.SysUser
//@return: err error, userInter *model.SysUser

func (forumService *BackendForumService) GetPosts() (*[]system.Post, error) {
	url := global.GVA_CONFIG.Backend.ForumApi + "/posts"
	posts := &[]system.Post{}
	err := utils.HttpGetJsonRes(url, &posts)
	if err != nil {
		return nil, err
	}
	print(posts)
	return posts, nil
}
