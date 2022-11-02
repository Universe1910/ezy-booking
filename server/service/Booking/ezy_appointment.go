package Booking

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Booking"
	BookingReq "github.com/flipped-aurora/gin-vue-admin/server/model/Booking/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type EzyAppointmentService struct {
}

// CreateEzyAppointment Create EzyAppointment记录
// Author [piexlmax](https://github.com/piexlmax)
func (ezyAppointmentService *EzyAppointmentService) CreateEzyAppointment(ezyAppointment Booking.EzyAppointment) (err error) {
	err = global.GVA_DB.Create(&ezyAppointment).Error
	return err
}

// DeleteEzyAppointment Delete EzyAppointment记录
// Author [piexlmax](https://github.com/piexlmax)
func (ezyAppointmentService *EzyAppointmentService) DeleteEzyAppointment(ezyAppointment Booking.EzyAppointment) (err error) {
	err = global.GVA_DB.Delete(&ezyAppointment).Error
	return err
}

// DeleteEzyAppointmentByIds 批量Delete EzyAppointment记录
// Author [piexlmax](https://github.com/piexlmax)
func (ezyAppointmentService *EzyAppointmentService) DeleteEzyAppointmentByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]Booking.EzyAppointment{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateEzyAppointment 更新EzyAppointment记录
// Author [piexlmax](https://github.com/piexlmax)
func (ezyAppointmentService *EzyAppointmentService) UpdateEzyAppointment(ezyAppointment Booking.EzyAppointment) (err error) {
	err = global.GVA_DB.Save(&ezyAppointment).Error
	return err
}

// GetEzyAppointment 根据id获取EzyAppointment记录
// Author [piexlmax](https://github.com/piexlmax)
func (ezyAppointmentService *EzyAppointmentService) GetEzyAppointment(id uint) (ezyAppointment Booking.EzyAppointment, err error) {
	// err = global.GVA_DB.Where("id = ?", id).First(&ezyAppointment).Error
	err = global.GVA_DB.Where("id = ?", id).Preload("StageObject").Preload("BranchObject").Preload("UserObject").First(&ezyAppointment).Error

	return
}

// GetEzyAppointmentInfoList 分页获取EzyAppointment记录
// Author [piexlmax](https://github.com/piexlmax)
func (ezyAppointmentService *EzyAppointmentService) GetEzyAppointmentInfoList(info BookingReq.EzyAppointmentSearch) (list []Booking.EzyAppointment, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// Create db
	db := global.GVA_DB.Model(&Booking.EzyAppointment{})
	var ezyAppointments []Booking.EzyAppointment
	// 如果有条件搜索 下方会自动Create 搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.AppointmentName != "" {
		db = db.Where("appointment_name LIKE ?", "%"+info.AppointmentName+"%")
	}
	if info.Singer != "" {
		db = db.Where("singer LIKE ?", "%"+info.Singer+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	// err = db.Limit(limit).Offset(offset).Find(&ezyAppointments).Error
	err = db.Limit(limit).Offset(offset).Preload("StageObject").Preload("BranchObject").Preload("UserObject").Find(&ezyAppointments).Error
	return ezyAppointments, total, err
}
