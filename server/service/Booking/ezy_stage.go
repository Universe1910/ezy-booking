package Booking

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Booking"
	BookingReq "github.com/flipped-aurora/gin-vue-admin/server/model/Booking/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type EzyStageService struct {
}

// CreateEzyStage Create EzyStage记录
// Author [piexlmax](https://github.com/piexlmax)
func (ezyStageService *EzyStageService) CreateEzyStage(ezyStage Booking.EzyStage) (err error) {
	err = global.GVA_DB.Create(&ezyStage).Error
	return err
}

// DeleteEzyStage Delete EzyStage记录
// Author [piexlmax](https://github.com/piexlmax)
func (ezyStageService *EzyStageService) DeleteEzyStage(ezyStage Booking.EzyStage) (err error) {
	err = global.GVA_DB.Delete(&ezyStage).Error
	return err
}

// DeleteEzyStageByIds 批量Delete EzyStage记录
// Author [piexlmax](https://github.com/piexlmax)
func (ezyStageService *EzyStageService) DeleteEzyStageByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]Booking.EzyStage{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateEzyStage 更新EzyStage记录
// Author [piexlmax](https://github.com/piexlmax)
func (ezyStageService *EzyStageService) UpdateEzyStage(ezyStage Booking.EzyStage) (err error) {
	err = global.GVA_DB.Save(&ezyStage).Error
	return err
}

// GetEzyStage 根据id获取EzyStage记录
// Author [piexlmax](https://github.com/piexlmax)
func (ezyStageService *EzyStageService) GetEzyStage(id uint) (ezyStage Booking.EzyStage, err error) {
	err = global.GVA_DB.Where("id = ?", id).Preload("BranchObject").First(&ezyStage).Error
	return
}

// GetEzyStageInfoList 分页获取EzyStage记录
// Author [piexlmax](https://github.com/piexlmax)
func (ezyStageService *EzyStageService) GetEzyStageInfoList(info BookingReq.EzyStageSearch) (list []Booking.EzyStage, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// Create db
	db := global.GVA_DB.Model(&Booking.EzyStage{})
	var ezyStages []Booking.EzyStage
	// 如果有条件搜索 下方会自动Create 搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	if info.Address != "" {
		db = db.Where("address LIKE ?", "%"+info.Address+"%")
	}
	if info.MapURL != "" {
		db = db.Where("map_url LIKE ?", "%"+info.MapURL+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Preload("BranchObject").Find(&ezyStages).Error
	return ezyStages, total, err
}
