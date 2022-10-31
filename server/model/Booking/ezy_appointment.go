// 自动生成模板EzyAppointment
package Booking

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// EzyAppointment 结构体
type EzyAppointment struct {
	global.GVA_MODEL
	AppointmentName    string     `json:"appointmentName" form:"appointmentName" gorm:"column:appointment_name;comment:;"`
	Singer             string     `json:"singer" form:"singer" gorm:"column:singer;comment:;"`
	AppointmentDate    *time.Time `json:"appointmentDate" form:"appointmentDate" gorm:"column:appointment_date;comment:;"`
	StartAt            *time.Time `json:"startAt" form:"startAt" gorm:"column:start_at;comment:;"`
	EndAt              *time.Time `json:"endAt" form:"endAt" gorm:"column:end_at;comment:;"`
	AppointmentContent string     `json:"appointmentContent" form:"appointmentContent" gorm:"column:appointment_content;comment:;"`
	AppointmentNote    string     `json:"appointmentNote" form:"appointmentNote" gorm:"column:appointment_note;comment:;"`
	Stage              *int       `json:"stage" form:"stage" gorm:"column:stage;comment:;"`
  
	StageMap           string     `json:"stageMap" form:"stageMap" gorm:"column:stage_map;comment:;"`
	StageArea          string     `json:"stageArea" form:"stageArea" gorm:"column:stage_area;comment:;"`
	DisableIndex       string     `json:"disableIndex" form:"disableIndex" gorm:"column:disable_index;comment:;"`
	TotalSeat          *int       `json:"totalSeat" form:"totalSeat" gorm:"column:total_seat;comment:;"`
	HideStageIndex     *bool      `json:"hideStageIndex" form:"hideStageIndex" gorm:"column:hide_stage_index;comment:;"`
	AllowBus           *bool      `json:"allowBus" form:"allowBus" gorm:"column:allow_bus;comment:;"`
	CreatedBy          *int       `json:"createdBy" form:"createdBy" gorm:"column:created_by;comment:;"`
	Branch             *int       `json:"branch" form:"branch" gorm:"column:branch;comment:;"`
	FeaturedImage      string     `json:"featuredImage" form:"featuredImage" gorm:"column:featured_image;comment:;"`
	Status             string     `json:"status" form:"status" gorm:"column:status;comment:;"`
	Slug               string     `json:"slug" form:"slug" gorm:"column:slug";comment:;"`
  
  StageObject EzyStage `json:ezyStage" gorm:"foreignKey:Stage"`
}

// TableName EzyAppointment 表名
func (EzyAppointment) TableName() string {
	return "ezy_appointment"
}
