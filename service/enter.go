package service

import (
	//"github.com/flipped-aurora/gin-vue-admin/server/service/example"
	"bff/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup system.ServiceGroup
	//ExampleServiceGroup example.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
