package initialize

import (
	"bff/global"
	"bff/middleware"
)

func OtherInit() {
	middleware.Proxy = middleware.NewReverseProxy(global.GVA_CONFIG.Backend.AuthApi)
	// dr, err := utils.ParseDuration(global.GVA_CONFIG.JWT.ExpiresTime)
	// if err != nil {
	// 	panic(err)
	// }
	// _, err = utils.ParseDuration(global.GVA_CONFIG.JWT.BufferTime)
	// if err != nil {
	// 	panic(err)
	// }

	// global.BlackCache = local_cache.NewCache(
	// 	local_cache.SetDefaultExpire(dr),
	// )
}
