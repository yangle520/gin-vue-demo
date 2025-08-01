package system

import "gorm.io/gorm"

type User struct {
	ID        uint   `gorm:"primarykey"`
	LoginName string `gorm:"size:50;comment:登录名"`
	Pwd       string `gorm:"size:50;comment:密码"`
	UserName  string `gorm:"size:50;comment:用户名"`
	gorm.Model
}
