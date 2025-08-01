package example

import (
	"github.com/gin-gonic/gin"
	"github.com/yangle520/gin-vue-demo/server/global"
	"github.com/yangle520/gin-vue-demo/server/model/common/response"
	"github.com/yangle520/gin-vue-demo/server/model/example"
	exampleRes "github.com/yangle520/gin-vue-demo/server/model/example/response"
	"github.com/yangle520/gin-vue-demo/server/utils"
	"go.uber.org/zap"
)

type CustomerApi struct{}

// 创建客户
func (e *CustomerApi) CreateExaCustomer(c *gin.Context) {
}

// 删除客户
func (e *CustomerApi) DeleteExaCustomer(c *gin.Context) {

}

// 更新客户信息
func (e *CustomerApi) UpdateExaCustomer(c *gin.Context) {
}

// 获取单一客户信息
func (e *CustomerApi) GetExaCustomer(c *gin.Context) {
	var customer example.ExaCustomer

	// 参数解析
	err := c.ShouldBindQuery(&customer)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 校验参数
	err = utils.Verify(customer.GVA_MODEL, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	data, err := customerService.GetExaCustomer(customer.ID)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(exampleRes.ExaCustomerResponse{Customer: data}, "success", c)
}

// 分页获取权限客户列表
func (e *CustomerApi) GetExaCustomerList(c *gin.Context) {
}
