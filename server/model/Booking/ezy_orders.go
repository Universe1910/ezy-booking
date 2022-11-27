// 自动生成模板EzyOrders
package Booking

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
)

// EzyOrders 结构体
type EzyOrders struct {
	global.GVA_MODEL
	Quantity       int        `json:"quantity" form:"quantity" gorm:"column:quantity;comment:;"`
	Seats          string     `json:"seats" form:"seats" gorm:"column:seats;comment:;"`
	Total          float64    `json:"total" form:"total" gorm:"column:total;comment:;"`
	CouponCode     *uint      `json:"couponCode" form:"couponCode" gorm:"column:coupon_code;comment:;"`
	CouponDiscount *float64   `json:"couponDiscount" form:"couponDiscount" gorm:"column:coupon_discount;comment:;"`
	Tax            *float64   `json:"tax" form:"tax" gorm:"column:tax;comment:;"`
	InvoicePayment float64    `json:"invoicePayment" form:"invoicePayment" gorm:"column:invoice_payment;comment:;"`
	PaymentType    string     `json:"paymentType" form:"paymentType" gorm:"column:payment_type;comment:;"`
	AdminDiscount  *float64   `json:"adminDiscount" form:"adminDiscount" gorm:"column:admin_discount;comment:;"`
	ArrivedNumber  int        `json:"arrivedNumber" form:"arrivedNumber" gorm:"column:arrived_number;comment:Quantity has arrived;"`
	BusNumber      *int       `json:"busNumber" form:"busNumber" gorm:"column:bus_number;comment:;"`
	CreatedBy      *int       `json:"createdBy" form:"createdBy" gorm:"column:created_by;comment:;"`
	LastActionBy   *int       `json:"lastActionBy" form:"lastActionBy" gorm:"column:last_action_by;comment:;"`
	LastEditBy     *int       `json:"lastEditBy" form:"lastEditBy" gorm:"column:last_edit_by;comment:Author ID;"`
	LastActionTime *time.Time `json:"lastActionTime" form:"lastActionTime" gorm:"column:last_action_time;comment:;"`
	Source         *string    `json:"source" form:"source" gorm:"column:source;comment:;"`
	Notes          *string    `json:"notes" form:"notes" gorm:"column:notes;comment:;"`
	Status         string     `json:"status" form:"status" gorm:"column:status;comment:;"`
	CustomerID     uint       `json:"customer_id" form:"customer_id" gorm:"column:customer_id;comment:;"`
	AppointmentID  uint       `json:"appointmentId" form:"appointmentId" gorm:"column:appointment_id;comment:;"`

	CustomerObject     EzyCustomer    `json:"customerObject" form:"customerObject" gorm:"foreignKey:CustomerID"`
	CreatedByObject    system.SysUser `json:"createdByObject" form:"createdByObject" gorm:"foreignKey:CreatedBy"`
	LastActionByObject system.SysUser `json:"lastActionByObject" form:"lastActionByObject" gorm:"foreignKey:LastActionBy"`
}

// TableName EzyOrders 表名
func (EzyOrders) TableName() string {
	return "ezy_orders"
}
