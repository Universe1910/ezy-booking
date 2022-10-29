// 自动生成模板EzyBranch
package Booking

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
)

// EzyBranch 结构体
type EzyBranch struct {
      global.GVA_MODEL
      Name  string `json:"name" form:"name" gorm:"column:name;comment:;"`
}


// TableName EzyBranch 表名
func (EzyBranch) TableName() string {
  return "ezy_branch"
}

