package Booking

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Booking"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    BookingReq "github.com/flipped-aurora/gin-vue-admin/server/model/Booking/request"
)

type EzyCustomerService struct {
}

// CreateEzyCustomer Create EzyCustomer记录
// Author [piexlmax](https://github.com/piexlmax)
func (ezyCustomerService *EzyCustomerService) CreateEzyCustomer(ezyCustomer Booking.EzyCustomer) (err error) {
	err = global.GVA_DB.Create(&ezyCustomer).Error
	return err
}

// DeleteEzyCustomer Delete EzyCustomer记录
// Author [piexlmax](https://github.com/piexlmax)
func (ezyCustomerService *EzyCustomerService)DeleteEzyCustomer(ezyCustomer Booking.EzyCustomer) (err error) {
	err = global.GVA_DB.Delete(&ezyCustomer).Error
	return err
}

// DeleteEzyCustomerByIds 批量Delete EzyCustomer记录
// Author [piexlmax](https://github.com/piexlmax)
func (ezyCustomerService *EzyCustomerService)DeleteEzyCustomerByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]Booking.EzyCustomer{},"id in ?",ids.Ids).Error
	return err
}

// UpdateEzyCustomer 更新EzyCustomer记录
// Author [piexlmax](https://github.com/piexlmax)
func (ezyCustomerService *EzyCustomerService)UpdateEzyCustomer(ezyCustomer Booking.EzyCustomer) (err error) {
	err = global.GVA_DB.Save(&ezyCustomer).Error
	return err
}

// GetEzyCustomer 根据id获取EzyCustomer记录
// Author [piexlmax](https://github.com/piexlmax)
func (ezyCustomerService *EzyCustomerService)GetEzyCustomer(id uint) (ezyCustomer Booking.EzyCustomer, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&ezyCustomer).Error
	return
}

// GetEzyCustomerInfoList 分页获取EzyCustomer记录
// Author [piexlmax](https://github.com/piexlmax)
func (ezyCustomerService *EzyCustomerService)GetEzyCustomerInfoList(info BookingReq.EzyCustomerSearch) (list []Booking.EzyCustomer, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // Create db
	db := global.GVA_DB.Model(&Booking.EzyCustomer{})
    var ezyCustomers []Booking.EzyCustomer
    // 如果有条件搜索 下方会自动Create 搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.Name != "" {
        db = db.Where("name LIKE ?","%"+ info.Name+"%")
    }
    if info.Phone != "" {
        db = db.Where("phone LIKE ?","%"+ info.Phone+"%")
    }
    if info.Email != "" {
        db = db.Where("email LIKE ?","%"+ info.Email+"%")
    }
    if info.Membership != "" {
        db = db.Where("membership LIKE ?","%"+ info.Membership+"%")
    }
    if info.Discount != "" {
        db = db.Where("discount LIKE ?","%"+ info.Discount+"%")
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&ezyCustomers).Error
	return  ezyCustomers, total, err
}
