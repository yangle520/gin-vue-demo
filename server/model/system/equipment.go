package system

import "gorm.io/gorm"

type Equipment struct {
	gorm.Model
	Name      string `gorm:"size:50;comment:名称"`
	Class     string `gorm:"size:50;comment:分类"`
	Quality   string `gorm:"size:50;comment:品质"`
	Color     string `gorm:"size:50;comment:颜色"`
	Img       string `gorm:"size:50;comment:图片"`
	Attr1     string `gorm:"size:50;comment:武力"`
	Attr2     string `gorm:"size:50;comment:魔力"`
	Attr3     string `gorm:"size:50;comment:统帅"`
	Attr4     string `gorm:"size:50;comment:防御"`
	TotalAttr string `gorm:"size:50;comment:总属性"`
}
