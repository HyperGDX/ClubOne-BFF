package router

import (
	//"bff/router/example"
	"bff/router/system"
)

type RouterGroup struct {
	System system.RouterGroup
	//Example example.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
