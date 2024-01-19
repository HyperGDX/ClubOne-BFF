package service

import (
	//"github.com/flipped-aurora/gin-vue-admin/server/service/example"
	"bff/service/backend"
	"bff/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup  system.ServiceGroup
	BackendServiceGroup backend.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
