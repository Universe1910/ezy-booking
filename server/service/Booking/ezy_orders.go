package Booking

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Booking"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    BookingReq "github.com/flipped-aurora/gin-vue-admin/server/model/Booking/request"
)

type EzyOrdersService struct {
}

// CreateEzyOrders Create EzyOrders记录
// Author [piexlmax](https://github.com/piexlmax)
func (ezyOrdersService *EzyOrdersService) CreateEzyOrders(ezyOrders Booking.EzyOrders) (err error) {
	err = global.GVA_DB.Create(&ezyOrders).Error
	return err
}

// DeleteEzyOrders Delete EzyOrders记录
// Author [piexlmax](https://github.com/piexlmax)
func (ezyOrdersService *EzyOrdersService)DeleteEzyOrders(ezyOrders Booking.EzyOrders) (err error) {
	err = global.GVA_DB.Delete(&ezyOrders).Error
	return err
}

// DeleteEzyOrdersByIds 批量Delete EzyOrders记录
// Author [piexlmax](https://github.com/piexlmax)
func (ezyOrdersService *EzyOrdersService)DeleteEzyOrdersByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]Booking.EzyOrders{},"id in ?",ids.Ids).Error
	return err
}

// UpdateEzyOrders 更新EzyOrders记录
// Author [piexlmax](https://github.com/piexlmax)
func (ezyOrdersService *EzyOrdersService)UpdateEzyOrders(ezyOrders Booking.EzyOrders) (err error) {
	err = global.GVA_DB.Save(&ezyOrders).Error
	return err
}

// GetEzyOrders 根据id获取EzyOrders记录
// Author [piexlmax](https://github.com/piexlmax)
func (ezyOrdersService *EzyOrdersService)GetEzyOrders(id uint) (ezyOrders Booking.EzyOrders, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&ezyOrders).Error
	return
}

// GetEzyOrdersInfoList 分页获取EzyOrders记录
// Author [piexlmax](https://github.com/piexlmax)
func (ezyOrdersService *EzyOrdersService)GetEzyOrdersInfoList(info BookingReq.EzyOrdersSearch) (list []Booking.EzyOrders, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // Create db
	db := global.GVA_DB.Model(&Booking.EzyOrders{})
    var ezyOrderss []Booking.EzyOrders
    // 如果有条件搜索 下方会自动Create 搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&ezyOrderss).Error
	return  ezyOrderss, total, err
}
