package backend

import (
	"bff/global"
	"bff/model/system"
	"bff/model/system/response"
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

func (forumService *BackendForumService) GetPosts(channelId int, offset int) ([]response.Post, error) {
	url := fmt.Sprintf("%s/posts/channel/%d?offset=%d", global.GVA_CONFIG.Backend.ForumApi, channelId, offset)
	posts := new([]response.Post)
	err := utils.HttpGetJsonRes(url, posts)
	if err != nil {
		return nil, err
	}
	return *posts, nil
}

func (forumService *BackendForumService) InsertPost(post *system.Post) (response.InsertPostRes,error) {
	insertNum := new(response.InsertPostRes)
	url := fmt.Sprintf("%s/posts", global.GVA_CONFIG.Backend.ForumApi)
	err := utils.HttpPostJsonRes(url, post, insertNum)
	if err != nil {
		return *insertNum, err
	}
	return *insertNum, nil
}

