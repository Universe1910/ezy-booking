package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/Booking"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type EzyOrdersSearch struct{
    Booking.EzyOrders
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    request.PageInfo
}
