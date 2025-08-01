package service

import (
	"github.com/yangle520/gin-vue-demo/server/service/example"
	"github.com/yangle520/gin-vue-demo/server/service/system"
)

var ServiceGroupApp = new(ServiceGroup)

type ServiceGroup struct {
	SystemServiceGroup  system.ServiceGroup
	ExampleServiceGroup example.ServiceGroup
}
