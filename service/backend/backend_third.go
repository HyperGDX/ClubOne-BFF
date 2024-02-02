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

func (thirdService *BackendThirdService) GetOssPolicy() (system.OssPolicy, error) {
	url := fmt.Sprintf("%s/OSS/Policy", global.GVA_CONFIG.Backend.ThirdApi)
	res, err := utils.HttpGetJsonRes(url)
	if err != nil {
		return system.OssPolicy{}, err
	}
	if res.Data != nil {
		return *res.Data.(*system.OssPolicy), nil
	} else {
		return system.OssPolicy{}, nil
	}
}
