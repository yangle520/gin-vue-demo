package router

import (
	"github.com/yangle520/gin-vue-demo/server/router/example"
)

var RouterGroupApp = new(RouterGroup)

type RouterGroup struct {
	// System  system.RouterGroup
	Example example.RouterGroup
}
