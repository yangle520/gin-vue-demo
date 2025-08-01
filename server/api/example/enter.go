package example

import "github.com/yangle520/gin-vue-demo/server/service"

type ApiGroup struct {
	CustomerApi
}

var (
	customerService = service.ServiceGroupApp.ExampleServiceGroup.CustomerService
)
