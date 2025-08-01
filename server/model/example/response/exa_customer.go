package response

import "github.com/yangle520/gin-vue-demo/server/model/example"

type ExaCustomerResponse struct {
	Customer example.ExaCustomer `json:"customer"`
}
