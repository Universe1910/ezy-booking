// 自动生成模板EzyStage
package Booking

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// EzyStage 结构体
type EzyStage struct {
	global.GVA_MODEL
	Name     string `json:"name" form:"name" gorm:"column:name;comment:;"`
	Address  string `json:"address" form:"address" gorm:"column:address;comment:;"`
	MapURL   string `json:"mapURL" form:"mapURL" gorm:"column:map_url;comment:;"`
	BranchID *int   `json:"branchID" form:"branchID" gorm:"column:branch_id;comment:;"`

	BranchObject EzyBranch `json:"branchObject"  gorm:"foreignKey:BranchID"`
}

// TableName EzyStage 表名
func (EzyStage) TableName() string {
	return "ezy_stage"
}
