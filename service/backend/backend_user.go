package backend

import (
	"bff/global"
	"bff/model/auth"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type BackendUserService struct{}

//var

//@author: [piexlmax](https://github.com/piexlmax)
//@author: [SliverHorn](https://github.com/SliverHorn)
//@function: Login
//@description: 用户登录
//@param: u *model.SysUser
//@return: err error, userInter *model.SysUser

func (userService *BackendUserService) Login(u *auth.LoginRequest) (userInter *auth.LoginResponse, err error) {
	// if nil == global.GVA_DB {
	// 	return nil, fmt.Errorf("db not init")
	// }
	url := global.GVA_CONFIG.Backend.AuthApi + "/login"
	loginReq := auth.LoginRequest{
		Username: u.Username,
		Password: u.Password,
	}

	// 将登录请求数据转换为 JSON 格式
	loginReqBytes, err := json.Marshal(loginReq)
	if err != nil {
		fmt.Println("Failed to marshal login request:", err)
		return
	}

	// 发送 POST 请求
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(loginReqBytes))
	if err != nil {
		fmt.Println("Failed to send login request:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, errors.New(resp.Status)
	}
	// 读取响应内容
	var loginResp auth.LoginResponse
	err = json.NewDecoder(resp.Body).Decode(&loginResp)
	if err != nil {
		fmt.Println("Failed to decode login response:", err)
		return
	}

	// 输出登录响应中的 token
	fmt.Println("Token:", loginResp.Token)
	return &loginResp, nil
}

// // 鉴权+认证
// func (userService *BackendUserService) Auth(u *system.SysUser) (userInter *system.SysUser, err error) {
// 	url := global.GVA_CONFIG.Backend.AuthApi
// 	user := system.SysUser{}
// 	res, err := utils.HttpRequest(url + "/login")
// 	// err = global.GVA_DB.Where("username = ?", u.Username).Preload("Authorities").Preload("Authority").First(&user).Error
// 	// if err == nil {
// 	// 	if ok := utils.BcryptCheck(u.Password, user.Password); !ok {
// 	// 		return nil, errors.New("密码错误")
// 	// 	}
// 	// 	MenuServiceApp.UserAuthorityDefaultRouter(&user)
// 	// }
// 	print(res)
// 	return &user, nil
// }
