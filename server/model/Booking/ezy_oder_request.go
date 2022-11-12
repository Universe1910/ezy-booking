// 自动生成模板EzyOrders
package Booking

// EzyOrders 结构体
type EzyOrdersRequest struct {
	// global.GVA_MODEL
	Quantity *int `json:"quantity" form:"quantity"`

	Seats string   `json:"seats" form:"seats"`
	Total *float64 `json:"total" form:"total"`

	CouponID       *int     `json:"couponID" form:"couponID"`
	CouponDiscount *float64 `json:"couponDiscount" form:"couponDiscount"`

	Tax            *float64 `json:"tax" form:"tax"`
	InvoicePayment *float64 `json:"invoicePayment" form:"invoicePayment"`

	PaymentType   string   `json:"paymentType" form:"paymentType"`
	AdminDiscount *float64 `json:"adminDiscount" form:"adminDiscount"`

	BusNumber *int `json:"busNumber" form:"busNumber"`
	CreatedBy *int `json:"createdBy" form:"createdBy"`

	Source string `json:"source" form:"source"`
	Status string `json:"status" form:"status"`

	CustomerID    *int `json:"customer_id" form:"customer_id" `
	AppointmentID *int `json:"appointmentId" form:"appointmentId"`

	CustomerName  string `json:"customerName" form:"customerName'`
	CustomerEmail string `json:"customerEmail" form:"customerEmail'`
	CustomerPhone string `json:"customerPhone" form:"customerPhone'`

	CustomerMembershipCode string `json:"customerMembershipCode" form:"customerMembershipCode"`
	
}
