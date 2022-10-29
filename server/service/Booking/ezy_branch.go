package Booking

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Booking"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    BookingReq "github.com/flipped-aurora/gin-vue-admin/server/model/Booking/request"
)

type EzyBranchService struct {
}

// CreateEzyBranch Create EzyBranch记录
// Author [piexlmax](https://github.com/piexlmax)
func (ezyBranchService *EzyBranchService) CreateEzyBranch(ezyBranch Booking.EzyBranch) (err error) {
	err = global.GVA_DB.Create(&ezyBranch).Error
	return err
}

// DeleteEzyBranch Delete EzyBranch记录
// Author [piexlmax](https://github.com/piexlmax)
func (ezyBranchService *EzyBranchService)DeleteEzyBranch(ezyBranch Booking.EzyBranch) (err error) {
	err = global.GVA_DB.Delete(&ezyBranch).Error
	return err
}

// DeleteEzyBranchByIds 批量Delete EzyBranch记录
// Author [piexlmax](https://github.com/piexlmax)
func (ezyBranchService *EzyBranchService)DeleteEzyBranchByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]Booking.EzyBranch{},"id in ?",ids.Ids).Error
	return err
}

// UpdateEzyBranch 更新EzyBranch记录
// Author [piexlmax](https://github.com/piexlmax)
func (ezyBranchService *EzyBranchService)UpdateEzyBranch(ezyBranch Booking.EzyBranch) (err error) {
	err = global.GVA_DB.Save(&ezyBranch).Error
	return err
}

// GetEzyBranch 根据id获取EzyBranch记录
// Author [piexlmax](https://github.com/piexlmax)
func (ezyBranchService *EzyBranchService)GetEzyBranch(id uint) (ezyBranch Booking.EzyBranch, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&ezyBranch).Error
	return
}

// GetEzyBranchInfoList 分页获取EzyBranch记录
// Author [piexlmax](https://github.com/piexlmax)
func (ezyBranchService *EzyBranchService)GetEzyBranchInfoList(info BookingReq.EzyBranchSearch) (list []Booking.EzyBranch, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // Create db
	db := global.GVA_DB.Model(&Booking.EzyBranch{})
    var ezyBranchs []Booking.EzyBranch
    // 如果有条件搜索 下方会自动Create 搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.Name != "" {
        db = db.Where("name LIKE ?","%"+ info.Name+"%")
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&ezyBranchs).Error
	return  ezyBranchs, total, err
}
