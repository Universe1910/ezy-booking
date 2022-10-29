// 自动生成模板EzyCustomer
package Booking

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
)

// EzyCustomer 结构体
type EzyCustomer struct {
      global.GVA_MODEL
      Name  string `json:"name" form:"name" gorm:"column:name;comment:;"`
      Phone  string `json:"phone" form:"phone" gorm:"column:phone;comment:;"`
      Email  string `json:"email" form:"email" gorm:"column:email;comment:;"`
      Membership  string `json:"membership" form:"membership" gorm:"column:membership;comment:;"`
      Discount  string `json:"discount" form:"discount" gorm:"column:discount;comment:;"`
}


// TableName EzyCustomer 表名
func (EzyCustomer) TableName() string {
  return "ezy_customer"
}

