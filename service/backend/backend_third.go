package backend

import (
	"bff/global"
	"bff/model/system"
	"bff/utils"
	"fmt"
)

type BackendThirdService struct{}

//var

//@author: [piexlmax](https://github.com/piexlmax)
//@author: [SliverHorn](https://github.com/SliverHorn)
//@function: Login
//@description: 用户登录
//@param: u *model.SysUser
//@return: err error, userInter *model.SysUser

func (thirdService *BackendThirdService) GetOssPolicy() (system.OssPolicy, error) {
	policy := new(system.OssPolicy)
	url := fmt.Sprintf("%s/oss/policy", global.GVA_CONFIG.Backend.ThirdApi)
	err := utils.HttpGetJsonRes(url, policy)
	    if err != nil {
        panic(err)
    }
	return *policy, nil
}
