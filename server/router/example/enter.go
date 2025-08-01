package example

import "github.com/yangle520/gin-vue-demo/server/api"

type RouterGroup struct {
	CustomerRouter
}

var (
	exaCustomerApi = api.ApiGroupApp.ExampleApiGroup.CustomerApi
)
