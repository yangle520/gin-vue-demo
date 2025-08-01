package example

import "github.com/yangle520/gin-vue-demo/server/global"

type ExaCustomer struct {
	global.GVA_MODEL
	CustomerName      string `json:"customerName" form:"customerName" gorm:"comment:客户名"`             // 客户名
	CustomerPhoneData string `json:"customerPhoneData" form:"customerPhoneData" gorm:"comment:客户手机号"` // 客户手机号

}
