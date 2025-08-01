package api

import (
	"github.com/yangle520/gin-vue-demo/server/api/example"
	"github.com/yangle520/gin-vue-demo/server/api/system"
)

var ApiGroupApp = new(ApiGroup)

type ApiGroup struct {
	SystemApiGroup  system.ApiGroup
	ExampleApiGroup example.ApiGroup
}
