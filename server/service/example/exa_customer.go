package example

import (
	"github.com/yangle520/gin-vue-demo/server/global"
	"github.com/yangle520/gin-vue-demo/server/model/common/request"
	"github.com/yangle520/gin-vue-demo/server/model/example"
)

type CustomerService struct{}

var CustomerServiceApp = new(CustomerService)

// 创建客户
func (exa *CustomerService) CreateExaCustomer(e example.ExaCustomer) (err error) {
	err = global.GVA_DB.Create(&e).Error
	return err
}

// 编辑用户
func (exa *CustomerService) UpdateExaCustomer(e example.ExaCustomer) (err error) {
	err = global.GVA_DB.Save(e).Error
	return err
}

// 删除客户
func (exa *CustomerService) DeleteExaCustomer(e example.ExaCustomer) (err error) {
	err = global.GVA_DB.Delete(&e).Error
	return err
}

// 查询详情
func (exa *CustomerService) GetExaCustomer(id uint) (customer example.ExaCustomer, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&customer).Error
	return
}

// 分页查询
func (exa *CustomerService) GetCustomerInfoList(info request.PageInfo) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)

	var CustomerList []example.ExaCustomer

	db := global.GVA_DB.Model(&example.ExaCustomer{})
	err = db.Count(&total).Error
	if err != nil {
		return CustomerList, total, err
	} else {
		err = db.Limit(limit).Offset(offset).Find(&CustomerList).Error
	}
	return CustomerList, total, err
}
