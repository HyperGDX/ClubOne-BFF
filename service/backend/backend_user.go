package backend

import (
	"bff/global"
	"bff/model/system"
	"bff/utils"
)

type BackendUserService struct{}

//var

//@author: [piexlmax](https://github.com/piexlmax)
//@author: [SliverHorn](https://github.com/SliverHorn)
//@function: Login
//@description: 用户登录
//@param: u *model.SysUser
//@return: err error, userInter *model.SysUser

func (userService *BackendUserService) Login(u *system.SysUser) (userInter *system.SysUser, err error) {
	// if nil == global.GVA_DB {
	// 	return nil, fmt.Errorf("db not init")
	// }
	url := global.GVA_CONFIG.Backend.BaseApi
	user := system.SysUser{}
	res, err := utils.HttpRequest(url + "/login")
	// err = global.GVA_DB.Where("username = ?", u.Username).Preload("Authorities").Preload("Authority").First(&user).Error
	// if err == nil {
	// 	if ok := utils.BcryptCheck(u.Password, user.Password); !ok {
	// 		return nil, errors.New("密码错误")
	// 	}
	// 	MenuServiceApp.UserAuthorityDefaultRouter(&user)
	// }
	print(res)
	return &user, nil
}
