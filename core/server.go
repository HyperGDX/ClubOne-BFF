package core

import (
	"bff/global"
	"bff/initialize"
	"fmt"
	"time"

	//"bff/service/system"
	"go.uber.org/zap"
)

type server interface {
	ListenAndServe() error
}

func RunServer() {
	if global.GVA_CONFIG.System.UseMultipoint || global.GVA_CONFIG.System.UseRedis {
		// 初始化redis服务
		initialize.Redis()
	}
	// if global.GVA_CONFIG.System.UseMongo {
	// 	err := initialize.Mongo.Initialization()
	// 	if err != nil {
	// 		zap.L().Error(fmt.Sprintf("%+v", err))
	// 	}
	// }
	// // 从db加载jwt数据
	// if global.GVA_DB != nil {
	// 	system.LoadAll()
	// }

	Router := initialize.Routers()
	Router.Static("/form-generator", "./resource/page")

	address := fmt.Sprintf(":%d", global.GVA_CONFIG.System.Addr)
	s := initServer(address, Router)
	// 保证文本顺序输出
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)
	global.GVA_LOG.Info("server run success on ", zap.String("address", address))

	fmt.Printf(`
			start server
		`, address)
	global.GVA_LOG.Error(s.ListenAndServe().Error())
}
